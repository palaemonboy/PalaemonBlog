package v1

import (
	"PalaemonBlog/model"
	"PalaemonBlog/utils/errormsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var code int

// AddNewUser 添加新用户 | add new user
func AddNewUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Get JSON error:", err)
		return
	}
	code = model.CheckUserStatus(data.Username)
	if code == errormsg.SUCCESS {
		model.CreateNewUser(&data)
	}
	if code == errormsg.ErrorUserNameUsed {
		code = errormsg.ErrorUserNameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

// 查询单个用户 | query user

// QueryUserList 查询用户列表 | query user list
func QueryUserList(c *gin.Context) {

}

// EditUser 编辑用户 | edit user
func EditUser(c *gin.Context) {

}

// DeleteUser 删除用户 | delete user
func DeleteUser(c *gin.Context) {

}
