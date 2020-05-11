package model

import (
	"errors"
	"log"

	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/db"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User user type
type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"type:text;not null"`
}

func hashPassword(password string) (string, error) {
	var err error
	hashedPwd := []byte(password)

	hashedPwd, err = bcrypt.GenerateFromPassword(hashedPwd, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("hashPassword error : %+v", err)
		return "", err
	}

	return string(hashedPwd), nil
}

//Register receive data and create new user in db
func (u *User) Register() error {
	var err error
	db := db.GetDB()
	// defer db.Close()

	u.Password, err = hashPassword(u.Password)
	if err != nil {
		return err
	}

	check := db.NewRecord(u)
	if check != true {
		return errors.New("User existed")
	}

	db.Create(u)

	return err
}

//CheckPass check if input password is in the db
func (u User) CheckPass() (bool, error) {
	db := db.GetDB()
	// defer db.Close()

	checkUser := new(User)

	err := db.Where(&User{
		Username: u.Username,
	}).Find(&checkUser).Error

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(u.Password))
	if err != nil {
		return false, err
	}

	return true, err
}

//GetByUsr query user by username
func (u *User) GetByUsr(username string) error {
	db := db.GetDB()
	// defer db.Close()

	var err error

	err = db.Where(&User{
		Username: username,
	}).Find(&u).Error

	if err != nil {
		return err
	}

	return err
}
