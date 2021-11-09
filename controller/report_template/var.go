package report_template

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ssopa/serializer"
	"ssopa/service/report_template"
	"ssopa/util"
	"time"
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
	result := report_template.DeleteReportTemplateVar(users, templateId, varName)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 获取渲染进度
func GetRenderProgress(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	renderId := c.Param("render_id")
	resourceId := c.Param("resource_id")
	resourceType := c.Param("resource_type")
	result := report_template.GetRenderProgress(renderId, resourceId, resourceType, users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 合并变量记录到内容主体
func MergeRenderRecordToContent(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	svc := report_template.MergeRenderRecordToContentParams{}
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
	result := svc.MergeRenderRecordToContent(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

func GetDate(c *gin.Context) {
	time.Sleep(time.Duration(10) * time.Second)
	c.JSON(http.StatusOK, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: time.Now().Format("2006-01-02 15:04:05"),
			Msg:  "获取时间成功",
		},
		ResCode: serializer.REPORT_TEMPLATE_VAR_DELETE_SUCCESS,
	})
}

func GetDate2(c *gin.Context) {
	time.Sleep(time.Duration(20) * time.Second)
	c.JSON(http.StatusOK, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: time.Now().Format("2006-01-02 15:04:05"),
			Msg:  "获取时间成功",
		},
		ResCode: serializer.REPORT_TEMPLATE_VAR_DELETE_SUCCESS,
	})
}

func GetDate3(c *gin.Context) {
	time.Sleep(time.Duration(30) * time.Second)
	c.JSON(http.StatusOK, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: time.Now().Format("2006-01-02 15:04:05"),
			Msg:  "获取时间成功",
		},
		ResCode: serializer.REPORT_TEMPLATE_VAR_DELETE_SUCCESS,
	})
}
