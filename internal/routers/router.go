/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 16:05:31
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-14 16:07:56
 */
package routers

import (
	"go-admin/internal/middleware"
	apigroup "go-admin/internal/routers/api_group"
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
	// v1Api.Use(middleware.JWT())
	v1Api.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"test": "test",
		})
	})

	// commonApi := v1Api.Group("/common")
	// common := api.NewCommonApi()
	// commonApi.POST("/login", common.Login)
	apigroup.InitCommonRouter(v1Api)
	// userApi := v1Api.Group("/user")
	// user := v1.NewUserApi()
	// userApi.GET("/role/:id", user.QueryUserAuth)
	apigroup.InitRbacRouter(v1Api)

	return r
}
