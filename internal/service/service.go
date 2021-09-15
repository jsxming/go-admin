/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 12:06:13
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-15 12:12:02
 */
package service

import (
	"context"
	"go-admin/global"
	"go-admin/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
