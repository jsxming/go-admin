/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:10:33
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-04 15:58:18
 */
package api

import (
	"go-admin/internal/service"
	"go-admin/pkg/errorcode"
	"go-admin/pkg/util"

	"github.com/gin-gonic/gin"
)

type Common struct{}

func NewCommonApi() *Common {
	return &Common{}
}

func (c *Common) Login(ctx *gin.Context) {
	params := service.LoginRequestParams{}
	err := util.BindAndValid(ctx, &params)

	if err != nil {
		util.ToResFail(ctx, errorcode.InvalidParams.AppendDetails(err.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	r, errs := svc.Login(&params)

	if errs != nil {
		util.ToResFail(ctx, errorcode.NewError(212, errs.Error()))
		return
	}

	util.ToResSuccess(ctx, r)
}
