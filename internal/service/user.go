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

func (svc *Service) List(param *UserListRequest, pager *app.Pager) ([]*model.UserList, error) {
	return svc.dao.List(*param.Status, param.BeginTime, param.EndTime, param.KeyWord, pager.Page, pager.PageSize)
}

func (svc *Service) ListCount(param *UserListRequest) (int, error) {
	return svc.dao.ListCount(*param.Status, param.BeginTime, param.EndTime, param.KeyWord)
}
