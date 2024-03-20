package service

import (
	"my-service/internal/model"
	"my-service/pkg/app"
)

type UserListRequest struct {
	Status    *int   `form: "status" binding:"required,gte=0"`
	BeginTime string `form: "beginTime"`
	EndTime   string `form: "endTime"`
	KeyWord   string `form: "keyWord"`
}

// 获取用户列表
func (svc *Service) UserList(param *UserListRequest, pager *app.Pager) ([]*model.UserList, error) {
	return svc.dao.UserList(*param.Status, param.BeginTime, param.EndTime, param.KeyWord, pager.Page, pager.PageSize)
}

// 获取用户列表总条数
func (svc *Service) UserListCount(param *UserListRequest) (int, error) {
	return svc.dao.UserListCount(*param.Status, param.BeginTime, param.EndTime, param.KeyWord)
}
