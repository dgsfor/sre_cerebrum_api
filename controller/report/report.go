package report

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ssopa/serializer"
	"ssopa/service/report"
	"ssopa/util"
)

// 创建报告
func CreateReport(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	svc := report.CreateReportParams{}
	if err := c.ShouldBind(&svc); err != nil {
		c.JSON(http.StatusBadRequest, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusBadRequest,
				Data: err.Error(),
				Msg:  "参数错误，请检查！",
			},
			ResCode: 40000,
		})
		return
	}
	result := svc.CreateReport(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 获取报告列表
func GetReportList(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	result := report.GetReportList(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 渲染报告
func RenderReport(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	reportId := c.Param("report_id")
	result := report.RenderReport(users,reportId)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}