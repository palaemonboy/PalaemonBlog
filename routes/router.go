package routes

import (
	"PalaemonBlog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Welcome Palaemon Blog Server.",
		})
	})
	// API入口
	API := router.Group("/api")
	// Service 入口
	ServiceAPI := API.Group("/v1")
	{
		ServiceAPI.GET("/start", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Palaemon Blog Running.",
			})
		})

	}
	err := router.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("Run Error.")
	}
}
