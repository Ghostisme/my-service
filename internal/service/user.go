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

// 获取用户列表
func (svc *Service) UserList(param *UserListRequest, pager *app.Pager) ([]*model.UserList, error) {
	return svc.dao.UserList(param.Status, param.BeginTime, param.EndTime, param.KeyWord, pager.Page, pager.PageSize)
}

// 获取用户列表总条数
func (svc *Service) UserListCount(param *UserListRequest) (int, error) {
	return svc.dao.UserListCount(param.Status, param.BeginTime, param.EndTime, param.KeyWord)
}
