package model

import (
	"DormAppBackend/forms"
	"errors"
	"log"
	"strings"

	"DormAppBackend/db"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var (
	authModel = new(AuthModel)
)

// User user type
type User struct {
	gorm.Model
	Role     int    `gorm:"not null; default:0" json:"role"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"type:text;not null" json:"password"`
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

//Login check if user login with correct information
func (u User) Login(form forms.LoginForm) (user User, token Token, err error) {
	err = db.GetDB().Where(&User{
		Username: form.Username,
	}).Find(&user).Error

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))

	if err != nil {
		return user, token, errors.New("Invalid Password")
	}

	//Generate JWT token
	tokenDetails, err := authModel.CreateToken(user.ID, uint(user.Role))
	if err != nil {
		return user, token, errors.New("Can not create token")
	}

	//If save to redis success return token to the user
	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

//Register receive data and create new user in db
func (u *User) Register(form forms.RegisterForm) (*User, error) {
	var err error
	db := db.GetDB()
	// defer db.Close()

	//check space in form
	if checkInputFormat(form.Username) || strings.Contains(form.Password, " ") {
		return nil, errors.New("Invalid format")
	}

	newUser := &User{
		Username: form.Username,
		Password: form.Password,
	}

	var checkUser = new(User)
	if !db.Where("username = ?", form.Username).Find(&checkUser).RecordNotFound() {
		return nil, errors.New("User existed")
	}

	newUser.Password, err = hashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}

	check := db.NewRecord(newUser)
	if check != true {
		return nil, errors.New("User existed")
	}

	db.Create(newUser)

	return newUser, err
}

func checkInputFormat(inputString string) bool {
	checkList := []string{" ", "*", "#", "/", "\\"}
	for _, check := range checkList {
		if strings.Contains(inputString, check) {
			return true
		}
	}
	return false
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
