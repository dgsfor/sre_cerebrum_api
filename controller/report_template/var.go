package report_template

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ssopa/serializer"
	"ssopa/service/report_template"
	"ssopa/util"
)

// 新增运营模板
func CreateReportTemplateVar(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	svc := report_template.CreateReportTemplateVarParams{}
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
	result := svc.CreateReportTemplateVar(users)
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
func GetReportTemplateVarList(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	result := report_template.GetReportTemplateVarList(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 删除指定变量
func DeleteReportTemplateVar(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	templateId := c.Param("template_id")
	varName := c.Param("var_name")
	result := report_template.DeleteReportTemplateVar(users,templateId,varName)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}