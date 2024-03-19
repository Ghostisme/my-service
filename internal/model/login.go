package model

import (
	"fmt"
	"my-service/pkg/utils"

	"github.com/jinzhu/gorm"
)

type Role struct {
	*Model
	RoleCode   string `json:"role" gorm:"Column:role_code"`
	RoleName   string `json:"role_name" gorm:"Column:role_name"`
	CreateTime string `json:"create_time" gorm:"Column:create_time"`
	UpdateTime string `json:"update_time" gorm:"Column:update_time"`
	DeleteTime string `json:"delete_time" gorm:"Column:delete_time"`
	Operator   string `json:"operator" gorm:"Column:operator"`
	OperatorId int    `json:"operator_id" gorm:"Column:operator_id"`
	Status     int    `json:"status" gorm:"Column:status"`
}

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

// 用户注册
func Register(db *gorm.DB, userName, passWord string) (int, error) {
	var role Role
	db = db.Where("role.code = 'user'")
	err := db.First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	fmt.Println("当前role", role)
	user := UserList{UserName: userName, Password: passWord, Status: 1, CreateTime: utils.NowTimeString(), UpdateTime: utils.NowTimeString(), RoleId: 1}
	res := db.Omit("DeleteTime", "ID").Create(&user)
	if res.Error != nil {
		return -1, res.Error
	}
	return 0, nil
}
