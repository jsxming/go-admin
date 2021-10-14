/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 14:26:30
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-14 14:26:31
 */
package apigroup

import (
	api "go-admin/internal/routers/api/common"

	"github.com/gin-gonic/gin"
)

func InitCommonRouter(r *gin.RouterGroup) {
	commonApi := r.Group("/common")
	common := api.NewCommonApi()
	commonApi.POST("/login", common.Login)
}
