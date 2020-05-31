package router

import (
	"DormAppBackend/controller"
	"DormAppBackend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func configCORS() gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"http://52.15.50.37:9090", "http://127.0.0.1:8080", "*"}
	cfg.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "authorization"}
	return cors.New(cfg)
}

// NewRouter routing
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(configCORS())
	// router.Use(cors.Default())
	router.Use(gin.Logger())

	authCtrl := new(controller.AuthController)
	router.POST("/tokenrf", authCtrl.Refresh)

	healthGroup := router.Group("health")
	{
		healthCtrl := new(controller.Health)
		healthGroup.GET("check", healthCtrl.Check)
	}

	userGroup := router.Group("user")
	{
		userCtrl := new(controller.UserController)
		userGroup.POST("/register", userCtrl.Register)
		userGroup.POST("/login", userCtrl.Login)
		userGroup.GET("/logout", userCtrl.Logout)
	}

	var authMid = new(middleware.AuthMiddleware)
	stdCtrl := new(controller.StudentController)
	roomCtrl := new(controller.RoomController)
	facCtrl := new(controller.FacilityController)
	rqCtrl := new(controller.RequestController)
	userCtrl := new(controller.UserController)

	level1 := router.Group("/lv1")
	{
		level1.Use(authMid.TokenAuth())
		level1.Use(authMid.CheckRoleLevelMid(1))

		level1.GET("/check/:usr", userCtrl.GetUserByUsername)
		level1.GET("/studentinfo", stdCtrl.GetStudentInfoLV1)
		level1.GET("/studentbyroom", stdCtrl.GetStudentByRoomID)

	}

	level0 := router.Group("/lv0")
	{
		level0.Use(authMid.TokenAuth())
		level0.Use(authMid.CheckRoleLevelMid(0))

		level0.GET("/usrinfo", stdCtrl.GetStudentInfo)
		level0.GET("/friends", stdCtrl.GetFriends)
		level0.GET("/roominfo", roomCtrl.GetRoomInfo)
		level0.GET("/dormmoney", stdCtrl.GetDormMoney)
		level0.GET("/fac/:id", facCtrl.GetFacilityByFacID)
		level0.GET("/listreq", rqCtrl.ListRequestByStudentID)
		level0.POST("/sendreq", rqCtrl.NewRequest)
	}
	router.NoRoute()
	return router
}
