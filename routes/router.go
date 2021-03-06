package routes

import (
	v1 "PalaemonBlog/api/v1"
	"PalaemonBlog/model"
	"PalaemonBlog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	// init DB | 初始化DB连接
	if err := model.InitDB(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		panic(err)
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Welcome Palaemon Blog Server.",
		})
	})
	// API入口
	API := router.Group("/api")
	// Service 入口
	V1API := API.Group("/v1")
	{
		// user module routers
		V1API.POST("/user/add", v1.AddNewUser)
		V1API.GET("/users", v1.QueryUserList)
		V1API.PUT("/user/:id", v1.EditUser)
		V1API.DELETE("/user/:id", v1.DeleteUser)

		// category module routers
		V1API.POST("/category/add", v1.AddNewCategory)
		V1API.GET("/category/:id", v1.QuerySingleCategory)
		V1API.GET("/categorys", v1.QueryCategoryList)
		V1API.PUT("/category/:id", v1.EditCategory)
		V1API.DELETE("/category/:id", v1.DeleteCategory)

		// article module routers
		V1API.POST("/article/add", v1.AddNewArticle)
		V1API.GET("/article/:id", v1.QuerySingleArticle)
		V1API.GET("/article/category/:id", v1.QueryArticlesByCategory)
		V1API.GET("/articles", v1.QueryArticleList)
		V1API.PUT("/article/:id", v1.EditArticle)
		V1API.DELETE("/article/:id", v1.DeleteArticle)

	}
	err := router.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("Run Error.")
	}
}
