package controller

import (
	"DormAppBackend/forms"
	"DormAppBackend/model"
	"DormAppBackend/tlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestController struct{}

var rqModel model.Request

func (rCtrl RequestController) NewRequest(c *gin.Context) {
	var requestForm forms.NewRequestForm

	accessDes, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		tlog.Info(tlog.Itf{
			"msg": "ExtractTokenMetadata",
			"err": err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not create request",
		})
		c.Abort()
		return
	}

	if c.ShouldBindJSON(&requestForm) != nil {
		tlog.Info(tlog.Itf{
			"msg": "BindJSON",
		})
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid form",
		})
		c.Abort()
		return
	}

	err, request := rqModel.NewRequest(requestForm, int(accessDes.UserID))
	if err != nil {
		tlog.Info(tlog.Itf{
			"msg":   "NewRequest",
			"error": err,
		})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not create request",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "New request created",
		"request": request,
	})
}

func (rCtrl RequestController) ListRequestByStudentID(c *gin.Context) {
	accessDes, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list request",
		})
		c.Abort()
		return
	}

	listReq, err := rqModel.GetListRequestStudentID(int(accessDes.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list request",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Get list request successfully",
		"list_request": listReq,
	})
}

func (rCtrl RequestController) GetAllRequest(c *gin.Context) {
	listRq, err := rqModel.GetAllRequest()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not get list requests",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Get list requests OK",
		"list_request": listRq,
	})
}

func (rCtrl RequestController) ReplyRequest(c *gin.Context) {
	var replyForm forms.ReplyRequestForm

	if c.ShouldBindJSON(&replyForm) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Can not reply request",
		})
		c.Abort()
		return
	}

	rep, err := rqModel.ReplyRequest(replyForm)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Can not reply request",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Reply request OK",
		"reply":   rep,
	})
}
