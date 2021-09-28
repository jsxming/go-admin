/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:50:07
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-17 10:27:59
 */
package dao

import "go-admin/internal/model"

type User struct {
	Id string `json:"id`
}

func (dao *Dao) QueryUserOne(tel, password string) (model.User, error) {
	a := model.User{
		Tel:      tel,
		Password: password,
	}
	return a.Get(dao.engine)
}

func (dao *Dao) QueryUserAuth(id int) ([]int, error) {
	a := model.UserRole{
		UserId: id,
	}
	return a.QueryUserRole(dao.engine)
}
