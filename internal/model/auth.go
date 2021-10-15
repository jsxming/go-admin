/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 16:37:13
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-14 17:17:31
 */

package model

import "gorm.io/gorm"

type Auth struct {
	Id       int    `json:"id"`
	Path     string `json:"path"`
	Label    string `json:"label"`
	ParentId int    `json:"parentId"`
}

func (a Auth) TabName() string {
	return "authority"
}

func (a Auth) All(DB *gorm.DB) ([]Auth, error) {
	var list []Auth
	err := DB.Table(a.TabName()).Select("id,path,label,parent_id").Scan(&list).Error
	return list, err
}
