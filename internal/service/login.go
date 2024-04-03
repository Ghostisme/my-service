package service

import (
	"my-service/internal/model"
)

// form中的内容表示该参数对应的key值，binding中required表示必填参数
type LoginRequest struct {
	UserName string `form:"username" binding:"required,max=100"`
	Password string `form:"password" binding:"required,max=100"`
}

// 登录
func (svc *Service) Login(userName, Password string) (*model.User, error) {
	return svc.dao.Login(userName, Password)
}

// 注册
func (svc *Service) Register(userName, Password string) (int, error) {
	return svc.dao.Register(userName, Password)
}
