package v1

import (
	"my-service/internal/service"
	"my-service/pkg/app"
	"my-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// @Summer 获取用户列表
// @Produce json
// @Param beginTime "开始时间"
// @Param endTime "结束时间"
// @Param keyWord "关键词"
// @Param status "是否有效"
// @Param page "页码"
// @Param pageSize ""条数
// @Success 200 {object} model.UserInfo "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/userlist [post]

func (u User) List(c *gin.Context) {
	param := service.UserListRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	total, err := svc.ListCount(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetUserCountFail)
		return
	}
	userList, err := svc.List(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetUserListFail)
		return
	}
	response.ToResponseList(userList, total)
}
