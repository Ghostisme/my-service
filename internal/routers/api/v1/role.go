package v1

import (
	"fmt"
	"my-service/internal/service"
	"my-service/pkg/app"
	"my-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Role struct{}

func NewRole() Role {
	return Role{}
}

// @Summer 获取角色列表
// @Produce json
// @Param token header string true "token"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Param keyWord query string false "关键词"
// @Param status query int true "是否有效"
// @Param page query int true "页码"
// @Param pageSize query int true "每页条数"
// @Success 200 {object} model.RoleInfo "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/role [post]
func (r Role) List(c *gin.Context) {
	param := service.RoleListRequest{}
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
	total, err := svc.RoleListCount(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorRoleCountFail)
		return
	}
	roleList, err := svc.RoleList(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorRoleListFail)
		return
	}
	response.ToResponseList(roleList, total)
}

// @Summer 编辑角色信息
// @Produce json
// @Param token header string true "token"
// @Param id query int true "角色主键id"
// @Param name query string false "角色姓名"
// @Success 200 {object} model.SwaggerSuccess "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/role [put]
func (r Role) Update(c *gin.Context) {
	// var (
	// 	token   string
	// 	errCode = errcode.Success
	// )
	// var token string
	param := service.RoleUpdateRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	// 通过jwt的token获取userId
	token := c.GetHeader("token")
	claims, _ := app.ParseToken(token)
	// res, err := svc. claims.UserId
	res, err := svc.UpdateRole(claims.UserId, &param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorRoleUpdateFail)
		return
	}
	fmt.Println("编辑角色%v", res)
	response.ToResponseSuccess()
}
