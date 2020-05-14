package controller

import (
	"DormAppBackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthController .....
type AuthController struct{}

var authModel = new(model.AuthModel)

//TokenValid token validation
func (authCtrl AuthController) TokenValid(c *gin.Context) {
	err := authModel.TokenValid(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization, please login again",
		})
		c.Abort()
		return
	}
}
