package dao

import (
	"my-service/internal/model"
)

/**
 * @brief 通过用户名获取用户信息
 * @param username-用户名
 * @return 用户ID，错误信息
 */
// func (d *Dao) GetUserByName(userName string) (uint32, error) {
// 	return model.GetUserByName(d.engine, userName)
// }

/**
 * @brief 通过用户名&密码获取用户信息
 * @param username-用户名
 * @param password-密码
 * @return 用户ID，错误信息
 */
// func (d *Dao) GetUser(userName, Password string) (*model.User, error) {
// 	return model.GetUser(d.engine, userName, Password)
// }

/*
 * @brief 用户登录
 * @param username-用户名
 * @param password-密码
 * @return 当前登录用户信息及角色信息，错误信息
 */
func (d *Dao) Login(UserName, Password string) (*model.User, error) {
	return model.Login(d.engine, UserName, Password)
}

/*
 * @brief 用户名注册
 * @param username-用户名
 * @param password-密码
 * @return
 */
func (d *Dao) RegisterUser(UserName, Password string) (int, error) {
	return model.RegisterUser(d.engine, UserName, Password)
}

/*
 * @brief 手机号注册
 * @param mobile-手机号
 * @return
 */
func (d *Dao) RegisterMobile(Mobile string) (int, error) {
	return model.RegisterMobile(d.engine, Mobile)
}
