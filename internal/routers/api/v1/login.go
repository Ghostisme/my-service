package v1

import (
	"fmt"
	"my-project-admin-service/internal/service"
	"my-project-admin-service/pkg/app"
	"my-project-admin-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// @Summer 登录验证
// @Produce json
// @Param userName query string true "用户名" maxlength(100)
// @Param Password query string true "用户密码" maxlength(100)
// @Success 200 {object} model.TokenInfoSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/login [post]
func GetUser(c *gin.Context) {
	param := service.LoginRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	_, err := svc.GetUserByName(param.UserName)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserNotExist)
		return
	}

	id, err := svc.GetUser(param.UserName, param.Password)
	fmt.Printf("id:%d", id)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserPasswordError)
		return
	}
	token := "asdasddsaadssadsda"
	// token, err := app.GenerateToken()
	// fmt.Printf("当前token%v", token)
	// if err != nil {
	// 	global.Logger.Errorf("app.GenerateToken err: %v", err)
	// 	response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	// 	return
	// }
	// 将生成的token插入会当前用户表
	// _, err1 := svc.SetToken(id, token)
	// if err1 != nil {
	// 	global.Logger.Errorf("app.GenerateToken err: %v", err1)
	// 	response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	// 	return
	// }
	response.ToResponse(gin.H{
		"token": token,
	})
}
// @Summer 登出
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.SwaggerSuccess "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/logout [post]
func Logout(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToResponseSuccess()
}