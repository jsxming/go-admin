/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 12:04:46
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-15 12:07:21
 */
package dao

import "gorm.io/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
