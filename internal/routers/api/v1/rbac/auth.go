/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 16:44:08
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-06 17:44:21
 */
package v1

import (
	"fmt"
	"go-admin/internal/service"
	"go-admin/pkg/errorcode"
	"go-admin/pkg/util"

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
	// id := ctx.
	id, ok := ctx.Params.Get("id")
	fmt.Println(id, ok)
}
