package service

import (
	"github.com/google/uuid"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"testSagara/models"
	"testSagara/repositories"
	"testSagara/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var d struct{}

func ProductPagination(repoProduct repositories.ProductRepository, c *gin.Context, pagination *utils.Pagination) utils.Response {
	operationResult := repoProduct.ProductPagination(pagination, c)
	if operationResult.Error != nil {
		return utils.Response{Success: false, Message: operationResult.Error.Error(), Data: d}
	}
	return utils.Response{Success: true, Message: "OK", Data: operationResult.Result}
}

func DetailProduct(id string, repoProduct repositories.ProductRepository) utils.Response {
	operationResult := repoProduct.DetailProduct(id)
	if operationResult.Error != nil {
		return utils.Response{Success: false, Message: operationResult.Error.Error(), Data: d}
	}
	return utils.Response{Success: true, Message: "OK", Data: operationResult.Result}
}

func CreateProduct(product *models.Product, repoProduct repositories.ProductRepository) utils.Response {
	product.ProductId = uuid.New().String()
	operationResult := repoProduct.CreateProduct(product)
	if operationResult.Error != nil {
		return utils.Response{Success: false, Message: operationResult.Error.Error(), Data: d}
	}
	return utils.Response{Success: true, Message: "OK", Data: operationResult.Result}
}

func UpdateProduct(id string, product *models.Product, repoProduct repositories.ProductRepository) utils.Response {
	operationResult := repoProduct.UpdateProduct(id, product)
	if operationResult.Error != nil {
		return utils.Response{Success: false, Message: operationResult.Error.Error(), Data: d}
	}
	return utils.Response{Success: true, Message: "OK", Data: operationResult.Result}
}

func DeleteProduct(id string, repoProduct repositories.ProductRepository) utils.Response {
	operationResult := repoProduct.DeleteProduct(id)
	if operationResult.Error != nil {
		return utils.Response{Success: false, Message: operationResult.Error.Error()}
	}
	return utils.Response{Success: true, Message: "OK", Data: operationResult.Result}
}

func UploadFiles(file *multipart.FileHeader, c *gin.Context, bucket string) utils.Response {
	path := os.Getenv("UPLOAD_PATH") + bucket
	os.Mkdir(path, 0777)
	ext := filepath.Ext(filepath.Base(file.Filename))
	filename := strings.ReplaceAll(filepath.Base(file.Filename), ext, "")
	filename = filename + "_" + time.Now().Format("20060102150405345")
	filename = strings.ReplaceAll(filename, " ", "_") + ext

	if err := c.SaveUploadedFile(file, path+"/"+filename); err != nil {
		return utils.Response{Success: false, Data: d, Message: err.Error()}
	}
	urlFile := os.Getenv("APP_URL") + "media/" + bucket + "/" + filename
	return utils.Response{Success: true, Data: urlFile, Message: "Ok"}
}
