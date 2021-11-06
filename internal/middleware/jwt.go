/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 18:13:49
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-04 15:16:01
 */
package middleware

import (
	errcode "go-admin/pkg/errorcode"
	"go-admin/pkg/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		token = c.GetHeader("token")
		if token == "" {
			ecode = errcode.NotFoundToken
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else {
				c.Set("user", claims.User)
			}
		}

		if ecode != errcode.Success {
			util.ToResFail(c, ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
