package service

import (
	"my-service/internal/model"
	"my-service/pkg/app"
)

type RoleListRequest struct {
	Status    *int   `form: "status" binding:"required,gte=0"`
	BeginTime string `form: "beginTime"`
	EndTime   string `form: "endTime"`
	KeyWord   string `form: "keyWord"`
}

// 获取角色列表
func (svc *Service) RoleList(param *RoleListRequest, pager *app.Pager) ([]*model.Role, error) {
	return svc.dao.RoleList(param.BeginTime, param.EndTime, param.KeyWord, *param.Status, pager.Page, pager.PageSize)
}

// 获取角色列表总条数
func (svc *Service) RoleListCount(param *RoleListRequest) (int, error) {
	return svc.dao.RoleListCount(param.BeginTime, param.EndTime, param.KeyWord, *param.Status)
}
