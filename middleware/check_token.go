package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	h "testSagara/helpers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(platform string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		payload, _ := ExtractClaims(c.Request)
		if payload["platform"] != platform {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token does not have access to endpoint"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func TokenApiMiddleware(mustLoggedIn bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		payload, _ := ExtractClaims(c.Request)
		restricted := []string{"android", "ios", "web"}
		var platform = payload["platform"].(string)
		// RESTRICTED PLATFORM
		if h.FindInArray(restricted, platform) == false {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token does not have access to endpoint"})
			c.Abort()
			return
		}

		// CHECK IF USER MUST LOGGED IN TO USE API
		if mustLoggedIn == true {

		}

		c.Next()
	}
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractClaims(r *http.Request) (jwt.MapClaims, bool) {
	tokenString := ExtractToken(r)
	hmacSecretString := os.Getenv("JWT_SECRET") // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
