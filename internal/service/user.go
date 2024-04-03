package service

import (
	"my-service/internal/model"
	"my-service/pkg/app"
)

type UserListRequest struct {
	Status    *int   `json:"status" form:"status"`
	BeginTime string `json:"beginTime" form:"beginTime"`
	EndTime   string `json:"endTime" form:"endTime"`
	KeyWord   string `json:"keyWord" form:"keyWord"`
	// Status    *int   `form: "status" binding:"required,gte=0"`
}

type UserCreateRequest struct {
	UserName string `json:"username" form:"username"`
	Status   int    `json:"status" form:"status"`
	Mobile   string `json:"mobile" form:"mobile"`
	Addr     string `json:"addr" form:"addr"`
	Email    string `json:"email" form:"email"`
	IsAdmin  int    `json:"is_admin" form:"is_admin"`
}

type UserUpdateRequest struct {
	ID       *int   `json:"id" binding:"required" form:"id"`
	Status   int    `json:"status" form:"status"`
	UserName string `json:"username" form:"username"`
	Mobile   string `json:"mobile" form:"mobile"`
	Addr     string `json:"addr" form:"addr"`
	Email    string `json:"email" form:"email"`
}

type UserDelRequest struct {
	ID *int `json:"id" binding:"required" form:"id"`
}

// 获取用户列表
func (svc *Service) UserList(param *UserListRequest, pager *app.Pager) ([]*model.UserList, error) {
	return svc.dao.UserList(param.Status, param.BeginTime, param.EndTime, param.KeyWord, pager.Page, pager.PageSize)
}

// 获取用户列表总条数
func (svc *Service) UserListCount(param *UserListRequest) (int, error) {
	return svc.dao.UserListCount(param.Status, param.BeginTime, param.EndTime, param.KeyWord)
}

// 创建用户
func (svc *Service) CreateUser(param *UserCreateRequest) (int, error) {
	return svc.dao.CreateUser(param.UserName, param.Mobile, param.Addr, param.Email, param.IsAdmin, param.Status)
}

// 编辑用户
func (svc *Service) UpdateUser(param *UserUpdateRequest) (int, error) {
	return svc.dao.UpdateUser(*param.ID, param.Status, param.UserName, param.Mobile, param.Email, param.Addr)
}

// 删除用户
func (svc *Service) DelUser(param *UserDelRequest) (int, error) {
	return svc.dao.DelUser(*param.ID)
}
