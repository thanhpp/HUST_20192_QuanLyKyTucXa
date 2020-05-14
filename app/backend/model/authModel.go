package model

import (
	"DormAppBackend/db"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/twinj/uuid"
)

// AuthModel auth
type AuthModel struct{}

//TokenDetails contains details for an user token
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

//AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     int64
}

//Token Token receive from user
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//CreateToken use user ID to create a token
func (authM *AuthModel) CreateToken(userID uint) (*TokenDetails, error) {

	tokenDetails := &TokenDetails{}
	tokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenDetails.AccessUUID = uuid.NewV4().String()

	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUUID = uuid.NewV4().String()

	var err error
	//Create access token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = tokenDetails.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = tokenDetails.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	//FIXME: Get access token config
	//FIXME: log format
	tokenDetails.AccessToken, err = at.SignedString([]byte("Access_token_secret"))
	if err != nil {
		return nil, err
	}

	//Create refresh token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokenDetails.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = tokenDetails.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	//FIXME: Get refresh token config
	//FIXME: log format
	tokenDetails.RefreshToken, err = rt.SignedString([]byte("Refresh_token_secret"))
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}

//CreateAuth save token to db
func (authM AuthModel) CreateAuth(userID uint, tokenDetails *TokenDetails) error {
	at := time.Unix(tokenDetails.AtExpires, 0)
	rt := time.Unix(tokenDetails.RtExpires, 0)
	now := time.Now()

	redisClient := db.GetRedisClient()

	errAccess := redisClient.Set(tokenDetails.AccessUUID, strconv.Itoa(int(userID)), at.Sub(now)).Err()
	//TODO: log format
	if errAccess != nil {
		return errAccess
	}

	errRefresh := redisClient.Set(tokenDetails.RefreshUUID, strconv.Itoa(int(userID)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
