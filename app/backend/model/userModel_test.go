package model

import (
	"DormAppBackend/forms"
	"testing"

	"DormAppBackend/config"
	"DormAppBackend/db"

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

func TestRegister1(t *testing.T) {
	config.Init()
	db.Init()

	newUser := &User{
		Username: "test /",
		Password: "test abc",
	}

	_, err := newUser.Register(forms.RegisterForm{
		Username: newUser.Username,
		Password: newUser.Password,
	})
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
