package report

import (
	"github.com/google/uuid"
	"net/http"
	"ssopa/conf"
	"ssopa/middleware"
	"ssopa/model"
	"ssopa/model/report"
	reportTemplate "ssopa/model/report_template"
	"ssopa/serializer"
	"ssopa/service/report_template"
	"ssopa/util"
	"strconv"
	"time"
)

type CreateReportParams struct {
	TemplateId string `form:"template_id" json:"template_id" binding:"required"` // 模板id
	Name       string `form:"name" json:"name" binding:"required"`               // 名称
	ReportType string `form:"report_type" json:"report_type" binding:"required"` // 报告类型
	StartTime  string `form:"start_time" json:"start_time" binding:"required"`   // 开始时间
	EndTime    string `form:"end_time" json:"end_time" binding:"required"`       // 结束时间
}

type UpdateReportParams struct {
	ReportId string `form:"report_id" json:"report_id" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"` // markdown内容
}

// 创建报告
func (p *CreateReportParams) CreateReport(users *util.UserCookie) serializer.SsopaResponse {
	var reportTemplateModel reportTemplate.ReportTemplate
	err := conf.Orm.Where("template_id = ?", p.TemplateId).Find(&reportTemplateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "get", users.Name, users.Email, "get report template specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取运营模板详情失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_GET_ERROR,
		}
	}
	reportId := "report-" + util.RandStringRunes(5) + strconv.FormatInt(time.Now().Unix(), 10)
	previewHash, _ := uuid.NewRandom()
	renderId, _ := uuid.NewRandom()
	reportModel := &report.Report{
		BaseModel:   model.BaseModel{},
		ReportId:    reportId,
		Name:        p.Name,
		TemplateId:  p.TemplateId,
		ReportType:  p.ReportType,
		Content:     reportTemplateModel.Content,
		VarList:     reportTemplateModel.VarList,
		SlotList:    reportTemplateModel.SlotList,
		StartTime:   p.StartTime,
		EndTime:     p.EndTime,
		PreviewHash: previewHash.String(),
		RenderId:    renderId.String(),
		Creator:     users.Name,
	}
	reportRsModel := &report.Replica{
		BaseModel: model.BaseModel{},
		ReportId:  reportId,
		Content:   reportTemplateModel.Content,
	}
	_ = conf.Orm.Create(&reportRsModel).Error
	err = conf.Orm.Create(&reportModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_MODULE, "create", users.Name, users.Email, "create report error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "创建报告失败！",
			},
			ResCode: serializer.REPORT_CREATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_MODULE, "create", users.Name, users.Email, "start to create slot record", nil)
	report_template.GenerateSlot(reportTemplateModel.TemplateId, reportId, reportTemplateModel.SlotList)
	middleware.CustomOutPutLog(serializer.REPORT_MODULE, "create", users.Name, users.Email, "create report success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "创建报告成功",
		},
		ResCode: serializer.REPORT_CREATE_SUCCESS,
	}

}

// 报告列表
func GetReportList(users *util.UserCookie) serializer.SsopaResponse {
	var reportModel []report.Report
	err := conf.Orm.Order("created_at desc").Find(&reportModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_MODULE, "get", users.Name, users.Email, "get report list error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取报告列表失败！",
			},
			ResCode: serializer.REPORT_GET_LIST_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_MODULE, "get", users.Name, users.Email, "get report list success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportModel,
			Msg:  "获取报告列表成功！",
		},
		ResCode: serializer.REPORT_GET_LIST_SUCCESS,
	}
}

// 渲染报告
func RenderReport(users *util.UserCookie, reportId string) serializer.SsopaResponse {
	var reportModel report.Report
	err := conf.Orm.Where("report_id = ?", reportId).Find(&reportModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_MODULE, "render", users.Name, users.Email, "get report list error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取报告列表失败！",
			},
			ResCode: serializer.REPORT_GET_LIST_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_MODULE, "create", users.Name, users.Email, "start to create var rendered record", nil)
	report_template.GenerateVarRenderedRecord(reportId, reportModel.VarList)
	middleware.CustomOutPutLog(serializer.REPORT_MODULE, "render", users.Name, users.Email, "start to rendered var", nil)
	report_template.RenderedVar(reportModel.RenderId, serializer.REPORT_MODULE, reportId, reportModel.TemplateId, reportModel.VarList)
	reportModel.Status = "RenderIng"
	err = conf.Orm.Save(&reportModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_MODULE, "render", users.Name, users.Email, "save report status error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "渲染报告时，保存报告状态失败！",
			},
			ResCode: serializer.REPORT_UPDATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_MODULE, "render", users.Name, users.Email, "render report success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportModel,
			Msg:  "渲染报告成功！",
		},
		ResCode: serializer.REPORT_RENDER_SUCCESS,
	}
}

// 获取报告详情
func GetReport(users *util.UserCookie, reportId string) serializer.SsopaResponse {
	var reportModel report.Report
	err := conf.Orm.Where("report_id = ?", reportId).Find(&reportModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_MODULE, "get", users.Name, users.Email, "get report specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取运营报告详情失败！",
			},
			ResCode: serializer.REPORT_GET_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_MODULE, "get", users.Name, users.Email, "report specified success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportModel,
			Msg:  "获取运营报告详情成功！",
		},
		ResCode: serializer.REPORT_GET_SUCCESS,
	}
}

// 更新报告内容

func (p *UpdateReportParams) UpdateReport(users *util.UserCookie) serializer.SsopaResponse {
	var reportModel report.Report
	err := conf.Orm.Where("report_id = ?", p.ReportId).Find(&reportModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_MODULE, "get", users.Name, users.Email, "get report specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取运营报告详情失败！",
			},
			ResCode: serializer.REPORT_GET_ERROR,
		}
	}
	reportModel.Content = p.Content
	err = conf.Orm.Save(&reportModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_MODULE, "update", users.Name, users.Email, "update report error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "更新运营报告失败！",
			},
			ResCode: serializer.REPORT_UPDATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_MODULE, "update", users.Name, users.Email, "update report success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "更新运营报告成功！",
		},
		ResCode: serializer.REPORT_UPDATE_SUCCESS,
	}
}

// 预览报告
func Preview(reportId string, previewHash string) serializer.SsopaResponse {
	var reportModel report.Report
	err := conf.Orm.Where("report_id = ? and preview_hash = ?", reportId, previewHash).Find(&reportModel).Error
	if err != nil {
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取运营报告详情失败！",
			},
			ResCode: serializer.REPORT_GET_ERROR,
		}
	}
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportModel,
			Msg:  "获取运营报告详情成功！",
		},
		ResCode: serializer.REPORT_GET_SUCCESS,
	}
}

// 完结报告
func FinishReport(users *util.UserCookie, reportId string) serializer.SsopaResponse {
	var reportModel report.Report
	err := conf.Orm.Where("report_id = ?", reportId).Find(&reportModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_MODULE, "update", users.Name, users.Email, "get report list error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取报告列表失败！",
			},
			ResCode: serializer.REPORT_GET_LIST_ERROR,
		}
	}
	reportModel.Status = "Published"
	err = conf.Orm.Save(&reportModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_MODULE, "update", users.Name, users.Email, "save report status error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "完结报告时，保存报告状态失败！",
			},
			ResCode: serializer.REPORT_UPDATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_MODULE, "update", users.Name, users.Email, "finish report success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportModel,
			Msg:  "完结报告成功！",
		},
		ResCode: serializer.REPORT_FINISH_SUCCESS,
	}
}

// 获取报告涉及到的所有批注列表
func GetReportSlotAnnotateList(users *util.UserCookie, reportId string) serializer.SsopaResponse {
	var slotAnnotateModel []reportTemplate.SlotAnnotate
	err := conf.Orm.Where("report_id = ?", reportId).Find(&slotAnnotateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "get", users.Name, users.Email, "get report slot specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取运营报告批注列表失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_SLOT_GET_LIST_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "get", users.Name, users.Email, "get report slot specified success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: slotAnnotateModel,
			Msg:  "获取运营报告批注列表成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_SLOT_GET_LIST_SUCCESS,
	}
}