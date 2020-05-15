package router

import (
	"DormAppBackend/controller"

	"github.com/gin-gonic/gin"
)

// NewRouter routing
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthGroup := router.Group("health")
	{
		healthCtrl := new(controller.Health)
		healthGroup.GET("check", healthCtrl.Check)
	}

	userGroup := router.Group("user")
	userGroup.Use()
	{
		userCtrl := new(controller.UserController)
		userGroup.POST("/login", userCtrl.Login)
		userGroup.GET("/logout", userCtrl.Logout)
	}

	return router
}
