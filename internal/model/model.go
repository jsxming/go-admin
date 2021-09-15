/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 12:13:53
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-15 16:34:21
 */
package model

import (
	"fmt"
	"go-admin/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	fmt.Println(databaseSetting)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
