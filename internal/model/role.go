package model

import (
	"fmt"
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
	IsDelete   int    `json:"is_delete" gorm:"Column:is_delete"`
	ActionTime string `json:"action_time" gorm:"Column:action_time"`
	ActionId   int    `json:"action_id" gorm:"Column:action_id"`
	Status     *int   `json:"status" gorm:"Column:status"`
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
		db = db.Where("create_time >= ? AND create_time <= ?", beginTime, endTime)
	}
	if keyWord != "" {

		db = db.Where("role_name LIKE ?", fmt.Sprintf(`%%%s%%`, keyWord))
		fmt.Println("keyWord", db)
	}
	if t.Status != nil {
		db = db.Where("status = ?", t.Status)
	}
	if page >= 0 && pageSize > 0 {
		db = db.Offset(page).Limit(pageSize)
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
	if beginTime != "" && endTime != "" {
		db = db.Where("create_time >= ? AND create_time <= ?", beginTime, endTime)
	}
	if keyWord != "" {
		db = db.Where("role_name LIKE ?", fmt.Sprintf(`%%%s%%`, keyWord))
	}
	if t.Status != nil {
		db = db.Where("status = ?", t.Status)
	}
	db = db.Select("Count(role.id)").Where("is_delete = 0")
	err := db.Find(&role).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 创建新角色
func (t Role) Create(db *gorm.DB) (int, error) {
	// var id int
	return 0, nil
}

// 编辑角色
func (t Role) Update(db *gorm.DB, userId uint32, id int, name string) (int, error) {
	fmt.Println("传递的userid，%d", userId)
	db = db.Update()
	return 0, nil
}

// 删除角色
func (t Role) Del(db *gorm.DB, id int) (int, error) {
	return 0, nil
}
