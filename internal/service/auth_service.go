package service

import (
	"errors"
	"gin-webapp/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	Config *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{Config: cfg}
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *AuthService) Authenticate(username, password string, expired bool) (string, string, error) {
	if username != "admin" || password != "password" {
		return "", "", errors.New("invalid username or password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	if expired {
		expirationTime = time.Now()
	}
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.Config.Signature))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.GenerateRefreshToken(username)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshToken, nil
}

func (s *AuthService) GenerateRefreshToken(username string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.Config.Signature))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config.Signature), nil
	})

	if err != nil || !token.Valid {
		return "", "", errors.New("invalid refresh token")
	}

	newToken, newRefreshToken, err := s.Authenticate(claims.Username, "password", false)
	if err != nil {
		return "", "", err
	}

	return newToken, newRefreshToken, nil
}
