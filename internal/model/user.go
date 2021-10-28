/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-16 15:03:02
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-28 17:42:09
 */
package model

import "gorm.io/gorm"

type User struct {
	ID       uint32 `json:"id"`
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
	err := DB.Table(u.TabName()).Select("role_id").Where("user_id = ?", u.UserId).Find(&result).Error
	return result, err
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
