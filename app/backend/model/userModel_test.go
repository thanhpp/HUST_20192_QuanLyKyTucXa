package model

import (
	"testing"

	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/config"
	"github.com/THANHPP/HUST_20192_QuanLyKyTucXa/app/backend/db"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	testPwd := "abc"
	val, err := hashPassword(testPwd)
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
	err = bcrypt.CompareHashAndPassword([]byte(val), []byte(testPwd))
	if err != nil {
		t.Error(err)
	}
}

func TestRegister(t *testing.T) {
	config.Init()
	db.Init()

	newUser := &User{
		Username: "test 1",
		Password: "test abc",
	}
	err := newUser.Register()
	if err != nil {
		t.Error(err)
	}
}

func TestCheckPass(t *testing.T) {
	config.Init()
	db.Init()

	loginUser := &User{
		Username: "test 1",
		Password: "test abc",
	}

	check, err := loginUser.CheckPass()
	if err != nil {
		t.Errorf("TestCheckPass err : %+v", err)
	}

	if check != true {
		t.Errorf("Wrong password : %+v", check)
	}
}
