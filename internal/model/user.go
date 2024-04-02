package model

import (
	"fmt"
	"my-service/global"
	"my-service/pkg/app"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	*Model
	UserName   string `json:"username" gorm:"Column:username"`
	Password   string `json:"-" gorm:"Column:password"`
	CreateTime string `json:"create_time" gorm:"Column:create_time"`
	UpdateTime string `json:"update_time" gorm:"Column:update_time"`
	DeleteTime string `json:"delete_time" gorm:"Column:delete_time"`
	Status     *int   `json:"status" gorm:"Column:status"`
	Mobile     string `json:"mobile" gorm:"Column:mobile"`
	Addr       string `json:"addr" gorm:"Column:addr"`
	Email      string `json:"email" gorm:"Column:email"`
	RoleId     int    `json:"-" gorm:"Column:role_id"`
	Role       Role   `json:"role" gorm:"foreignKey:RoleId;references:ID;"`
}

type UserList struct {
	*Model
	UserName   string `json:"username" gorm:"Column:username"`
	Password   string `json:"-" gorm:"Column:password"`
	CreateTime string `json:"create_time" gorm:"Column:create_time"`
	UpdateTime string `json:"update_time" gorm:"Column:update_time"`
	DeleteTime string `json:"delete_time" gorm:"Column:delete_time"`
	Status     int    `json:"status" gorm:"Column:status"`
	Mobile     string `json:"mobile" gorm:"Column:mobile"`
	Addr       string `json:"addr" gorm:"Column:addr"`
	Email      string `json:"email" gorm:"Column:email"`
	IsDelete   int    `json:"is_delete" gorm:"Column:is_delete"`
	RoleId     int    `json:"-" gorm:"Column:role_id"`
}

type userList struct {
	List  []*User
	Pager *app.Pager `json:"pager"`
}

type UserInfo struct {
	*SwaggerSuccess
	Data *userList `json:"data"`
}

// 查询用户列表集合
func (t User) List(db *gorm.DB, beginTime, endTime, keyWord string, page, pageSize int) ([]*UserList, error) {
	var user []*UserList
	if beginTime != "" && endTime != "" {
		db = db.Where("create_time >= ? AND create_time <= ?", beginTime, endTime)
	}
	if keyWord != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf(`%%%s%%`, keyWord))
		fmt.Println("keyWord", db)
	}
	if t.Status != nil {
		db = db.Where("status = ?", t.Status)
	}
	if page >= 0 && pageSize > 0 {
		db = db.Offset(page).Limit(pageSize)
	}
	db = db.Where("`user`.is_delete = 0")
	err := db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	// query := "SELECT `user`.* FROM `user` "
	// query += fmt.Sprintf("`user`.status = %d", t.Status)
	// if beginTime != "" && endTime != "" || keyWord != "" || t.Status != nil {
	// 	query += "WHERE "
	// }
	// if beginTime != "" && endTime != "" {
	// 	query += fmt.Sprintf("`user`.create_time >= '%s' AND `user`.create_time <= '%s' AND ", beginTime, endTime)
	// }
	// if keyWord != "" {
	// 	query += fmt.Sprintf("`user`.username LIKE '%%%s%%' AND ", keyWord)
	// }
	// if t.Status != nil {
	// 	query += fmt.Sprintf("`user`.status = %d ", t.Status)
	// }
	// query += " ORDER BY `user`.id ASC "
	// if page >= 0 && pageSize > 0 {
	// 	query += fmt.Sprintf("LIMIT %d, %d", page, pageSize)
	// }
	// res := db.Raw(query).Scan(&user)
	// if res.Error != nil {
	// 	return nil, res.Error
	// }
	return user, nil
}

// 查询用户列表总条数
func (t User) ListCount(db *gorm.DB, beginTime, endTime, keyWord string) (int, error) {
	var user User
	count := 0
	if beginTime != "" && endTime != "" {
		db = db.Where("create_time >= ? AND end_time <= ?", beginTime, endTime)
	}
	if keyWord != "" {
		db = db.Where("`user`.username LIKE ?", fmt.Sprintf(`%%%s%%`, keyWord))
	}
	if t.Status != nil {
		db = db.Where("status = ?", t.Status)
	}
	db = db.Select("Count(`user`.id)").Where("is_delete = 0")
	err := db.Find(&user).Count(&count).Error
	if err != nil {
		return 0, err
	}
	// query := "SELECT Count(`user`.id) as count FROM `user` "
	// if beginTime != "" && endTime != "" || keyWord != "" || t.Status != nil {
	// 	query += "WHERE "
	// }
	// if beginTime != "" && endTime != "" {
	// 	query += fmt.Sprintf("`user`.create_time >= '%s' AND `user`.create_time <= '%s' AND ", beginTime, endTime)
	// }
	// if keyWord != "" {
	// 	query += fmt.Sprintf("`user`.username LIKE '%%%s%%' AND ", keyWord)
	// }
	// if t.Status != nil {
	// 	query += fmt.Sprintf("`user`.status = %d ", t.Status)
	// }
	// res := db.Raw(query).Count(&count)
	// if res.Error != nil {
	// 	return -1, res.Error
	// }
	return count, nil
}

// 创建新用户
func (t User) Create(db *gorm.DB, username, password, mobile, addr, email string, isAdmin, status int) (int, error) {
	var roleCode string
	if isAdmin == 0 {
		roleCode = "admin"
	} else {
		roleCode = "user"
	}
	var role Role
	db = db.Select("id").Where("role_code = ?", roleCode)
	err := db.Find(&role).Error
	if err != nil {
		return -1, err
	}
	global.ModelLogger.Info("查询的role_id", role.ID)
	user := UserList{
		UserName:   username,
		Password:   password,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
		Mobile:     mobile,
		Addr:       addr,
		Email:      email,
		IsDelete:   0,
		Status:     status,
		RoleId:     int(role.ID),
	}
	db = db.Omit("ID").Create(&user)
	err = db.Error
	if err != nil {
		return -1, err
	}
	return 0, nil
}

// 编辑用户
func (t User) Update(db *gorm.DB, id int) (int, error) {
	var user UserList
	db = db.Find(&user, id)
}
