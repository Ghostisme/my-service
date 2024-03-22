package app

import (
	"my-service/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

/**
 * @brief 创建回复实例
 * @param ctx-网络请求上下文
 * @return 回复实例
 */
func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

/**
 * @brief 回复列表
 * @param list-列表
 * @param total-总条数
 */
func (r *Response) ToResponseList(list interface{}, total int) {
	err := errcode.Success
	r.Ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list": list,
			"page": Pager{
				Page:     GetPage(r.Ctx),
				PageSize: GetPageSize(r.Ctx),
				Total:    total,
			},
		},
		"code": err.Code(),
		"msg":  err.Msg(),
	})
}

/**
 * @brief 回复列表
 * @param list-列表
 */
func (r *Response) ToResponseListWithoutPager(list interface{}) {
	err := errcode.Success
	r.Ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list": list,
		},
		"code": err.Code(),
		"msg":  err.Msg(),
	})
}

/**
 * @brief 错误回复
 * @param err-错误信息
 */
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(http.StatusOK, response)
}

/**
 * @brief 回复成功
 * @param res-
 */
func (r *Response) ToResponse(res interface{}) {
	err := errcode.Success
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
		"data": res,
	}
	r.Ctx.JSON(http.StatusOK, response)
}

/**
 * @brief 回复成功
 */
func (r *Response) ToResponseSuccess() {
	err := errcode.Success
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}
	r.Ctx.JSON(http.StatusOK, response)
}

/**
 * @brief 回复下载
 * @param fileName-文件名
 */
func (r *Response) ToResponseDownload(fileName string) {
	fileContentDisposition := "attachment;filename=\"" + fileName + "\""
	r.Ctx.Header("Content-Type", "application/zip")
	r.Ctx.Header("Content-Disposition", fileContentDisposition)
	r.Ctx.Header("Cache-Control", "no-cache")
	r.Ctx.File("./" + fileName)
}
