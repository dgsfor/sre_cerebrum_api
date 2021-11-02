package middleware

import (
	"github.com/gin-gonic/gin"
	"ssopa/serializer"
	"net/http"
	"ssopa/util"
)

// 登录认证拦截器
func MiddlewareAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := util.GetUserCookie(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, serializer.Response{
				Code:  http.StatusUnauthorized,
				Data:  nil,
				Msg:   "认证失败",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
