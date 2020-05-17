package controller

import (
	"net/http"

	"DormAppBackend/forms"
	"DormAppBackend/model"

	"github.com/gin-gonic/gin"
)

var userModel = new(model.User)

// UserController controller for user
type UserController struct{}

//GetUserByUsername return an user by username
func (u UserController) getUserByUsername(c *gin.Context) {
	if c.Param("usr") != "" {
		user := new(model.User)

		err := user.GetByUsr(c.Param("usr"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error to retrive user",
				"error":   err,
			})

			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User founded",
			"user":    user,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "bad request",
	})
}

//Login validate user and return token
func (u UserController) Login(c *gin.Context) {
	var loginForm forms.LoginForm

	if c.ShouldBindJSON(&loginForm) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid form",
		})
		c.Abort()
		return
	}

	user, token, err := userModel.Login(loginForm)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid login details",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "login success",
			"user":    user,
			"token":   token,
		})
	}
}

//Logout remove user accessToken
func (u UserController) Logout(c *gin.Context) {
	au, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not logged in",
		})
		return
	}

	deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully",
	})
}

//TODO: Register

//Register create new user
func (u UserController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm
	if c.ShouldBindJSON(&registerForm) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid form",
		})
		c.Abort()
		return
	}

	user, err := userModel.Register(registerForm)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register successfully",
		"user":    user,
	})

}
