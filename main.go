package main

import (
	"fmt"
	"gin-webapp/config"
	controller "gin-webapp/internal/controller"
	middleware "gin-webapp/internal/middleware"
	"gin-webapp/internal/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	port := os.Getenv("PORT")
	signature := os.Getenv("SIGNATURE")
	if port == "" || signature == "" {
		fmt.Println("Error: PORT and SIGNATURE must be set in .env file")
		return
	}

	cfg := &config.Config{
		Port:      port,
		Signature: signature,
	}

	authService := service.NewAuthService(cfg)
	authController := controller.NewAuthController(authService)

	router := gin.Default()

	enquiryController := controller.NewEnquiryController()

	router.POST("/auth/login", authController.Login)
	router.POST("/auth/refresh", authController.Refresh)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(authService))
	{
		api.POST("/enquiry/save", enquiryController.SaveEnquiry)
	}

	router.Run(":8080")
}
