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
		//【测试】检查用户状态

		// user module routers
		V1API.POST("/user/add", v1.AddNewUser)
		V1API.GET("/users", v1.QueryUserList)
		V1API.PUT("/user/:id", v1.EditUser)
		V1API.DELETE("/user/:id", v1.DeleteUser)

		// category module routers

		// article module routers

	}
	err := router.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("Run Error.")
	}
}
