package v1

import (
	"PalaemonBlog/model"
	"PalaemonBlog/utils/errormsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QueryArticleListReq struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}

// AddNewArticle 添加文章 | add new article
func AddNewArticle(c *gin.Context) {
	var data model.Article
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Get JSON error:", err)
		return
	}
	code := model.CreateNewArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

// QueryArticlesByCategory 查询分类下的所有文章 | query all articles under category
func QueryArticlesByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req QueryArticleListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println("Get Query ArticleReq error:", err)
		return
	}
	if req.PageSize == 0 {
		req.PageSize = -1
	}
	if req.PageNum == 0 {
		req.PageNum = -1
	}
	data, code, total := model.QueryArticlesByCategory(id, req.PageSize, req.PageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormsg.GetErrMsg(code),
	})
}

// QuerySingleArticle 查询文章 | query article
func QuerySingleArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.QuerySingleArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

// QueryArticleList 查询文章列表 | query article list
func QueryArticleList(c *gin.Context) {
	var req QueryArticleListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println("Get Query ArticleReq error:", err)
		return
	}
	if req.PageSize == 0 {
		req.PageSize = -1
	}
	if req.PageNum == 0 {
		req.PageNum = -1
	}
	data, code, total := model.QueryArticleList(req.PageSize, req.PageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormsg.GetErrMsg(code),
	})
}

// EditArticle 编辑文章 | edit article
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Get JSON error:", err)
		return
	}
	code = model.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})

}

// DeleteArticle 删除文章 | delete article
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})

}
