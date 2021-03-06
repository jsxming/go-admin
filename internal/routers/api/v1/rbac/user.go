/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:02:25
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-30 14:55:04
 */
package v1

import (
	"go-admin/internal/model"
	"go-admin/internal/service"
	"go-admin/pkg/errorcode"
	"go-admin/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryUserRole(ctx *gin.Context) {
	id := ctx.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		util.ToResFail(ctx, errorcode.SearchFail.AppendDetails(err.Error()))
		return
	}

	svc := service.New(ctx.Request.Context())
	ids, err := svc.QueryUserRole(numId)
	if err != nil {
		util.ToResFail(ctx, errorcode.SearchFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, ids)
}

func QueryUserAuth(ctx *gin.Context) {
	var id int
	if u, ok := ctx.Get("user"); ok {
		a := u.(model.User)
		id = int(a.ID)
	}

	svc := service.New(ctx.Request.Context())
	ids, err := svc.QueryUserAuth(id)
	if err != nil {
		util.ToResFail(ctx, errorcode.SearchFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, ids)
}

func QueryUserAll(ctx *gin.Context) {
	svc := service.New(ctx.Request.Context())
	arr, err := svc.QueryUserAll()
	if err != nil {
		util.ToResFail(ctx, errorcode.SearchFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, arr)
}

/*
更新用户角色
*/
func UpdateUserRole(ctx *gin.Context) {
	var params service.UserRoleUpdateParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		util.ToResFail(ctx, errorcode.InvalidParams.AppendDetails(err.Error()))
		return
	}
	svc := service.New(ctx.Request.Context())
	err = svc.UpdateUserRole(params)
	if err != nil {
		util.ToResFail(ctx, errorcode.UpdateFail.AppendDetails(err.Error()))
		return
	}
	util.ToResSuccess(ctx, nil)
}

// type Page struct {
// 	Size    int `json:"size"`
// 	Current int `json:"current"`
// }

func QueryUserPage(ctx *gin.Context) {
	// var params RoleAuthParams
	// err := ctx.ShouldBind(&params)
	// if err != nil {
	// 	util.ToResFail(ctx, errorcode.InvalidParams.AppendDetails(err.Error()))
	// 	return
	// }

}
