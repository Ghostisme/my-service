package model

import (
	"my-service/pkg/app"

	"github.com/jinzhu/gorm"
)

type Role struct {
	*Model
	RoleCode   string `json:"role" gorm:"Column:role_code"`
	RoleName   string `json:"role_name" gorm:"Column:role_name"`
	CreateTime string `json:"create_time" gorm:"Column:create_time"`
	UpdateTime string `json:"update_time" gorm:"Column:update_time"`
	DeleteTime string `json:"delete_time" gorm:"Column:delete_time"`
	ActionTime string `json:"action_time" gorm:"Column:action_time"`
	ActionId   int    `json:"action_id" gorm:"Column:action_id"`
	Status     int    `json:"status" gorm:"Column:status"`
}

type roleList struct {
	List  []*Role
	Pager *app.Pager `json:"pager"`
}

type RoleInfo struct {
	*SwaggerSuccess
	Data *roleList `json:"data"`
}

// 查询角色列表
func (t Role) List(db *gorm.DB, beginTime, endTime, keyWord string, page, pageSize int) ([]*Role, error) {
	var role []*Role
	if beginTime != "" && endTime != "" {
		db = db.Where("create_time >= ? AND end_time <= ?", beginTime, endTime)
	}
	if keyWord != "" {
		db = db.Where("")
	}
	db = db.Where("")
	if pageSize > 0 {
		db = db.Limit(pageSize)
	}
	err := db.Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

// 查询角色列表总数
func (t Role) ListCount(db *gorm.DB, beginTime, endTime, keyWord string) (int, error) {
	var role Role
	var count int

	// db = db.Select("Count(role.id)").Where("")
	err := db.Find(&role).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
