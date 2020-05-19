package router

import (
	"DormAppBackend/controller"
	"DormAppBackend/middleware"

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

	// authMiddleware := new(middleware.AuthMiddleware)
	userGroup := router.Group("user")
	{
		userCtrl := new(controller.UserController)
		userGroup.POST("/register", userCtrl.Register)
		userGroup.POST("/login", userCtrl.Login)
		userGroup.GET("/logout", userCtrl.Logout)
	}

	level1 := router.Group("/lv1")
	var authMid = new(middleware.AuthMiddleware)
	{
		userCtrl := new(controller.UserController)
		level1.Use(authMid.CheckRoleLevelMid(1))
		level1.GET("/check/:usr", userCtrl.GetUserByUsername)
	}

	router.NoRoute()
	return router
}
