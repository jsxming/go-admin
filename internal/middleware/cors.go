/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 09:38:18
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-16 09:40:31
 */
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			//c.AbortWithStatus(http.StatusNoContent)
			c.JSON(http.StatusOK, "Options Request Ok!")
		}

		// 处理请求
		c.Next()
	}
}
