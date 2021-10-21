/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:02:25
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-21 17:42:10
 */
package v1

import (
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
