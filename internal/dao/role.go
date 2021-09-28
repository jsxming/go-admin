/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-17 15:17:23
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-17 15:22:10
 */
package dao

import "go-admin/internal/model"

func (dao *Dao) QueryRoleAuth(id int) (result []model.RoleAuth, err error) {
	a := model.RoleAuth{
		RoleId: id,
	}
	return a.FindAuth(dao.engine)
}

func (dao *Dao) QueryRoleAll() (result []model.Role, err error) {
	r := model.Role{}
	return r.All(dao.engine)
}
