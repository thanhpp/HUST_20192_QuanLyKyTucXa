package model

import (
	"DormAppBackend/db"
	"DormAppBackend/forms"
	"DormAppBackend/tlog"
	"errors"

	"github.com/jinzhu/gorm"
)

type Request struct {
	gorm.Model
	Status    string `gorm:"type:text; not null" json:"status"`
	StudentID int    `gorm:"type:int; not null" json:"student_id"`
	Title     string `gorm:"type:text; not null" json:"title"`
	Message   string `gorm:"type:text; not null" json:"message"`
	Reply     string `gorm:"type:text" json:"reply"`
	Note      string `gorm:"type:text" json:"note"`
}

func (r Request) NewRequest(rqForm forms.NewRequestForm, studentID int) (error, *Request) {
	rq := &Request{
		Status:    "new request",
		StudentID: studentID,
		Title:     rqForm.Title,
		Message:   rqForm.Message,
	}

	err := db.GetDB().Table("request").Where(&rq).Find(&rq).Error
	if err == nil {
		return errors.New("request existed"), nil
	}

	err = db.GetDB().Table("request").Create(rq).Error
	if err != nil {
		tlog.Error("Can not create new request", err)
		return err, nil
	}

	return nil, rq
}

func (r Request) GetListRequestStudentID(studentID int) ([]Request, error) {
	var listRq []Request
	rows, err := db.GetDB().Table("request").Where("student_id = ?", studentID).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rq Request
		db.GetDB().ScanRows(rows, &rq)
		listRq = append(listRq, rq)
	}

	return listRq, nil
}

func (r Request) ReplyRequest(repForm forms.ReplyRequestForm) (Request, error) {
	var rq Request

	err := db.GetDB().Table("request").Where("id = ?", repForm.RequestID).Find(&rq).Error
	if err != nil {
		return rq, err
	}

	rq.Reply = repForm.Reply
	rq.Status = "replied"

	err = db.GetDB().Table("request").Save(&rq).Error
	if err != nil {
		return rq, err
	}

	return rq, nil
}

func (r Request) GetAllRequest() ([]Request, error) {
	var listRq []Request

	rows, err := db.GetDB().Table("request").Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rq Request
		db.GetDB().ScanRows(rows, &rq)
		listRq = append(listRq, rq)
	}

	return listRq, nil
}
