package model

import (
	"PalaemonBlog/utils/errormsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"权限"`
}

// CheckUserStatus 查询用户状态 | query user status
func CheckUserStatus(name string) (code int) {
	var user User
	Db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errormsg.ErrorUserNameUsed //1001
	}
	return errormsg.SUCCESS
}

// CreateNewUser 添加新用户 | add new user
func CreateNewUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := Db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 查询单个用户 | query user

// GetUserList 查询用户列表 | query user list
func GetUserList(PageSize int, PageNum int) ([]User, int64) {
	var users []User
	var total int64
	err = Db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&users).Count(&total).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return nil, 0
	}
	return users, total
}

// EditUser 编辑用户 | edit user
func EditUser(ID int, data *User) (code int) {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = Db.Model(&user).Where("id = ?", ID).Updates(maps).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// DeleteUser 删除用户 | delete user
func DeleteUser(ID int) (code int) {
	var user User
	err := Db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// ScryptPw 密码加密 | Password encryption
func ScryptPw(password string) string {
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	keyLen := 10
	pwHash, err := scrypt.Key([]byte(password), salt, 1024, 8, 1, keyLen)
	if err != nil {
		log.Fatal("Password encryption error", err)
	}
	finalPw := base64.StdEncoding.EncodeToString(pwHash)
	return finalPw
}

// CheckLogin 查询登录状态| Check login status
func CheckLogin(username, password string) (User, int) {
	var user User
	Db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return user, errormsg.ErrorUserNotExist
	}
	if user.Password != ScryptPw(password) {
		return user, errormsg.ErrorPasswordWrong
	}
	// no admin permission
	if user.Role != 1 {
		return user, errormsg.ErrorUserNoPermission
	}
	return user, errormsg.SUCCESS

}
