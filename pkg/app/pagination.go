package app

import (
	"my-project-admin-service/global"
	"my-project-admin-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := utils.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := utils.StrTo(c.Query("pageSize")).MustInt()
	if pageSize <= 0 {
		return global.AppSettings.DefaultPageSize
	}
	if pageSize > global.AppSettings.MaxPageSize {
		return global.AppSettings.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
