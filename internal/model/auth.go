/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 16:37:13
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-30 11:38:36
 */

package model

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Auth struct {
	Id       int                   `json:"id"`
	Path     string                `json:"path"`
	Label    string                `json:"label"`
	ParentId int                   `json:"parentId"`
	IsPage   uint8                 `json:"isPage"`
	IsDel    soft_delete.DeletedAt `json:"-" gorm:"softDelete:flag"`
}

func (a Auth) TabName() string {
	return "authority"
}

func (a Auth) All(DB *gorm.DB) ([]Auth, error) {
	var list []Auth
	err := DB.Table(a.TabName()).Select("id,path,label,parent_id,is_page").Where("is_del =0").Scan(&list).Error
	return list, err
}

func (a Auth) Del(DB *gorm.DB) error {
	err := DB.Table(a.TabName()).Where("id = ? or parent_id = ?", a.Id, a.Id).Update("is_del", 1).Error
	// err := DB.Delete(&a).Error
	return err
}

func (a Auth) Update(DB *gorm.DB) error {
	fmt.Println(a.Id, a.Label, a.Path, "...")
	return DB.Table(a.TabName()).Where("id = ?", a.Id).Updates(a).Error
}
