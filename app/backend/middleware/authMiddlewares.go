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
