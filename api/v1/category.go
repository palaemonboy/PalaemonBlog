package v1

import (
	"PalaemonBlog/model"
	"PalaemonBlog/utils/errormsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QueryCategoryListReq struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}

// AddNewCategory 添加分类 | add new category
func AddNewCategory(c *gin.Context) {
	var data model.Category
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Get JSON error:", err)
		return
	}
	code = model.CheckCategoryStatus(data.Name)
	if code == errormsg.SUCCESS {
		model.CreateNewCategory(&data)
	}
	if code == errormsg.ErrorCategoryUsed {
		code = errormsg.ErrorCategoryUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

// QuerySingleCategory 查询单个分类下的文章 | query article under category
func QuerySingleCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.QuerySingleCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

// QueryCategoryList 查询分类列表 | query category list
func QueryCategoryList(c *gin.Context) {
	var req QueryCategoryListReq
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
	data, total := model.QueryCategoryList(req.PageSize, req.PageNum)
	code = errormsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormsg.GetErrMsg(code),
	})
}

// EditCategory 编辑分类 | edit category
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Get JSON error:", err)
		return
	}
	code := model.CheckUserStatus(data.Name)
	if code == errormsg.SUCCESS {
		model.EditCategory(id, &data)
	}
	if code == errormsg.ErrorCategoryUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

// DeleteCategory 删除分类 | delete category
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})

}
