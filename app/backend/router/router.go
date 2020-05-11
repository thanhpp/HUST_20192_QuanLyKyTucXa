package router

import (
	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/controller"
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

	//TODO: Auth middlewares

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			userCtrl := new(controller.UserController)
			userGroup.GET("/:usr", userCtrl.GetUserByUsername)
		}
	}

	return router
}
