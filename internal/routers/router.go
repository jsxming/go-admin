/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 16:05:31
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-21 16:18:48
 */
package routers

import (
	"go-admin/internal/middleware"
	apigroup "go-admin/internal/routers/api_group"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.Use(middleware.AccessLog())
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "hello world----",
			// "token": util.GenerateToken(),
		})
	})

	v1Api := r.Group("/v1")
	v1Api.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"test": "test",
		})
	})

	apigroup.InitCommonRouter(v1Api)
	apigroup.InitRbacRouter(v1Api)

	return r
}
