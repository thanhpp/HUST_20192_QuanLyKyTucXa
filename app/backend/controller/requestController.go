package controller

import (
	"DormAppBackend/forms"
	"DormAppBackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestController struct{}

var rqModel model.Request

func (rCtrl RequestController) NewRequest(c *gin.Context) {
	var requestForm forms.NewRequestForm

	accessDes, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can create request",
		})
		c.Abort()
		return
	}

	if c.ShouldBindJSON(&requestForm) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid form",
		})
		c.Abort()
		return
	}

	err, request := rqModel.NewRequest(requestForm, int(accessDes.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can create request",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "New request created",
		"request": request,
	})

}
