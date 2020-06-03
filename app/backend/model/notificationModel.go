package model

import (
	"DormAppBackend/db"
	"DormAppBackend/forms"

	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	Title   string `json:"title" gorm:"type:text; not null"`
	Content string `json:"content" gorm:"type:text; not null"`
}

func (noti Notification) GetAllNotification() ([]Notification, error) {
	var listNoti []Notification

	rows, err := db.GetDB().Table("notification").Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var newNoti Notification
		db.GetDB().ScanRows(rows, &newNoti)
		listNoti = append(listNoti, newNoti)
	}

	return listNoti, nil
}

func (noti Notification) CreateNotification(notiForm forms.NotificationForm) (*Notification, error) {
	newNoti := &Notification{
		Title:   notiForm.Title,
		Content: notiForm.Content,
	}

	err := db.GetDB().Create(newNoti).Error
	if err != nil {
		return nil, err
	}

	err = db.GetDB().Table("notification").Find(&newNoti).Error
	if err != nil {
		return nil, err
	}

	return newNoti, nil
}
