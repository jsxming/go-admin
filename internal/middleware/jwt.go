/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 18:13:49
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-14 15:44:27
 */
package middleware

import (
	"fmt"
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
		fmt.Print(token, "token----")
		if token == "" {
			ecode = errcode.NotFoundToken
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
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
