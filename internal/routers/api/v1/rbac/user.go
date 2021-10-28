/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:02:25
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-28 17:44:35
 */
package v1

import (
	"fmt"
	"go-admin/internal/model"
	"go-admin/internal/service"
	"go-admin/pkg/errorcode"
	"go-admin/pkg/util"

	"github.com/gin-gonic/gin"
)

func QueryUserAuth(ctx *gin.Context) {
	// id := ctx.Param("id")
	var id int
	if u, ok := ctx.Get("user"); ok {
		a := u.(model.User)
		id = int(a.ID)
	}

	// num, err := strconv.Atoi(id)
	// if err != nil {
	// 	util.ToResFail(ctx, errorcode.InvalidParams)
	// 	return
	// }

	svc := service.New(ctx.Request.Context())
	ids, err := svc.QueryUserAuth(id)
	if err != nil {
		util.ToResFail(ctx, errorcode.NewError(1100, err.Error()))
		return
	}
	util.ToResSuccess(ctx, ids)
}

func QueryUserAll(ctx *gin.Context) {
	svc := service.New(ctx.Request.Context())
	arr, err := svc.QueryUserAll()
	fmt.Println("...")
	if err != nil {
		util.ToResFail(ctx, errorcode.SearchFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, arr)
}
