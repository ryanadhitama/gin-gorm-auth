package router

import (
	"gin-gorm-auth/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(authController *controller.AuthController) *gin.Engine {
	service := gin.Default()

	router := service.Group("/auth")

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	return service
}
