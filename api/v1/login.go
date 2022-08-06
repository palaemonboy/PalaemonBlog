package v1

import (
	"PalaemonBlog/middleware"
	"PalaemonBlog/model"
	"PalaemonBlog/utils/errormsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Get JSON error:", err)
		return
	}
	user, code := model.CheckLogin(data.Username, data.Password)
	if code == errormsg.SUCCESS {
		j := middleware.NewJWT()
		token, setTokenErrCode := j.SetToken(data.Username)
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    user.Username,
			"id":      user.ID,
			"message": errormsg.GetErrMsg(setTokenErrCode),
			"token":   token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  errormsg.ErrorTokenWrong,
			"data":    data.Username,
			"id":      data.ID,
			"message": errormsg.GetErrMsg(errormsg.ErrorTokenWrong),
		})
	}

}
