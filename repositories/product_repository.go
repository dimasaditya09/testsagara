package repositories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"testSagara/models"
	"testSagara/utils"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) ProductPagination(pagination *utils.Pagination, c *gin.Context) RepositoryResult {
	var products []models.Product
	var totalRow int64 = 0

	queryGroup := r.db
	if pagination.Keyword != "" {
		queryGroup = queryGroup.Where(models.Product{ProductName: pagination.Keyword}).Or(models.Product{ProductId: pagination.Keyword})
	}
	countQuery := queryGroup
	errCount := countQuery.Model(&models.Product{}).Count(&totalRow).Error
	if errCount != nil {
		return RepositoryResult{Error: errCount}
	}
	offset := (pagination.Page - 1) * pagination.Limit
	var sort = pagination.Sort
	if sort != "" && strings.ToUpper(sort) == "ASC" {
		sort = "created_at asc"
	} else {
		sort = "created_at desc"
	}

	errFind := queryGroup.Limit(pagination.Limit).Offset(offset).Order(sort).Find(&products).Error
	if errFind != nil {
		return RepositoryResult{Error: errFind}
	}
	pagination.Items = products
	pagination.TotalRows = totalRow
	return RepositoryResult{Result: pagination}
}

func (r *ProductRepository) DetailProduct(id string) RepositoryResult {
	var product models.Product
	err := r.db.Where(&models.Product{ProductId: id}).Take(&product).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: &product}
}

func (r *ProductRepository) CreateProduct(dataProduct *models.Product) RepositoryResult {
	err := r.db.Create(dataProduct).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: dataProduct}
}

func (r *ProductRepository) UpdateProduct(id string, product *models.Product) RepositoryResult {
	var products models.Product
	err := r.db.Where(&models.Product{ProductId: id}).Updates(&models.Product{
		ProductName:  product.ProductName,
		ProductImage: product.ProductImage,
	}).First(&products).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &products}
}

func (r *ProductRepository) DeleteProduct(id string) RepositoryResult {
	err := r.db.Delete(&models.Product{ProductId: id}).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: nil}
}
