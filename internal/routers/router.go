/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 16:05:31
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-28 15:12:01
 */
package routers

import (
	"go-admin/internal/middleware"
	api "go-admin/internal/routers/api/common"
	v1 "go-admin/internal/routers/api/v1/rbac"
	"go-admin/pkg/util"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg":   "hello world----",
			"token": util.GenerateToken(),
		})
	})

	v1Api := r.Group("/v1")
	v1Api.Use(middleware.JWT())
	v1Api.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"test": "test",
		})
	})

	commonApi := v1Api.Group("/common")
	common := api.NewCommonApi()
	commonApi.POST("/login", common.Login)

	userApi := v1Api.Group("/user")
	user := v1.NewUserApi()
	userApi.GET("/role/:id", user.QueryUserAuth)

	return r
}
