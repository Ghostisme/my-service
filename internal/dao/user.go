package dao

import (
	"my-service/internal/model"
	"my-service/pkg/app"
)

/**
 * @brief 获取用户列表
 * @param beginTime-开始时间
 * @param endTime-结束时间
 * @param keyWord-关键词
 * @param status-是否有效
 * @param page-页码
 * @param pageSize-每页条数
 * @return 用户数组集合，错误信息
 */
func (d *Dao) UserList(status int, beginTime, endTime, keyWord string, page, pageSize int) ([]*model.UserList, error) {
	user := model.User{Status: status}
	pageOffset := app.GetPageOffset(page, pageSize)
	return user.List(d.engine, beginTime, endTime, keyWord, pageOffset, pageSize)
}

// 获取用户列表总数
func (d *Dao) UserListCount(status int, beginTime, endTime, keyWord string) (int, error) {
	user := model.User{Status: status}
	return user.ListCount(d.engine, beginTime, endTime, keyWord)
}
