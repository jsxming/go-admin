/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-17 13:45:17
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-17 13:47:48
 */
package model

import "gorm.io/gorm"

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

// func (ra *RoleAuth) UpdateRoleAuth(roleId int, auths []int) error {
// 	return DB.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Where("role_id = ?", roleId).Delete(RoleAuth{}).Error; err != nil {
// 			return err
// 		}
// 		userRoles := make([]RoleAuth, 0)
// 		for _, v := range auths {
// 			userRoles = append(userRoles, RoleAuth{
// 				RoleId: roleId,
// 				AuthId: v,
// 			})
// 		}
// 		if err := tx.Create(&userRoles).Error; err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }

func (ra *RoleAuth) TabName() string {
	return "role_auth"
}

func (ra RoleAuth) FindAuth(DB *gorm.DB) (result []RoleAuth, err error) {
	err = DB.Where("role_id = ?", ra.RoleId).Find(&result).Error
	return
}
