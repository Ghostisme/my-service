package dao

import (
	"my-service/internal/model"
	"my-service/pkg/app"
)

// 获取用户列表
func (d *Dao) UserList(status *int, beginTime, endTime, keyWord string, page, pageSize int) ([]*model.UserList, error) {
	user := model.User{Status: status}
	pageOffset := app.GetPageOffset(page, pageSize)
	return user.List(d.engine, beginTime, endTime, keyWord, pageOffset, pageSize)
}

// 获取用户列表总数
func (d *Dao) UserListCount(status *int, beginTime, endTime, keyWord string) (int, error) {
	user := model.User{Status: status}
	return user.ListCount(d.engine, beginTime, endTime, keyWord)
}

// 创建用户
func (d *Dao) CreateUser(username, mobile, addr, email string, isAdmin, status int) (int, error) {
	user := model.User{}
	option := model.NewOption(model.WithCreate(username, mobile, addr, email, isAdmin, status))
	return user.Create(d.engine, option)
}

// 编辑用户
func (d *Dao) UpdateUser(id, status int, username, mobile, email, addr string) (int, error) {
	user := model.User{}
	return user.Update(d.engine, id, status, username, mobile, email, addr)
}

// 删除用户
func (d *Dao) DelUser(id int) (int, error) {
	user := model.User{}
	return user.Del(d.engine, id)
}
