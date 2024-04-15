package model

import (
	"my-service/global"

	"github.com/jinzhu/gorm"
)

type tokenInfo struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type TokenInfoSwagger struct {
	*SwaggerSuccess
	Data *tokenInfo `json:"data"`
}

func (u User) TableName() string {
	return "user"
}

// func GetUserByName(db *gorm.DB, userName string) (uint32, error) {
// 	var user User

// 	db = db.Where("username = ? and status = 1", userName)
// 	err := db.First(&user).Error

// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return 0, err
// 	}

// 	return user.Model.ID, nil
// }

// 用户登录
func Login(db *gorm.DB, userName, Password string) (*User, error) {
	var user User
	db = db.Select("`user`.*, role.*").Preload("Role").Joins("LEFT JOIN `role` ON role.id = `user`.role_id").Where("`user`.username = ? and `user`.password = ? and `user`.status = 1", userName, Password)
	err := db.Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

// 用户名注册
func RegisterUser(db *gorm.DB, userName, password string) (int, error) {
	user := User{}
	res, err := user.Create(db, NewOption(WithUserRegister(userName, password)))
	global.ApiLogger.Info("当前注册调用新增", res)
	// err = db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return 0, nil
}

// 手机号注册
func RegisterMobile(db *gorm.DB, Mobile string) (int, error) {
	user := User{}
	res, err := user.Create(db, NewOption(WithMobileRegister(Mobile)))
	global.ApiLogger.Info("当前注册调用新增", res)
	// err = db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return 0, nil
}
