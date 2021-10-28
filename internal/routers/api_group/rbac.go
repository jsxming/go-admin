/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 14:19:26
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-28 17:44:26
 */
package apigroup

import (
	v1 "go-admin/internal/routers/api/v1/rbac"

	"github.com/gin-gonic/gin"
)

func InitRbacRouter(r *gin.RouterGroup) {
	// r.Use(middleware.JWT())
	userApi := r.Group("/user")
	userApi.GET("/all", v1.QueryUserAll)
	userApi.GET("/role", v1.QueryUserAuth)

	roleApi := r.Group("/role")
	roleApi.GET("/all", v1.QueryRoleAll)
	roleApi.GET("/auth/:id", v1.QueryRoleAuth)
	roleApi.POST("/auth/update", v1.UpdateRoleAuth)

	authApi := r.Group("/auth")
	authApi.GET("/all", v1.QueryAuthAll)

}
