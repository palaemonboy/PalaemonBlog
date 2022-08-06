package middleware

import (
	"PalaemonBlog/utils"
	"PalaemonBlog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/errgo.v2/errors"
	"net/http"
	"strings"
	"time"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

var (
	TokenExpired     = errors.New("Token Expired")
	TokenNotValidYet = errors.New("Token Not ValidYet")
	TokenMalformed   = errors.New("Token Malformed")
	TokenInvalid     = errors.New("Token Invalid")
)

type MyClaims struct {
	jwt.RegisteredClaims
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

// SetToken 生成token | generate token
func (j *JWT) SetToken(UserName string) (Token string, ErrorCode int) {
	ExpireTime := jwt.NewNumericDate(jwt.TimeFunc().Add(10 * time.Hour))
	SetClaims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: ExpireTime,
			Issuer:    "PalaemonBlog",
		},
		UserName: UserName,
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(j.JwtKey)
	if err != nil {
		return "", errormsg.ERROR
	}
	return token, errormsg.SUCCESS
}

// VerifyToken 验证 token
func (j *JWT) VerifyToken(Token string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(Token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil {
		if v, ok := err.(jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid

}

// JwtToken jwt 中间件 | jwt middleware
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": errormsg.ErrorTokenExist,
				"msg":  errormsg.GetErrMsg(errormsg.ErrorTokenExist),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(authHeader, " ", 2)
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"code": errormsg.ErrorTokenWrong,
				"msg":  errormsg.GetErrMsg(errormsg.ErrorTokenWrong),
			})
			c.Abort()
			return
		}
		//get correct token
		j := NewJWT()
		claims, err := j.VerifyToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code": errormsg.ErrorTokenRuntime,
					"msg":  errormsg.GetErrMsg(errormsg.ErrorTokenRuntime),
					"data": nil,
				})
				c.Abort()
				return
			}
			// other errors
			c.JSON(http.StatusOK, gin.H{
				"code": errormsg.ERROR,
				"msg":  err.Error(),
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Set("user_name", claims)
		c.Next()
	}
}
