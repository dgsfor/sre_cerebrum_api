package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	authorityMessageController "ssopa/controller/authority_message"
	oauthController "ssopa/controller/oauth"
	reportController "ssopa/controller/report"
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
		r.Use(middleware.GinZap(Logger, time.RFC3339, true))
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
		ur.GET("/userinfo", atLeastLoginMiddleware, oauthController.GetUserInfo)
		ur.GET("/check_login", atLeastLoginMiddleware, oauthController.CheckLogin)
		// report template
		reportTemplate := v1.Group("/report_template")
		reportTemplate.POST("/template", atLeastLoginMiddleware, reportTemplateController.CreateReportTemplate)
		reportTemplate.GET("/template", atLeastLoginMiddleware, reportTemplateController.GetReportTemplateList)
		reportTemplate.PUT("/template", atLeastLoginMiddleware, reportTemplateController.UpdateReportTemplate)
		reportTemplate.GET("/template/:template_id", atLeastLoginMiddleware, reportTemplateController.GetReportTemplate)
		reportTemplate.DELETE("/template/:template_id", atLeastLoginMiddleware, reportTemplateController.DeleteReportTemplate)
		reportTemplate.PUT("/template/:template_id/:status", atLeastLoginMiddleware, reportTemplateController.UpdateReportTemplateStatus)
		// report template var
		reportTemplate.POST("/var", atLeastLoginMiddleware, reportTemplateController.CreateReportTemplateVar)
		reportTemplate.PUT("/var/merge", atLeastLoginMiddleware, reportTemplateController.MergeRenderRecordToContent)
		reportTemplate.GET("/var", atLeastLoginMiddleware, reportTemplateController.GetReportTemplateVarList)
		reportTemplate.DELETE("/var/:template_id/:var_name", atLeastLoginMiddleware, reportTemplateController.DeleteReportTemplateVar)
		reportTemplate.GET("/var/render/:render_id/:resource_id/:resource_type", atLeastLoginMiddleware, reportTemplateController.GetRenderProgress)
		reportTemplate.GET("/date", reportTemplateController.GetDate)
		reportTemplate.GET("/date2", reportTemplateController.GetDate2)
		reportTemplate.GET("/date3", reportTemplateController.GetDate3)
		// report template slot
		reportTemplate.POST("/slot", atLeastLoginMiddleware, reportTemplateController.CreateReportTemplateSlot)
		reportTemplate.GET("/slot", atLeastLoginMiddleware, reportTemplateController.GetReportTemplateSlotList)
		reportTemplate.DELETE("/slot/:template_id/:slot_name", atLeastLoginMiddleware, reportTemplateController.DeleteReportTemplateSlot)
		// report
		report := v1.Group("/report")
		report.POST("/report",atLeastLoginMiddleware,reportController.CreateReport)
		report.GET("/report",atLeastLoginMiddleware,reportController.GetReportList)
		report.GET("/report/:report_id",atLeastLoginMiddleware,reportController.GetReport)
		report.GET("/slot/:report_id",atLeastLoginMiddleware,reportController.GetReportSlotAnnotateList)
		report.PUT("/report",atLeastLoginMiddleware,reportController.UpdateReport)
		report.GET("/report_render/:report_id",atLeastLoginMiddleware,reportController.RenderReport)
		report.GET("/report_preview/:report_id/:preview_hash",reportController.Preview)
		report.PUT("/report_finish/:report_id",atLeastLoginMiddleware,reportController.FinishReport)
		// authority message
		authorityMessage := v1.Group("/authority_message")
		authorityMessage.POST("/notice_channel", atLeastLoginMiddleware, authorityMessageController.CreateNoticeChannel)
		authorityMessage.GET("/notice_channel", atLeastLoginMiddleware, authorityMessageController.GetNoticeChannelList)
		authorityMessage.GET("/message", atLeastLoginMiddleware, authorityMessageController.GetAuthorityMessageList)
		authorityMessage.PUT("/message", atLeastLoginMiddleware, authorityMessageController.UpdateAuthorityMessage)
		authorityMessage.GET("/message/:message_id", atLeastLoginMiddleware, authorityMessageController.GetAuthorityMessage)
		authorityMessage.POST("/message", atLeastLoginMiddleware, authorityMessageController.CreateAuthorityMessage)
		authorityMessage.GET("/message_send/:message_id/:send_type", atLeastLoginMiddleware, authorityMessageController.AuthorityMessageSend)
		authorityMessage.PUT("/message_audit/:message_id/:audit_status", atLeastLoginMiddleware, authorityMessageController.AuditAuthorityMessage)
		authorityMessage.GET("/message_history/:message_id", atLeastLoginMiddleware, authorityMessageController.GetAuthorityMessageSendHistory)
	}
	return r
}
