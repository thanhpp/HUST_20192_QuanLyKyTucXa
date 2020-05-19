package model

import (
	"DormAppBackend/db"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	Role         int64
}

//AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	Role       int64
	UserID     int64
}

//Token Token receive from user
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//CreateToken use user ID to create a token
func (authM *AuthModel) CreateToken(userID uint, role uint) (*TokenDetails, error) {

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
	atClaims["role"] = role

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	//FIXME: Get access token config
	//FIXME: log format
	tokenDetails.AccessToken, err = at.SignedString([]byte("Access_token_secret"))
	if err != nil {
		return nil, err
	}

	//Create refresh token
	rtClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	rtClaims["refresh_uuid"] = tokenDetails.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = tokenDetails.RtExpires
	rtClaims["role"] = role
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

//ExtractToken get token from request header
func (authM AuthModel) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//VerifyToken validation token
func (authM AuthModel) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := authM.ExtractToken(r)

	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method : %+v", token.Header["alg"])
			}
			return []byte("Access_token_secret"), nil
		})

	if err != nil {
		return nil, err
	}

	return token, nil
}

//TokenValid valid token of a http request
func (authM AuthModel) TokenValid(r *http.Request) error {
	token, err := authModel.VerifyToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && token.Valid {
		return err
	}

	return nil
}

//ExtractTokenMetadata get UUID from token
func (authM AuthModel) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := authM.VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		role, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["role"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
			Role:       role,
		}, nil
	}

	return nil, err
}

//GetRoleFromToken get role comes with token
func (authM AuthModel) GetRoleFromToken(r *http.Request) (int64, error) {
	token, err := authM.VerifyToken(r)
	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		role, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["role"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return role, nil
	}

	return -1, err
}

//FetchAuth get userID from UUID
func (authM AuthModel) FetchAuth(authD *AccessDetails) (int64, error) {
	userid, err := db.GetRedisClient().Get(authD.AccessUUID).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseInt(userid, 10, 64)

	return userID, nil
}

//DeleteAuth rermove auth from redis
func (authM AuthModel) DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := db.GetRedisClient().Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}
