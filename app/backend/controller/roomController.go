package controller

import (
	"DormAppBackend/model"
	"DormAppBackend/tlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RoomController ...
type RoomController struct{}

var roomMod = new(model.Room)

//GetRoomInfo ...
func (rCtrl RoomController) GetRoomInfo(c *gin.Context) {
	accessDes, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Canot extract metadata from Token",
			"error":   err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get room info",
		})
		c.Abort()
		return
	}

	roomID, err := studenMod.GetRoomID(int(accessDes.UserID))
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Can not get room id from student id",
			"error":   err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get room info",
		})
		c.Abort()
		return
	}

	room, err := roomMod.GetRoomInfo(roomID)
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Can not get room info from room id",
			"error":   err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get room info",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get room info successfully",
		"room":    room,
	})
}

func (rCtrl RoomController) GetAllRoom(c *gin.Context) {
	listRoom, err := roomMod.GetAllRoom()
	if err != nil {
		tlog.Info(tlog.Itf{
			"msg": "Can not get list room",
			"err": err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list room",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Get list room OK",
		"list_rooms": listRoom,
	})
}
