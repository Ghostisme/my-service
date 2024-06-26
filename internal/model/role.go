package model

import (
	"fmt"
	"my-service/pkg/app"
	"time"

	"github.com/jinzhu/gorm"
)

type Role struct {
	*Model
	// gorm.Model
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

// 为Role绑定表名
func (t Role) TableName() string {
	return "role"
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
	db = db.Where("is_delete = 0")
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
func (t Role) Create(db *gorm.DB, userId uint32, isAdmin, status int, name string) (int, error) {
	var roleCode string
	if isAdmin == 0 {
		roleCode = "admin"
	} else {
		roleCode = "user"
	}
	role := Role{
		RoleName:   name,
		Status:     &status,
		RoleCode:   roleCode,
		ActionId:   int(userId),
		ActionTime: time.Now().Format("2006-01-02 15:04:05"),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
		// DeleteTime: sql.NullString{Valid: false},
	}
	db = db.Omit("ID", "DeleteTime").Create(&role)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return int(role.ID), nil
}

// 编辑角色
func (t Role) Update(db *gorm.DB, userId uint32, id int, name string) (int, error) {
	// role := Role{}
	// db = db.Model(&role).Select("role_name", "status").Where("id = ?", id).Updates(Role{
	// 	Model:    &Model{ID: uint32(id)},
	// 	RoleName: "123",
	// 	Status:   t.Status,
	// })
	var role Role
	db = db.Find(&role, id)
	// if name != "" {
	// 	return t.UpdateName(db, userId, id, name)
	// } else {
	// 	return t.UpdateStatus(db, userId, id)
	// }
	// db = db.Update("status", t.Status)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	db = db.Updates(map[string]interface{}{
		"role_name":   name,
		"status":      t.Status,
		"action_id":   userId,
		"action_time": time.Now().Format("2006-01-02 15:04:05"),
	})
	err = db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return 0, nil
}
func (t Role) UpdateName(db *gorm.DB, userId uint32, id int, name string) (int, error) {
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return 0, nil
}
func (t Role) UpdateStatus(db *gorm.DB, userId uint32, id int) (int, error) {
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return 0, nil
}

// 删除角色
func (t Role) Delete(db *gorm.DB, userId uint32, id int) (int, error) {
	// role := DelRole{
	// 	&Model{ID: uint32(id)},
	// }
	var role Role
	db = db.Find(&role, id)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	db = db.Updates(map[string]interface{}{
		"is_delete":   1,
		"delete_time": time.Now().Format("2006-01-02 15:04:05"),
		"action_id":   userId,
		"action_time": time.Now().Format("2006-01-02 15:04:05"),
	})
	err = db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return 0, nil
}
