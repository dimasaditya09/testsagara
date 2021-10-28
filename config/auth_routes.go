package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testSagara/helpers"
	"testSagara/service"
)

type LoginJson struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthRoutes(route *gin.Engine) {
	group := route.Group("/api/auth")

	group.POST("/login", func(c *gin.Context) {
		var postLogin LoginJson
		err := c.ShouldBind(&postLogin)
		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		code := http.StatusOK
		response := service.Login(postLogin.Username, postLogin.Password, c)
		if !response.Success {
			code = http.StatusBadRequest
		}
		c.JSON(code, response)
	})
}
