package controller

import (
	"DormAppBackend/forms"
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

func (fCtrl FacilityController) GetListFacByStudentID(c *gin.Context) {
	accessDes, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list facilities",
		})
		c.Abort()
		return
	}

	roomid, err := studenMod.GetRoomID(int(accessDes.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list facilities",
		})
		c.Abort()
		return
	}

	listFac, err := facManMod.GetFacListByRoom(roomid)
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

func (fCtrl FacilityController) NewFacility(c *gin.Context) {
	var newFacForm forms.NewFacilityForm
	if c.ShouldBindJSON(&newFacForm) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid new facility form",
		})
		c.Abort()
		return
	}

	newFac, err := facMod.NewFacility(newFacForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not create new facility",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Create new facility OK",
		"facility": newFac,
	})
}

func (fCtrl FacilityController) NewFacilityManage(c *gin.Context) {
	var newFacMngForm forms.NewFacilityManageForm
	if c.ShouldBindJSON(&newFacMngForm) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid new facility manage form",
		})
		c.Abort()
		return
	}

	newFacMng, err := facManMod.NewFacilityManage(newFacMngForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not create new facility manage",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Create new facility manage OK",
		"facility_manage": newFacMng,
	})
}

func (fCtrl FacilityController) UpdateFacilityManageInfo(c *gin.Context) {
	facMngID := c.Query("fmngid")
	quantity := c.Query("quantity")

	if facMngID == "" || quantity == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Not enough arguments",
		})
		c.Abort()
		return
	}

	facMngIDInt, err := strconv.Atoi(facMngID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid arguments",
		})
		c.Abort()
		return
	}

	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid arguments",
		})
		c.Abort()
		return
	}

	returnFacMng, err := facManMod.UpdateFacilityManage(facMngIDInt, quantityInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not update facility manage",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Update facility manage OK",
		"facility_manage": returnFacMng,
	})
}
