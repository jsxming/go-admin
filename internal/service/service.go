/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 12:06:13
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-18 10:04:06
 */
package service

import (
	"context"
	"go-admin/global"

	"gorm.io/gorm"
)

type Service struct {
	ctx context.Context
	// dao *dao.Dao
	db *gorm.DB
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.db = global.DBEngine
	return svc
}
