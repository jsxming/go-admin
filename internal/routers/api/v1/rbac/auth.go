/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 16:44:08
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-30 11:44:11
 */
package v1

import (
	"go-admin/internal/service"
	"go-admin/pkg/errorcode"
	"go-admin/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryAuthAll(ctx *gin.Context) {
	svc := service.New(ctx.Request.Context())
	roleArr, err := svc.QueryAuthAll()
	if err != nil {
		util.ToResFail(ctx, errorcode.SearchFail.AppendDetails(err.Error()))
		// util.ToResFail(ctx, errorcode.NewError(1400, err.Error()))
		return
	}
	util.ToResSuccess(ctx, roleArr)
}

func DelAuth(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	svc := service.New(ctx.Request.Context())
	err := svc.DelAuth(idInt)
	if err != nil {
		util.ToResFail(ctx, errorcode.DeleteFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, nil)
}

func UpdateAuth(ctx *gin.Context) {
	var params service.AuthParams
	err := ctx.ShouldBind(&params)
	if params.Id == 0 {
		util.ToResFail(ctx, errorcode.InvalidParams.AppendDetails())
		return
	}
	if err != nil {
		util.ToResFail(ctx, errorcode.InvalidParams.AppendDetails(err.Error()))
		return
	}
	svc := service.New(ctx.Request.Context())
	err = svc.UpdateAuth(params)
	if err != nil {
		util.ToResFail(ctx, errorcode.UpdateFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, nil)
}
