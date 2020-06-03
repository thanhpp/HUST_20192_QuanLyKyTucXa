package controller

import (
	"DormAppBackend/forms"
	"DormAppBackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationController struct{}

var notiMod = new(model.Notification)

func (notiCtrl NotificationController) GetAllNoti(c *gin.Context) {
	listNoti, err := notiMod.GetAllNotification()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list notifications",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           "Get all notification OK",
		"list_notification": listNoti,
	})
}

func (notiCtrl NotificationController) CreateNoti(c *gin.Context) {
	var notiForm forms.NotificationForm
	if c.ShouldBindJSON(&notiForm) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid notification form",
		})
		c.Abort()
		return
	}

	noti, err := notiMod.CreateNotification(notiForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not create new notification",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Create new notification OK",
		"notification": noti,
	})
}
