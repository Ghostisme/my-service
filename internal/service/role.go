package service

import (
	"my-service/global"
	"my-service/internal/model"
	"my-service/pkg/app"
)

type RoleListRequest struct {
	Status    *int   `json:"status" form:"status"`
	BeginTime string `json:"beginTime" form:"beginTime"`
	EndTime   string `json:"endTime" form:"endTime"`
	KeyWord   string `json:"keyWord" form:"keyWord"`
	// Status    *int   `form: "status" binding:"required,gte=0"`
}

type RoleUpdateRequest struct {
	Id     *int   `json:"id" binding:"required" form:"id"`
	Name   string `json:"name" form:"name"`
	Status int    `json:"status" form:"status"`
	// UserId int    `form:"-"`
}

type RoleDelRequest struct {
	Id *int `json:"id" binding:"required" form:"id"`
}

// 获取角色列表
func (svc *Service) RoleList(param *RoleListRequest, pager *app.Pager) ([]*model.Role, error) {
	return svc.dao.RoleList(param.BeginTime, param.EndTime, param.KeyWord, param.Status, pager.Page, pager.PageSize)
}

// 获取角色列表总条数
func (svc *Service) RoleListCount(param *RoleListRequest) (int, error) {
	return svc.dao.RoleListCount(param.BeginTime, param.EndTime, param.KeyWord, param.Status)
}

// 编辑角色状态
func (svc *Service) UpdateRole(UserId uint32, param *RoleUpdateRequest) (int, error) {
	global.ServiceLogger.Info("status", param.Status)
	return svc.dao.UpdateRole(UserId, *param.Id, param.Status, param.Name)
}

// 删除角色
func (svc *Service) DelRole(UserId uint32, param *RoleDelRequest) (int, error) {
	return svc.dao.DelRole(UserId, *param.Id)
}
