package middleware

import (
	"DormAppBackend/controller"

	"github.com/gin-gonic/gin"
)

//AuthMiddleware ...
type AuthMiddleware struct{}

var (
	auCtrl = new(controller.AuthController)
)

//TokenAuth authentication for token
func (auMid AuthMiddleware) TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auCtrl.TokenValid(c)
		c.Next()
	}
}

//CheckRoleLevelMid allow user with given role level to access
func (auMid AuthMiddleware) CheckRoleLevelMid(level int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		auCtrl.CheckRoleLevel(c, level)
		c.Next()
	}
}
