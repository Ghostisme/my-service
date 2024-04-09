package service

import (
	"errors"
	"fmt"
	"my-service/global"
	"my-service/internal/model"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
)

// form中的内容表示该参数对应的key值，binding中required表示必填参数
type LoginRequest struct {
	UserName string `form:"username" binding:"required,max=100"`
	Password string `form:"password" binding:"required,max=100"`
}
type RegisterRequest struct {
	UserName string `form:"username" binding:"required,max=100"`
	Password string `form:"password" binding:"required,max=100"`
	Code     string `form:"code" binding:"required,max=6"`
}

var (
	store           = base64Captcha.DefaultMemStore
	errCodeNotExist = errors.New("code not exist")
	errCodeNotMatch = errors.New("code not match")
	// webSecretKey     = "qgajvd17wljhaicq"
	// serviceSecretKey = "mxalxjzj9oeffag9"
)

// 登录
func (svc *Service) Login(userName, Password string) (*model.User, error) {
	return svc.dao.Login(userName, Password)
}

// 注册
func (svc *Service) Register(userName, Password string) (int, error) {
	return svc.dao.Register(userName, Password)
}

// 创建验证码
func (svc *Service) CreateCaptcha(tag, from string, ttl, codeLen int) (string, string, error) {
	key := GenCaptchaCodeKey(tag, "register")
	code, b64s, err := MakeCaptcha(codeLen)
	if err != nil {
		return "", "", err
	}
	err = global.RedisClient.SetNX(global.Ctx, key, code, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		global.ServiceLogger.Info("创建验证码错误信息", err)
		return "", "", err
	}
	return code, b64s, err
}

// 获取验证码
func MakeCaptcha(codeLen int) (string, string, error) {
	//定义一个driver
	var driver base64Captcha.Driver
	driverDigit := &base64Captcha.DriverDigit{
		Height:   80,  //高度
		Width:    240, //宽度
		MaxSkew:  0.7,
		Length:   codeLen, //数字个数
		DotCount: 80,
	}
	driver = driverDigit
	//生成验证码
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := c.Generate()
	global.ServiceLogger.Info("这是什么", answer)
	global.ServiceLogger.Info("这个id是什么", id)
	code := GetCodeById(id)
	return code, b64s, err
}

// tag:唯一标记，如phone username等
// from: 标记是哪个业务申请的验证码
func GenCaptchaCodeKey(tag, from string) string {
	return "CAPTCHA-CODE-" + from + "-" + tag
}

func CreateCaptcha(tag, from string, ttl, codeLen int) (string, string, error) {
	key := GenCaptchaCodeKey(tag, "login")
	code, b64s, err := MakeCaptcha(codeLen)
	if err != nil {
		return "", "", err
	}

	// 如果code没有过期，是不允许再发送的
	err = global.RedisClient.Set(global.Ctx, key, code, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		fmt.Printf("SendSmsCode AliyunSendSms fail %v\n", err)
		return "", "", err
	}

	return code, b64s, err
}

func GetCodeById(id string) string {
	return store.Get(id, true)
}

func (svc *Service) VerifyCaptchaCode(tag, inputCode, from string) error {
	key := GenCaptchaCodeKey(tag, from)
	code, err := global.RedisClient.Get(global.Ctx, key).Result()
	global.ServiceLogger.Info("查看当前的code", code)
	if err != nil {
		if err == redis.Nil {
			return errCodeNotExist
		}
		return err
	}

	// 对比后马上删除
	err = global.RedisClient.Del(global.Ctx, key).Err()
	if err != nil {
		fmt.Printf("redis del fail %v\n", err)
		return err
	}
	// 针对输入的inputCode解密
	// requestCode := cryptor.AesSimpleDecrypt(param.Password, webSecretKey)
	if inputCode != code {
		return errCodeNotMatch
	}

	return nil
}
