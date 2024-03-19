package v1

import (
	"my-service/global"
	"my-service/internal/service"
	"my-service/pkg/app"
	"my-service/pkg/cryptor"
	"my-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// @Summer 登录验证
// @Produce json
// @Param token header string true "token"
// @Param userName query string true "用户名" maxlength(100)
// @Param Password query string true "用户密码" maxlength(100)
// @Success 200 {object} model.TokenInfoSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/login [post]
func Login(c *gin.Context) {
	param := service.LoginRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	var webSecretKey = "qgajvd17wljhaicq"
	// 根据前端秘钥进行解密
	decrypted := cryptor.AesSimpleDecrypt("2jLdyu6zzDZKHJgREyYsEw==", webSecretKey)
	// 将解密秘钥结合后端key加密
	var serviceSecretKey = "mxalxjzj9oeffag9"
	password := cryptor.AesSimpleEncrypt(decrypted, serviceSecretKey)
	// fmt.Println("数据库结果", password)
	user, err := svc.Login(param.UserName, password)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserPasswordError)
		return
	}
	token, err := app.GenerateToken(user.ID)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
		"user":  user,
	})
}

// @Summer 登出
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.SwaggerSuccess "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/logout [get]
func Logout(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToResponseSuccess()
}

// @Summer 注册
// @Produce json
// @Param token header string true "token"
// @Param username query string true "用户名" maxlength(100)
// @Param password query string true "密码" maxlength(100)
// @Param code query string true "验证码" maxlength(6)
func Register(c *gin.Context)  {
	// response := app.NewResponse(c)

	// valid, errs := app.BindAndValid(c, &param)
	// if !valid {
	// 	response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	// 	return
	// }
}

// 获取验证码
func GenValidateCode(c *gin.Context) {
	// response := app.NewResponse(c)

	// valid, errs := app.BindAndValid(c, &param)
	// if !valid {
	// 	response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	// 	return
	// }
}