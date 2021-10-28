package helpers

import (
	"strconv"
	u "testSagara/utils"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationRequest(c *gin.Context) *u.Pagination {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	sort := c.DefaultQuery("sort", "desc")
	sortBy := c.DefaultQuery("sortBy", "created_at")
	keyword := c.DefaultQuery("keyword", "")

	return &u.Pagination{Limit: limit, Page: page, Sort: sort, SortBy: sortBy, Keyword: keyword}
}
