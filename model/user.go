package model

import (
	"PalaemonBlog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" `
	Password string `gorm:"type:varchar(20);not null" json:"password" `
	Role     int    `gorm:"type:int" json:"role" `
}

// 查询用户状态 | query user status
func CheckUserStatus(name string) (code int) {
	var user User
	Db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		//fmt.Println("user.ID = ", user.ID)
		return errormsg.ErrorUserNameUsed //1001
	}
	return errormsg.SUCCESS
}

// CreateNewUser 添加新用户 | add new user
func CreateNewUser(data *User) int {
	err := Db.Create(&data)
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 查询单个用户 | query user

// 查询用户列表 | query user list
func QueryUserList(c *gin.Context) {

}

// 编辑用户 | edit user
func EditUser(c *gin.Context) {

}

// 删除用户 | delete user
func DeleteUser(c *gin.Context) {

}
