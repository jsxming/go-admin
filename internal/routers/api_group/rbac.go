/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-10-14 14:19:26
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-30 14:55:20
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
	userApi.GET("/page", v1.QueryUserPage)
	userApi.GET("/role/:id", v1.QueryUserRole)
	userApi.GET("/auth", v1.QueryUserAuth)
	userApi.POST("/updaterole", v1.UpdateUserRole)

	roleApi := r.Group("/role")
	roleApi.GET("/all", v1.QueryRoleAll)
	roleApi.GET("/auth/:id", v1.QueryRoleAuth)
	roleApi.POST("/auth/update", v1.UpdateRoleAuth)

	authApi := r.Group("/auth")
	authApi.GET("/all", v1.QueryAuthAll)
	authApi.POST("/update", v1.UpdateAuth)
	authApi.POST("/delete/:id", v1.DelAuth)

}
