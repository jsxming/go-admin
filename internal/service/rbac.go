/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:08:32
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-14 17:58:44
 */
package service

import (
	"go-admin/internal/model"
	"go-admin/pkg/util"
)

type LoginRequestParams struct {
	Tel      string `form:"tel" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResult struct {
	Tel   string `json:"tel"`
	Token string `json:"token"`
}

func (svc *Service) QueryUser(params *LoginRequestParams) (*LoginResult, error) {
	u := model.User{
		Tel:      params.Tel,
		Password: params.Password,
	}
	r, err := u.Get(svc.db)
	return &LoginResult{
		Tel:   r.Tel,
		Token: util.GenerateToken(),
	}, err
}

func (svc *Service) QueryUserAuth(id int) ([]int, error) {
	a := model.UserRole{
		UserId: id,
	}
	arr, err := a.QueryUserRole(svc.db)
	return arr, err
}

func (svc *Service) QueryRoleAll() ([]model.Role, error) {
	r := model.Role{}
	arr, err := r.All(svc.db)
	return arr, err
}

func (svc *Service) QueryAuthAll() ([]model.Auth, error) {
	r := model.Auth{}
	arr, err := r.All(svc.db)
	return arr, err
}

func (svc *Service) QueryRoleAuth(id int) ([]model.RoleAuth, error) {
	r := model.RoleAuth{
		RoleId: id,
	}
	arr, err := r.FindAuth(svc.db)
	return arr, err
}

func (svc *Service) UpdateRoleAuth(id int, auths []int) error {
	r := model.RoleAuth{
		RoleId: id,
	}
	err := r.UpdateRoleAuth(svc.db, auths)
	return err
}
