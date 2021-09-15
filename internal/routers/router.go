/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 16:05:31
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-15 16:55:26
 */
package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	return r
}
