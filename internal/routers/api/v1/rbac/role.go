/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 14:11:37
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-15 15:25:39
 */
package v1

import (
	"fmt"
	"go-admin/internal/service"
	"go-admin/pkg/errorcode"
	"go-admin/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type Role struct{}

// func NewRoleApi() Role {
// 	return Role{}
// }

func QueryRoleAll(ctx *gin.Context) {
	svc := service.New(ctx.Request.Context())
	roleArr, err := svc.QueryRoleAll()
	if err != nil {
		util.ToResFail(ctx, errorcode.SearchFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, roleArr)
}

func QueryRoleAuth(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		util.ToResFail(ctx, errorcode.SearchFail.AppendDetails(err.Error()))
		return
	}
	svc := service.New(ctx.Request.Context())
	r, err := svc.QueryRoleAuth(num)
	if err != nil {
		util.ToResFail(ctx, errorcode.SearchFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, r)
}

type RoleAuthParams struct {
	RoleId int   `form:"roleId"`
	Auth   []int `form:"auth"`
}

func UpdateRoleAuth(ctx *gin.Context) {
	var params RoleAuthParams
	err := ctx.ShouldBind(&params)
	fmt.Printf("params === %v", params)
	if err != nil {
		util.ToResFail(ctx, errorcode.InvalidParams.AppendDetails(err.Error()))
		return
	}
	svc := service.New(ctx.Request.Context())
	err = svc.UpdateRoleAuth(params.RoleId, params.Auth)
	if err != nil {
		util.ToResFail(ctx, errorcode.UpdateFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, nil)
}
