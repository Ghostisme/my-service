package service

import (
	"my-service/internal/model"
)

// form中的内容表示该参数对应的key值，binding中required表示必填参数
type LoginRequest struct {
	UserName string `form:"username" binding:"required,max=100"`
	Password string `form:"password" binding:"required,max=100"`
}

// func (svc *Service) GetUserByName(userName string) (uint32, error) {
// 	id, err := svc.dao.GetUserByName(userName)
// 	if err != nil {
// 		return 0, err
// 	}

// 	if id > 0 {
// 		return id, nil
// 	}

// 	return 0, errors.New("user does not exist")
// }

// func (svc *Service) GetUser(userName, Password string) (*model.User, error) {
// 	// user, err := svc.dao.GetUser(userName, Password)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// if id > 0 {
// 	// 	return id, nil
// 	// }

// 	return svc.dao.GetUser(userName, Password)
// }

func (svc *Service) Login(userName, Password string) (*model.User, error) {
	return svc.dao.Login(userName, Password)
}
