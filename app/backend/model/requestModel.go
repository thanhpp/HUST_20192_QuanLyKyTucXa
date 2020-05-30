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
	Message   string `gorm:"type:text; not null" json:"message"`
	Note      string `gorm:"type:text" json:"note"`
}

func (r Request) NewRequest(rqForm forms.NewRequestForm, studentID int) (error, *Request) {
	rq := &Request{
		Status:    "new request",
		StudentID: studentID,
		Message:   rqForm.Message,
	}

	if !db.GetDB().Table("request").Select("id").Where("student_id = ?", rq.StudentID).Where("message = ?", rq.Message).Where("status = ?", rq.Status).RecordNotFound() {
		return errors.New("Request existed"), nil
	}

	err := db.GetDB().Create(rq).Error
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

func (r Request) UpdateRequestStatus() {

}
