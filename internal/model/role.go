/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-17 13:45:17
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-14 17:57:34
 */
package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Role struct {
	Id       int    `json:"id"`
	RoleName string `json:"roleName"`
}

type RoleAuth struct {
	Id     int `json:"id"`
	RoleId int `json:"roleId"`
	AuthId int `json:"authId"`
}

func (role *Role) TabName() string {
	return "role"
}

func (role *Role) Create(DB *gorm.DB) error {
	return DB.Create(role).Error
}

func (role Role) Update(DB *gorm.DB) error {
	return DB.Model(&role).Updates(role).Error
}

func (role Role) All(DB *gorm.DB) ([]Role, error) {
	var list []Role
	err := DB.Table(role.TabName()).Select("id,role_name").Scan(&list).Error
	return list, err
}

func (ra *RoleAuth) TabName() string {
	return "role_auth"
}

func (ra RoleAuth) FindAuth(DB *gorm.DB) (result []RoleAuth, err error) {
	fmt.Print(ra.RoleId)
	err = DB.Where("role_id = ?", ra.RoleId).Find(&result).Error
	return
}

func (ra RoleAuth) UpdateRoleAuth(DB *gorm.DB, auths []int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", ra.RoleId).Delete(RoleAuth{}).Error; err != nil {
			return err
		}
		userRoles := make([]RoleAuth, 0)
		for _, v := range auths {
			userRoles = append(userRoles, RoleAuth{
				RoleId: ra.RoleId,
				AuthId: v,
			})
		}
		if err := tx.Create(&userRoles).Error; err != nil {
			return err
		}
		return nil
	})
}
