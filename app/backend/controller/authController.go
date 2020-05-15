package controller

import (
	"DormAppBackend/forms"
	"DormAppBackend/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

//AuthController .....
type AuthController struct{}

var authModel = new(model.AuthModel)

//TokenValid token validation
func (authCtrl AuthController) TokenValid(c *gin.Context) {
	err := authModel.TokenValid(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization, please login again",
		})
		c.Abort()
		return
	}
}

//Refresh renew the token
func (authCtrl AuthController) Refresh(c *gin.Context) {
	var tokenForm forms.Token

	if c.ShouldBindJSON(&tokenForm) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid form",
			"form":    tokenForm,
		})
		c.Abort()
		return
	}

	//verify the token
	token, err := jwt.Parse(tokenForm.RefreshToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method : %+v", token.Header["alg"])
			}
			return []byte("Refresh_token_secret"), nil
		})

	// Token exipred
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization, please login",
		})
		return
	}

	// Check token valid
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization, please login",
		})
		return
	}

	//If valid, get the UUID
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization, please login",
			})
			return
		}

		userID, err := strconv.ParseInt((fmt.Sprintf("%f", claims["user_id"])), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization, please login",
			})
			return
		}

		//Remove the previous Refresh Token
		deleted, delErr := authModel.DeleteAuth(refreshUUID)
		if delErr != nil || deleted == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization, please login",
			})
			return
		}

		//Create new refresh Token and Access Token
		ts, createErr := authModel.CreateToken(uint(userID))
		if createErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization, please login",
			})
			return
		}

		//Save token pair to redis
		saveErr := authModel.CreateAuth(uint(userID), ts)
		if saveErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization, please login",
			})
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusOK, tokens)

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization, please login",
		})
	}

}
