package v1

import (
	"PalaemonBlog/model"
	"PalaemonBlog/utils/errormsg"
	"PalaemonBlog/utils/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QueryUserListReq struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}

var code int

// AddNewUser 添加新用户 | add new user
func AddNewUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Get JSON error:", err)
		return
	}
	valMsg, code := validate.Validator(&data)
	if code != errormsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": valMsg,
		})
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
	var req QueryUserListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println("Get QueryUserListReq error:", err)
		return
	}
	if req.PageSize == 0 {
		req.PageSize = -1
	}
	if req.PageNum == 0 {
		req.PageNum = -1
	}
	data, total := model.GetUserList(req.PageSize, req.PageNum)
	code = errormsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户 | edit user
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Get JSON error:", err)
		return
	}
	code := model.CheckUserStatus(data.Username)
	if code == errormsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errormsg.ErrorUserNameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户 | delete user
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})

}
