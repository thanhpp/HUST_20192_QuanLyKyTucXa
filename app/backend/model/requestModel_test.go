package model

import (
	"DormAppBackend/config"
	"DormAppBackend/db"
	"DormAppBackend/forms"
	"DormAppBackend/tlog"
	"testing"
)

func TestNewRequest(t *testing.T) {
	config.Init()
	db.Init()
	tlog.Init()
	var r Request
	rqForm := &forms.NewRequestForm{
		Title:   "Test title",
		Message: "Test request",
	}

	rq := &Request{
		Status:    "new request 123",
		StudentID: 1,
		Title:     rqForm.Title,
		Message:   rqForm.Message,
	}

	checkrq := new(Request)
	err := db.GetDB().Table("request").Where(&rq).Find(checkrq).Error
	if err == nil {
		t.Error(err, checkrq)
		return

	}

	err, _ = r.NewRequest(*rqForm, 1)
	if err != nil {
		t.Error(err)
	}
}
