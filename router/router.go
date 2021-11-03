package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	oauthController "ssopa/controller/oauth"
	reportTemplateController "ssopa/controller/report_template"
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
		ur.POST("/login", oauthController.Login)
		ur.POST("/register", oauthController.Register)
		ur.GET("/logout", oauthController.Logout)
		ur.GET("/userinfo",atLeastLoginMiddleware, oauthController.GetUserInfo)
		ur.GET("/check_login", atLeastLoginMiddleware, oauthController.CheckLogin)
		// report template
		reportTemplate := v1.Group("/report_template")
		reportTemplate.POST("/template",atLeastLoginMiddleware, reportTemplateController.CreateReportTemplate)
		reportTemplate.GET("/template",atLeastLoginMiddleware, reportTemplateController.GetReportTemplateList)
		reportTemplate.PUT("/template",atLeastLoginMiddleware, reportTemplateController.UpdateReportTemplate)
		reportTemplate.GET("/template/:template_id",atLeastLoginMiddleware, reportTemplateController.GetReportTemplate)
	}
	return r
}
