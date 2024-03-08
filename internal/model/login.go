package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	*Model
	UserName     string `json:"username" gorm:"Column:username"`
	Password string `json:"password" gorm:"Column:password"`
}

type tokenInfo struct {
	Token string `json:"token"`
}

type TokenInfoSwagger struct {
	*SwaggerSuccess
	Data *tokenInfo `json:"data"`
}

func (u User) TableName() string {
	return "user"
}

func GetUserByName(db *gorm.DB, userName string) (uint32, error) {
	var user User

	db = db.Where("username = ? and status = 1", userName)
	err := db.First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return user.Model.ID, nil
}

func GetUser(db *gorm.DB, userName, Password string) (uint32, error) {
	var user User

	db = db.Where("username = ? and password = ? and status = 1", userName, Password)
	err := db.First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return user.Model.ID, nil
}

func SetCurrentToken(db *gorm.DB, id uint32, token string) (uint32, error) {
	db = db.Where("id = ?", id)
	// 更新用户
	err := db.Update(&tokenInfo{token}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return 0, nil
}