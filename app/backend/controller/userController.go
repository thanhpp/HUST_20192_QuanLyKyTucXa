package controller

import (
	"net/http"

	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/model"
	"github.com/gin-gonic/gin"
)

// UserController controller for user
type UserController struct{}

//GetUserByUsername return an user by username
func (u UserController) GetUserByUsername(c *gin.Context) {
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
