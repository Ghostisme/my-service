package model

import (
	"fmt"
	"my-service/pkg/app"
	"my-service/pkg/cryptor"
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

type CreateUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Status   int    `json:"status"`
	Mobile   string `json:"mobile"`
	Addr     string `json:"addr"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
	IsAdmin  int    `json:"is_admin"`
}

// type Option *CreateUser
type Option func(*CreateUser)

type userList struct {
	List  []*User
	Pager *app.Pager `json:"pager"`
}

type UserInfo struct {
	*SwaggerSuccess
	Data *userList `json:"data"`
}

var serviceSecretKey = "mxalxjzj9oeffag9"

// 为UserList绑定表名
func (t UserList) TableName() string {
	return "user"
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
	if err != nil && err != gorm.ErrRecordNotFound {
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
	if err != nil && err != gorm.ErrRecordNotFound {
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
func WithUserRegister(username, password string) Option {
	return func(c *CreateUser) {
		c.UserName = username
		c.Password = password
		c.Mobile = ""
		c.Email = ""
		c.Addr = ""
		c.Status = 1
		c.IsAdmin = 1
	}
}
func WithMobileRegister(mobile string) Option {
	return func(c *CreateUser) {
		password := cryptor.AesSimpleEncrypt("123456", serviceSecretKey)
		c.UserName = mobile
		c.Password = password
		c.Mobile = mobile
		c.Status = 1
		c.IsAdmin = 1
	}
}

func WithCreate(username, mobile, addr, email string, isAdmin, status int) Option {
	return func(c *CreateUser) {
		password := cryptor.AesSimpleEncrypt("123456", serviceSecretKey)
		c.UserName = username
		c.Password = password
		c.Mobile = mobile
		c.Addr = addr
		c.Email = email
		c.IsAdmin = isAdmin
		c.Status = status
	}
}

func NewOption(opts ...Option) *CreateUser {
	user := &CreateUser{}
	for _, opt := range opts {
		opt(user)
	}
	return user
}

func (t User) Create(db *gorm.DB, cu *CreateUser) (int, error) {
	var roleCode string
	role := Role{}
	if cu.IsAdmin == 0 {
		roleCode = "admin"
	} else {
		roleCode = "user"
	}
	db = db.Where("role_code = ?", roleCode).First(&role)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	if cu.Email != "" || cu.Addr != "" {
		user := UserList{
			UserName:   cu.UserName,
			Password:   cu.Password,
			Mobile:     cu.Mobile,
			Addr:       cu.Addr,
			Email:      cu.Email,
			IsDelete:   0,
			Status:     cu.Status,
			CreateTime: time.Now().Format("2006-01-02 15:04:05"),
			UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
			RoleId:     int(role.ID),
		}
		db = db.Omit("ID", "DeleteTime").Create(&user)
	} else {
		if cu.Mobile != "" {
			user := UserList{
				UserName:   cu.UserName,
				Password:   cu.Password,
				Mobile:     cu.UserName,
				IsDelete:   0,
				Status:     1,
				CreateTime: time.Now().Format("2006-01-02 15:04:05"),
				UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
				RoleId:     int(role.ID),
			}
			db = db.Omit("DeleteTime", "ID", "Addr", "Email").Create(&user)
		} else {
			user := UserList{
				UserName:   cu.UserName,
				Password:   cu.Password,
				Mobile:     cu.Mobile,
				IsDelete:   0,
				Status:     1,
				CreateTime: time.Now().Format("2006-01-02 15:04:05"),
				UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
				RoleId:     int(role.ID),
			}
			db = db.Omit("DeleteTime", "ID", "Addr", "Email").Create(&user)
		}
	}
	err = db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return 0, nil
}

// 编辑用户
func (t User) Update(db *gorm.DB, id, status int, username, mobile, email, addr string) (int, error) {
	var user UserList
	db = db.Find(&user, id)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	db = db.Updates(map[string]interface{}{
		"username": username,
		"mobile":   mobile,
		"email":    email,
		"addr":     addr,
		"status":   status,
	})
	err = db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return 0, nil
}

// 删除用户
func (t User) Del(db *gorm.DB, id int) (int, error) {
	var user UserList
	db = db.Find(&user, id)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	db = db.Updates(UserList{
		IsDelete:   1,
		DeleteTime: time.Now().Format("2006-01-02 15:04:05"),
	})
	err = db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return 0, nil
}
