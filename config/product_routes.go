package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testSagara/helpers"
	"testSagara/middleware"
	"testSagara/models"
	"testSagara/repositories"
	"testSagara/service"
)

func ProductRoutes(productRepository *repositories.ProductRepository, route *gin.Engine) {
	url := "/api/product"

	group := route.Group(url)
	group.Use(middleware.TokenApiMiddleware(true))

	group.GET("/fetch", func(c *gin.Context) {
		statusCode := http.StatusOK
		pagination := helpers.GeneratePaginationRequest(c)
		response := service.ProductPagination(*productRepository, c, pagination)
		if !response.Success {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, response)
	})

	group.GET("/detail/:id", func(c *gin.Context) {
		statusCode := http.StatusOK
		id := c.Param("id")
		response := service.DetailProduct(id, *productRepository)
		if !response.Success {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, response)
	})

	group.POST("/create", func(c *gin.Context) {
		var product models.Product
		statusCode := http.StatusOK

		err := c.ShouldBindJSON(&product)
		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		response := service.CreateProduct(&product, *productRepository)
		if !response.Success {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, response)
	})

	group.PUT("/update/:productId", func(c *gin.Context) {
		productId := c.Param("productId")
		var product models.Product
		statusCode := http.StatusOK

		err := c.ShouldBindJSON(&product)
		if err != nil {
			response := helpers.GenerateValidationResponse(err)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		response := service.UpdateProduct(productId, &product, *productRepository)
		if !response.Success {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, response)
	})

	group.DELETE("/delete/:productId", func(c *gin.Context) {
		productId := c.Param("productId")
		statusCode := http.StatusOK
		response := service.DeleteProduct(productId, *productRepository)
		if !response.Success {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, response)
	})

	group.POST("/upload-file", func(c *gin.Context) {
		statusCode := http.StatusOK
		bucket := c.PostForm("bucket")
		if bucket == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Bucket is required",
			})
			return
		}
		file, errFile := c.FormFile("file")
		if errFile != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": errFile.Error(),
			})
			return
		}
		response := service.UploadFiles(file, c, bucket)
		if !response.Success {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, response)

	})
}
