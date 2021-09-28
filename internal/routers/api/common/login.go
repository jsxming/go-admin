/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 10:10:33
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-17 09:18:56
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
	_, err := util.BindAndValid(ctx, &params)

	if err != nil {
		util.ToResFail(ctx, errorcode.InvalidParams.AppendDetails(err.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	r, errs := svc.QueryUser(&params)
	if errs != nil {
		util.ToResFail(ctx, errorcode.NewError(212, errs.Error()))
		return
	}

	util.ToResSuccess(ctx, r)
}
