/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 18:09:53
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-21 14:18:41
 */

package util

import (
	"go-admin/global"
	"go-admin/internal/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	AppSecret string     `json:"app_secret"`
	User      model.User `json:"user"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(u model.User) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		// AppKey:    EncodeMD5(appKey),
		User:      u,
		AppSecret: EncodeMD5(global.JWTSetting.Secret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenClaims.SignedString(GetJWTSecret())
	return token
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 获取token 中的user
func GetTokenUser(c *gin.Context) *model.User {
	b := c.MustGet("user").(*model.User)
	return b
}
