package controller

import (
	"DormAppBackend/model"
	"DormAppBackend/tlog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//FacilityController ....
type FacilityController struct{}

var (
	facMod    = new(model.Facility)
	facManMod = new(model.FacilitiesManage)
)

func (fCtrl FacilityController) GetFacilityByFacID(c *gin.Context) {
	facID := c.Param("id")

	searchFacID, err := strconv.Atoi(facID)
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Can not parse id to int",
			"error":   err,
		})
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "not an ID for facility search",
		})
		c.Abort()
		return
	}

	returnFac, err := facMod.GetFacInfo(searchFacID)
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Can not find facility in db",
			"error":   err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not find this facility",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Get facility info success",
		"facility": returnFac,
	})
}

func (fCtrl FacilityController) GetListFacByRoomID(c *gin.Context) {
	roomID := c.Param("roomid")

	searchRoomID, err := strconv.Atoi(roomID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid roomID",
		})
		c.Abort()
		return
	}

	listFac, err := facManMod.GetFacListByRoom(searchRoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list facilities",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Get list facilities by roomid OK",
		"list_facilities": listFac,
	})
}
