/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:08:32
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-30 14:53:05
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
	Tel   string        `json:"tel"`
	Token string        `json:"token"`
	Auth  []*model.Auth `json:"auth"`
}

type AuthParams struct {
	Id    int    `from:"id"`
	Label string `form:"label"`
	Path  string `form:"path"`
}

type UserRoleUpdateParams struct {
	UserId int   `form:"userId"`
	RoleId []int `form:"roleId"`
}

func (svc *Service) Login(params *LoginRequestParams) (*LoginResult, error) {
	u := model.User{
		Tel:      params.Tel,
		Password: params.Password,
	}

	user, err := u.Get(svc.db)
	var auth []*model.Auth
	// fmt.Println(user, err)
	auth, err = user.QueryUserAuth(svc.db)
	return &LoginResult{
		Tel:   user.Tel,
		Auth:  auth,
		Token: util.GenerateToken(user),
	}, err
}

func (svc *Service) QueryUser(params *LoginRequestParams) (model.User, error) {
	u := model.User{
		Tel:      params.Tel,
		Password: params.Password,
	}
	r, err := u.Get(svc.db)
	return r, err
}

func (svc *Service) QueryUserAll() ([]model.User, error) {
	r := model.User{}
	arr, err := r.All(svc.db)
	return arr, err
}

func (svc *Service) QueryUserRole(id int) ([]int, error) {
	a := model.UserRole{
		UserId: id,
	}
	arr, err := a.QueryUserRole(svc.db)
	return arr, err
}

func (svc *Service) QueryUserAuth(id int) ([]*model.Auth, error) {
	a := model.User{
		ID: id,
	}
	arr, err := a.QueryUserAuth(svc.db)
	return arr, err
}

func (svc *Service) QueryRoleAll() ([]model.Role, error) {
	r := model.Role{}
	arr, err := r.All(svc.db)
	return arr, err
}

type AuthTree struct {
	Children []*AuthTree `json:"children,omitempty"`
	model.Auth
}

func (svc *Service) QueryAuthAll() ([]AuthTree, error) {
	r := model.Auth{}
	arr, err := r.All(svc.db)
	if len(arr) > 0 && err == nil {
		m := make(map[int][]*AuthTree)
		var tree []*AuthTree
		for _, authItem := range arr {
			treeNode := AuthTree{Auth: authItem, Children: nil}
			tree = append(tree, &treeNode)
			if authItem.ParentId != 0 {
				if _, ok := m[authItem.ParentId]; ok {
					m[authItem.ParentId] = append(m[authItem.ParentId], &treeNode)
				} else {
					m[authItem.ParentId] = []*AuthTree{&treeNode}
				}
			}
		}
		var result []AuthTree
		for i, treeNode := range tree {
			if _, ok := m[treeNode.Id]; ok {
				treeNode.Children = m[treeNode.Id]
			}
			if treeNode.ParentId == 0 {
				result = append(result, *tree[i])
			}

		}
		return result, nil
	}
	return nil, err
}

func (svc *Service) DelAuth(id int) error {
	a := model.Auth{
		Id: id,
	}
	return a.Del(svc.db)
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

func (svc *Service) UpdateAuth(params AuthParams) error {
	r := model.Auth{
		Id:    params.Id,
		Label: params.Label,
		Path:  params.Path,
	}
	err := r.Update(svc.db)
	return err
}

func (svc *Service) UpdateUserRole(params UserRoleUpdateParams) error {
	r := model.UserRole{
		UserId: params.UserId,
	}
	err := r.UpdateUserRole(svc.db, params.RoleId)
	return err
}
