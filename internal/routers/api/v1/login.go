package v1

import (
	"my-service/global"
	"my-service/internal/model"
	"my-service/internal/service"
	"my-service/pkg/app"
	"my-service/pkg/cryptor"
	"my-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

var webSecretKey = "qgajvd17wljhaicq"
var serviceSecretKey = "mxalxjzj9oeffag9"
var from = "register"

type CodeRequest struct {
	Mobile string `form:"mobile" binding:"required,max=11"`
}

type code struct {
	Code string `json:"code"`
}

type CodeResponse struct {
	*model.SwaggerSuccess
	Data *code `json:"data"`
}

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

	// 根据前端秘钥进行解密
	decrypted := cryptor.AesSimpleDecrypt(param.Password, webSecretKey)
	// 将解密秘钥结合后端key加密
	password := cryptor.AesSimpleEncrypt(decrypted, serviceSecretKey)
	// fmt.Println("数据库结果", password)
	user, err := svc.Login(param.UserName, password)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserPasswordError)
		return
	}
	token, err := app.GenerateToken(user.ID)
	if err != nil {
		global.ApiLogger.Errorf("app.GenerateToken err: %v", err)
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
// @Param username query string true "用户名" maxlength(100)
// @Param password query string true "密码" maxlength(100)
// @Param code query string true "验证码" maxlength(6)
// @Success 200 {object} model.SwaggerSuccess "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/register [post]
func Register(c *gin.Context) {
	param := service.RegisterRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.VerifyCaptchaCode(param.UserName, param.Code, from)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorValidCodeFail)
		return
	}

}

// @Summer 获取验证码
// @Produce json
// @Param mobile query string true "手机号" maxlength(11)
// @Success 200 {object} CodeResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/register [post]
// func GenValidateCode(c *gin.Context) {
func CreateCode(c *gin.Context) {
	codeLen := 6

	captchaTTL := 300
	param := CodeRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	code, b64s, err := svc.CreateCaptcha(param.Mobile, from, captchaTTL, codeLen)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateCodeFail)
		return
	}
	// err = svc.VerifyCaptchaCode(param.Mobile, code, from)
	// if err != nil {
	// 	response.ToErrorResponse(errcode.ErrorValidCodeFail)
	// 	return
	// }
	response.ToResponse(gin.H{
		"code": code,
		"img":  b64s,
	})
}
