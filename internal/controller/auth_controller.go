package controller

import (
	"gin-webapp/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Expired  bool   `json:"expired"`
	}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, refreshToken, err := ac.AuthService.Authenticate(loginReq.Username, loginReq.Password, loginReq.Expired)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	var refreshReq struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := ctx.ShouldBindJSON(&refreshReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, newRefreshToken, err := c.AuthService.RefreshToken(refreshReq.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": newRefreshToken})
}
