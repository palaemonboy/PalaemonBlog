package routes

import (
	v1 "PalaemonBlog/api/v1"
	"PalaemonBlog/middleware"
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
	router := gin.New()
	router.Use(
		middleware.Logger(),
		gin.Recovery(),
		middleware.Cors(),
	)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Welcome Palaemon Blog Server.",
		})
	})
	// API入口
	API := router.Group("/api")
	// Service 入口
	auth := API.Group("/v1")
	auth.Use(middleware.JwtToken())
	{
		// user module routers
		auth.GET("/users", v1.QueryUserList)
		auth.PUT("/user/:id", v1.EditUser)
		auth.DELETE("/user/:id", v1.DeleteUser)

		// category module routers
		auth.POST("/category/add", v1.AddNewCategory)

		auth.PUT("/category/:id", v1.EditCategory)
		auth.DELETE("/category/:id", v1.DeleteCategory)

		// article module routers
		auth.POST("/article/add", v1.AddNewArticle)

		auth.PUT("/article/:id", v1.EditArticle)
		auth.DELETE("/article/:id", v1.DeleteArticle)

	}
	public := API.Group("/v1")
	{
		public.POST("/login", v1.Login)

		public.POST("/user/add", v1.AddNewUser)

		public.GET("/category/:id", v1.QuerySingleCategory)
		public.GET("/categorys", v1.QueryCategoryList)

		public.GET("/article/:id", v1.QuerySingleArticle)
		public.GET("/article/category/:id", v1.QueryArticlesByCategory)
		public.GET("/articles", v1.QueryArticleList)
	}

	err := router.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("Run Error.")
	}
}
