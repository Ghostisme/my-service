package dao

import (
	"my-project-admin-service/internal/model"
)

/**
 * @brief 通过用户名获取用户信息
 * @param username-用户名
 * @return 用户ID，错误信息
 */
func (d *Dao) GetUserByName(userName string) (uint32, error) {
	return model.GetUserByName(d.engine, userName)
}

/**
 * @brief 通过用户名&密码获取用户信息
 * @param username-用户名
 * @param password-密码
 * @return 用户ID，错误信息
 */
func (d *Dao) GetUser(userName, Password string) (uint32, error) {

	return model.GetUser(d.engine, userName, Password)
}

/**
 * @brief 通过id设置用户token
 * @param id-用户主键
 * @param token-登录秘钥
 * @return 0，错误信息
 */
func (d *Dao) SetToken(id uint32, token string) (uint32, error) {
	return model.SetCurrentToken(d.engine, id, token)
}