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
// @Param token header string true "token"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Param keyWord query string false "关键词"
// @Param status query int true "是否有效"
// @Param page query int true "页码"
// @Param pageSize query int true "每页条数"
// @Success 200 {object} model.UserInfo "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user [post]
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
	total, err := svc.UserListCount(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserCountFail)
		return
	}
	userList, err := svc.UserList(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserListFail)
		return
	}
	response.ToResponseList(userList, total)
}

// @Summer 创建用户
// @Produce json
// @Param token header string true "token"
// @Param username query string false "用户姓名"
// @Param mobile query string false "用户联系方式"
// @Param addr query string false "地址"
// @Param email query string false "邮箱"
// @Param status query int false "是否启用"
// @Param is_admin query int false "是否是管理员"
// @Success 200 {object} model.SwaggerSuccess "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user/save [post]
func (u User) Create(c *gin.Context) {
	param := service.UserCreateRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	_, err := svc.CreateUser(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserCreateFail)
		return
	}
	response.ToResponseSuccess()
}

// @Summer 更新用户
// @Produce json
// @Param token header string true "token"
// @Param id query int true "用户id主键"
// @Param username query string false "用户姓名"
// @Param mobile query string false "用户联系方式"
// @Param addr query string false "地址"
// @Param email query string false "邮箱"
// @Param status query int false "是否启用"
// @Success 200 {object} model.SwaggerSuccess "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user [put]
func (u User) Update(c *gin.Context) {
	param := service.UserUpdateRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	_, err := svc.UpdateUser(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserUpdateFail)
		return
	}
	response.ToResponseSuccess()
}

// @Summer 删除用户
// @Produce json
// @Param token header string true "token"
// @Param id query int true "用户id主键"
// @Success 200 {object} model.SwaggerSuccess "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user [delete]
func (u User) Del(c *gin.Context) {
	param := service.UserDelRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	_, err := svc.DelUser(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserDelFail)
		return
	}
	response.ToResponseSuccess()
}
