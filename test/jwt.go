/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 18:06:02
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-15 18:08:31
 */

package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	name string
	age  int
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

}
