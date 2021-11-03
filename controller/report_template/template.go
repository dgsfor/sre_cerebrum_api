package report_template

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ssopa/serializer"
	"ssopa/service/report_template"
	"ssopa/util"
)

// 新增运营模板
func CreateReportTemplate(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	svc := report_template.CreateReportTemplateParams{}
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
	result := svc.CreateReportTemplate(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 获取所有运营模板
func GetReportTemplateList(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	result := report_template.GetReportTemplateList(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 获取指定运营模板
func GetReportTemplate(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	templateId := c.Param("template_id")
	result := report_template.GetReportTemplate(users,templateId)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 更新运营模板
func UpdateReportTemplate(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	svc := report_template.UpdateReportTemplateParams{}
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
	result := svc.UpdateReportTemplate(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 删除指定运营模板
func DeleteReportTemplate(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	templateId := c.Param("template_id")
	result := report_template.DeleteReportTemplate(users,templateId)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 更新指定运营模板状态
func UpdateReportTemplateStatus(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	templateId := c.Param("template_id")
	status := c.Param("status")
	result := report_template.UpdateReportTemplateStatus(users,templateId,status)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}