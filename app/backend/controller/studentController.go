package controller

import (
	"DormAppBackend/forms"
	"DormAppBackend/model"
	"DormAppBackend/tlog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//StudentController ...
type StudentController struct{}

var (
	studenMod = new(model.Student)
	monMng    = new(model.MoneyManage)
)

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

func (sCtrl StudentController) GetStudentInfoLV1(c *gin.Context) {
	stdID := c.Query("id")
	if stdID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No student id found",
		})
		c.Abort()
		return
	}

	findID, err := strconv.Atoi(stdID)
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Can not parse to int",
			"error":   err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get student info",
		})
		c.Abort()
		return
	}

	returnStd, err := studenMod.GetStudentInfo(findID)
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "GetStudentInfo",
			"error":   err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get student info",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Get student info OK",
		"student_info": returnStd,
	})
}

func (sCtrl StudentController) GetStudentByRoomID(c *gin.Context) {
	roomID := c.Query("id")
	if roomID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No room id found",
		})
		c.Abort()
		return
	}

	findID, err := strconv.Atoi(roomID)
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Can not parse to int",
			"error":   err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list students",
		})
		c.Abort()
		return
	}

	returnListStd, err := studenMod.GetAllStudentsByRoom(findID)
	if err != nil {
		tlog.Info(tlog.Itf{
			"message": "Can not GetAllStudentsByRoom",
			"error":   err,
		})

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list students",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Get list student OK",
		"list_students": returnListStd,
	})
}

func (sCtrl StudentController) UpdateStudentRoom(c *gin.Context) {
	var err error

	stdID := c.Query("stdid")
	if stdID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No studentid found",
		})
		c.Abort()
		return
	}

	roomID := c.Query("roomid")
	if roomID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No room id found",
		})
		c.Abort()
		return
	}

	stdIDInt, err := strconv.Atoi(stdID)
	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid arguments",
			"error":   err,
		})
		c.Abort()
		return
	}

	std, err := studenMod.ChangeRoom(stdIDInt, roomIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not update room info for student",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update room info for student OK",
		"student": std,
	})
}

func (sCtrl StudentController) CalNewMonthMoney(c *gin.Context) {
	var err error

	listMonMong, err := monMng.CalculateNewMonth()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":           "Error while calculate dorm money",
			"list_money_manage": listMonMong,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           "Calculate new month money OK",
		"list_money_manage": listMonMong,
	})
}

func (sCtrl StudentController) GetAllMoneyManage(c *gin.Context) {
	listMonMng, err := monMng.GetAllMoneyManage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get money manage",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           "Get all money manage OK",
		"list_money_manage": listMonMng,
	})
}

func (sCtrl StudentController) UpdatePayment(c *gin.Context) {
	var updatePay forms.UpdatePaymentForm

	if c.ShouldBindJSON(&updatePay) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid post form",
		})
		c.Abort()
		return
	}

	returnMonMng, err := monMng.UpdatePaymentStatus(updatePay.MoneyManageID, updatePay.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not update payment status",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Update payment status OK",
		"money_manage": returnMonMng,
	})
}

func (sCtrl StudentController) NewStudentInfo(c *gin.Context) {
	var stdInfoForm forms.StudentInfoForm

	accessDes, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not find user",
		})
		c.Abort()
		return
	}

	if c.ShouldBindJSON(&stdInfoForm) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid student info form",
		})
		c.Abort()
		return
	}

	std, err := studenMod.NewStudentInfto(int(accessDes.UserID), stdInfoForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not create new student info",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update student info OK",
		"student": std,
	})
}

func (sCtrl StudentController) GetAllStudent(c *gin.Context) {
	listStd, err := studenMod.GetAllStudent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list students",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Get list student OK",
		"list_student": listStd,
	})
}

func (sCtrl StudentController) GetPaymentReport(c *gin.Context) {
	paid, unpaid, err := monMng.GetReportStudentPayment()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get payment report",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get payment report OK",
		"paid":    paid,
		"unpaid":  unpaid,
	})
}
