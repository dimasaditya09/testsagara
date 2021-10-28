package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"os"
	"testSagara/utils"
	"time"
)

type LoginStruct struct {
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

func Login(username, password string, c *gin.Context) utils.Response {
	if username != "admin" {
		return utils.Response{Success: false, Message: "username not found", Data: d}
	}
	bytePass := []byte("$2a$14$z4D8eJtyo5RPjuINOIidRueDs45oCpiX4N32/a6iRl4mNW8Kr59Ya")
	checkPassword := bcrypt.CompareHashAndPassword(bytePass, []byte(password))
	if checkPassword != nil {
		return utils.Response{Success: false, Message: checkPassword.Error(), Data: d}
	}
	atClaims := jwt.MapClaims{}
	atClaims["userName"] = username
	atClaims["platform"] = "web"
	atClaims["exp"] = time.Now().Add(time.Minute * 60 * 60 * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, errToken := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if errToken != nil {
		return utils.Response{Success: false, Message: errToken.Error(), Data: d}
	}
	dataResponse := LoginStruct{
		UserName: username,
		Token:    token,
	}
	return utils.Response{Success: true, Data: dataResponse, Message: "Data Login"}
}
