package dao

import (
	"my-service/internal/model"
	"my-service/pkg/app"
)

// 获取角色列表
func (d *Dao) RoleList(beginTime, endTime, keyWord string, status *int, page, pageSize int) ([]*model.Role, error) {
	role := model.Role{Status: status}
	pageOffset := app.GetPageOffset(page, pageSize)
	return role.List(d.engine, beginTime, endTime, keyWord, pageOffset, pageSize)
}

// 获取角色列表总条数
func (d *Dao) RoleListCount(beginTime, endTime, keyWord string, status *int) (int, error) {
	role := model.Role{Status: status}
	return role.ListCount(d.engine, beginTime, endTime, keyWord)
}

// 编辑角色信息
func (d *Dao) UpdateRole(userId uint32, id, status int, name string) (int, error) {
	role := model.Role{Status: &status}
	return role.Update(d.engine, userId, id, name)
}

// 删除角色信息
func (d *Dao) DelRole(userId uint32, id int) (int, error) {
	role := model.Role{}
	return role.Delete(d.engine, userId, id)
}
