package report_template

import (
	"net/http"
	"ssopa/conf"
	"ssopa/middleware"
	"ssopa/model"
	"ssopa/model/report_template"
	"ssopa/serializer"
	"ssopa/util"
	"strconv"
	"time"
)

type CreateReportTemplateParams struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Status string `form:"status" json:"status" binding:"required"`
	Type   string `form:"type" json:"type" binding:"required"`
}

type UpdateReportTemplateParams struct {
	TemplateId string `form:"template_id" json:"template_id" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"` // markdown内容
	VarList    string `form:"var_list" json:"var_list"`                  // 变量json串
	SlotList   string `form:"slot_list" json:"slot_list"`                // 批注插槽json串
}

func (p *CreateReportTemplateParams) CreateReportTemplate(users *util.UserCookie) serializer.SsopaResponse {
	templateId := "rep-temp-" + util.RandStringRunes(5) + strconv.FormatInt(time.Now().Unix(), 10)
	record := conf.Orm.Where("name = ?", p.Name).Find(&report_template.ReportTemplate{}).RowsAffected
	if record >= 1 {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "create", users.Name, users.Email, "report template exist!", nil)
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: nil,
				Msg:  "模板名已经存在，请检查！",
			},
			ResCode: serializer.REPORT_TEMPLATE_EXIST,
		}
	}
	reportTemplateModel := &report_template.ReportTemplate{
		BaseModel:  model.BaseModel{},
		TemplateId: templateId,
		Name:       p.Name,
		Creator:    users.Name,
		Status:     p.Status,
		Type:       p.Type,
	}
	err := conf.Orm.Create(&reportTemplateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "create", users.Name, users.Email, "report template create error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "新增模板失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_CREATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "create", users.Name, users.Email, "report template create success!", reportTemplateModel)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportTemplateModel,
			Msg:  "新增模板成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_CREATE_SUCCESS,
	}
}

func GetReportTemplateList(users *util.UserCookie) serializer.SsopaResponse {
	var reportTemplateModel []report_template.ReportTemplate
	err := conf.Orm.Order("created_at desc").Find(&reportTemplateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "get", users.Name, users.Email, "get report template list error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取运营模板列表失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_GET_LIST_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "get", users.Name, users.Email, "report template list success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportTemplateModel,
			Msg:  "获取运营模板列表成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_GET_LIST_SUCCESS,
	}
}

func GetReportTemplate(users *util.UserCookie, templateId string) serializer.SsopaResponse {
	var reportTemplateModel report_template.ReportTemplate
	err := conf.Orm.Where("template_id = ?", templateId).Find(&reportTemplateModel).Error
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
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "get", users.Name, users.Email, "report template specified success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportTemplateModel,
			Msg:  "获取运营模板详情成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_GET_SUCCESS,
	}
}

func (p *UpdateReportTemplateParams) UpdateReportTemplate(users *util.UserCookie) serializer.SsopaResponse {
	var reportTemplateModel report_template.ReportTemplate
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
	reportTemplateModel.Content = p.Content
	reportTemplateModel.VarList = p.VarList
	reportTemplateModel.SlotList = p.SlotList
	err = conf.Orm.Save(&reportTemplateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "update", users.Name, users.Email, "update report template error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "更新运营模板失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_UPDATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "update", users.Name, users.Email, "update report template success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "更新运营模板成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_UPDATE_SUCCESS,
	}
}

func DeleteReportTemplate(users *util.UserCookie, templateId string) serializer.SsopaResponse {
	var reportTemplateModel report_template.ReportTemplate
	err := conf.Orm.Where("template_id = ?", templateId).Delete(&reportTemplateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "delete", users.Name, users.Email, "delete report template specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "删除运营模板失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_DELETE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "delete", users.Name, users.Email, "delete report template success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "删除运营模板成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_DELETE_SUCCESS,
	}
}

func UpdateReportTemplateStatus(users *util.UserCookie, templateId string, status string) serializer.SsopaResponse {
	var reportTemplateModel report_template.ReportTemplate
	err := conf.Orm.Where("template_id = ?", templateId).Find(&reportTemplateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "get", users.Name, users.Email, "get report template specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取运营模板失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_GET_ERROR,
		}
	}
	reportTemplateModel.Status = status
	err = conf.Orm.Save(&reportTemplateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "update", users.Name, users.Email, "update report template error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "更新运营模板失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_UPDATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "update", users.Name, users.Email, "update report template success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "更新运营模板成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_UPDATE_SUCCESS,
	}
}
