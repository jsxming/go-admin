/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 14:19:26
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-14 14:25:11
 */
package apigroup

import (
	v1 "go-admin/internal/routers/api/v1/rbac"

	"github.com/gin-gonic/gin"
)

func InitRbacRouter(r *gin.RouterGroup) {
	userApi := r.Group("/user")
	// user := v1.NewUserApi()
	user := v1.NewUserApi()
	userApi.GET("/role/:id", user.QueryUserAuth)

	roleApi := r.Group("/role")
	role := v1.NewRoleApi()
	roleApi.GET("/all", role.QueryRoleAll)

}
