/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:08:32
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-18 10:05:19
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
