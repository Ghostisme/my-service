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
