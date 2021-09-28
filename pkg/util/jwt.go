/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 18:09:53
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-16 15:19:06
 */

package util

import (
	"go-admin/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken() string {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		// AppKey:    EncodeMD5(appKey),
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
