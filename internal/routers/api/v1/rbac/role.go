/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 14:11:37
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-14 14:16:47
 */
package v1

import (
	"go-admin/internal/service"
	"go-admin/pkg/errorcode"
	"go-admin/pkg/util"

	"github.com/gin-gonic/gin"
)

type Role struct{}

func NewRoleApi() Role {
	return Role{}
}

func (r Role) QueryRoleAll(ctx *gin.Context) {
	svc := service.New(ctx.Request.Context())
	roleArr, err := svc.QueryRoleAll()
	if err != nil {
		util.ToResFail(ctx, errorcode.NewError(1200, err.Error()))
		return
	}
	util.ToResSuccess(ctx, roleArr)
}
