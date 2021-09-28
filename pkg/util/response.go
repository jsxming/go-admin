/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 18:29:32
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-16 10:42:39
 */

package util

import (
	"go-admin/pkg/errorcode"

	"github.com/gin-gonic/gin"
)

func ToResSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"code":    1,
		"message": "",
		"data":    data,
	})
}

func ToResFail(ctx *gin.Context, e *errorcode.Error) {
	ctx.JSON(200, gin.H{
		"code":           e.Code(),
		"message":        e.Msg(),
		"messageDetails": e.Details(),
		"data":           nil,
	})
}
