package controller

import (
	"DormAppBackend/model"
	"DormAppBackend/tlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

//StudentController ...
type StudentController struct{}

var studenMod = new(model.Student)

//GetStudentInfo ...
func (sCtrl StudentController) GetStudentInfo(c *gin.Context) {
	accessDes, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not find user",
		})
		c.Abort()
		return
	}
	returnStudent, err := studenMod.GetStudentInfo(int(accessDes.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not find user",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "User found",
		"student_info": returnStudent,
	})
}

//GetFriends ...
func (sCtrl StudentController) GetFriends(c *gin.Context) {
	accessDes, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get friend list",
		})
		c.Abort()
		return
	}

	returnFrList, err := studenMod.GetFriends(int(accessDes.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get friend list",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Get list friends",
		"friend_list": returnFrList,
	})
}

func (sCtrl StudentController) GetDormMoney(c *gin.Context) {
	accessDes, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Can not extract metadata from token",
			"error":   err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get dorm money",
		})
		c.Abort()
		return
	}

	returnMoneyList, err := studenMod.GetDormMoney(int(accessDes.UserID))
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Can not get money list from db",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Dorm money history",
		"money_list": returnMoneyList,
	})
}
