/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:02:25
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-14 14:17:30
 */
package v1

import (
	"go-admin/internal/service"
	"go-admin/pkg/errorcode"
	"go-admin/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUserApi() User {
	return User{}
}

func (a User) QueryUserAuth(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		util.ToResFail(ctx, errorcode.InvalidParams)
		return
	}

	svc := service.New(ctx.Request.Context())
	ids, err := svc.QueryUserAuth(num)
	if err != nil {
		util.ToResFail(ctx, errorcode.NewError(1100, err.Error()))
		return
	}
	util.ToResSuccess(ctx, ids)
}
