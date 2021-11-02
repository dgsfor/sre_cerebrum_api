package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	oauth_controller "ssopa/controller/oauth"
	"ssopa/middleware"
	"time"
)

var (
	Logger, _ = zap.NewProduction()
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	atLeastLoginMiddleware := middleware.MiddlewareAuthorization()
	r := gin.Default()
	r.Use(middleware.Cors())
	// 关闭gin默认AccessLog日志
	if gin.Mode() == "release" {
		gin.DefaultWriter = ioutil.Discard
		r.Use(middleware.GinZap(Logger,time.RFC3339,true))
		r.Use(gin.Recovery())
	}
	// 路由
	v1 := r.Group("/api/ssopa/v1")
	{
		// user
		ur := v1.Group("/oauth")
		ur.POST("/login", oauth_controller.Login)
		ur.POST("/register",oauth_controller.Register)
		ur.GET("/logout", oauth_controller.Logout)
		ur.GET("/userinfo",atLeastLoginMiddleware,oauth_controller.GetUserInfo)
		ur.GET("/check_login", atLeastLoginMiddleware, oauth_controller.CheckLogin)
	}
	return r
}
