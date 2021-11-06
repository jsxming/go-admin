/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 15:03:02
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-06 11:41:36
 */
package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Password string `json:"password" gorm:"-"`
	NickName string `json:"nick_name"`
}

//用户 角色
type UserRole struct {
	Id     int `json:"id"`
	RoleId int `json:"roleId"`
	UserId int `json:"userId"`
}

func (a User) Get(db *gorm.DB) (User, error) {
	db = db.Where(`tel=? and password=?`, a.Tel, a.Password)
	err := db.First(&a).Error
	return a, err
}

func (a User) QueryUserPage(db *gorm.DB, page, size int) ([]User, error) {
	var result []User
	err := db.Select(`id,name,tel,nick_name`).Where(`tel like ?%`, a.Tel).Limit(size).Offset(page * 10).Find(&result).Error
	return result, err
}

func (u *User) TabName() string {
	return "user"
}

func (u *UserRole) TabName() string {
	return "user_role"
}

func (u *User) Create(DB *gorm.DB) error {
	return DB.Create(u).Error
}

func (u User) Update(DB *gorm.DB) error {
	return DB.Model(&u).Updates(u).Error
}

func (u User) All(DB *gorm.DB) ([]User, error) {
	var list []User
	err := DB.Table(u.TabName()).Select("id,name").Scan(&list).Error
	return list, err
}

func (u UserRole) QueryUserRole(DB *gorm.DB) ([]int, error) {
	var result []int
	fmt.Println(u.UserId, "ssss---")
	err := DB.Table(u.TabName()).Select("role_id").Where("user_id = ?", u.UserId).Find(&result).Error
	return result, err
}

func (u User) QueryUserAuth(DB *gorm.DB) (result []*Auth, err error) {
	sql := `select id,path,label,parent_id,is_page from authority where id in (
    			select auth_id from role_auth where role_id in (
    			    select role_id from user inner join user_role on user.id = user_role.user_id where user.id =?) group by auth_id
    		)`
	db, _ := DB.DB()
	rows, err := db.Query(sql, u.ID)
	for rows.Next() {
		item := Auth{}
		rows.Scan(&item.Id, &item.Path, &item.Label, &item.ParentId, &item.IsPage)
		result = append(result, &item)
	}
	return
}

// func UpdateUserRole(userId int, roles []int) error {
// 	return DB.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Where("user_id = ?", userId).Delete(UserRole{}).Error; err != nil {
// 			return err
// 		}
// 		userRoles := make([]UserRole, 0)
// 		for _, v := range roles {
// 			userRoles = append(userRoles, UserRole{
// 				RoleId: v,
// 				UserId: userId,
// 			})
// 		}
// 		if err := tx.Create(&userRoles).Error; err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
