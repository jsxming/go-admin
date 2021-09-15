/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 16:56:42
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-15 17:00:21
 */
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "adf"})
	})
	if err := r.Run(fmt.Sprintf(":%d", viper.Get("port"))); err != nil {
		panic(err)
	}
	// r.Run("9000")
}
