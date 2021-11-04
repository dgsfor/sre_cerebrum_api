package report_template

import (
	"net/http"
	"ssopa/conf"
	"ssopa/middleware"
	"ssopa/model"
	reportTemplateModel "ssopa/model/report_template"
	"ssopa/serializer"
	"ssopa/util"
)

type CreateReportTemplateVarParams struct {
	TemplateId     string `form:"template_id" json:"template_id"`                              // 模板id
	VarName        string `form:"var_name" json:"var_name" binding:"required"`                 // 变量名称
	VarUrl         string `form:"var_url" json:"var_url" binding:"required"`                   // 变量获取地址
	VarHeader      string `form:"var_header" json:"var_header"`                                // 请求的header头
	VarResultField string `form:"var_result_field" json:"var_result_field" binding:"required"` // 变量结果获取字段
	VarType        string `form:"var_type" json:"var_type" binding:"required"`                 // 变量类型 ,内置变量、自定义变量、图片变量，inner_var、custom_var、img_var
	VarDesc        string `form:"var_desc" json:"var_desc" binding:"required"`                 // 变量备注，变量说明
}

func (p *CreateReportTemplateVarParams) CreateReportTemplateVar(users *util.UserCookie) serializer.SsopaResponse {
	var row int64
	if p.TemplateId != "" {
		row = conf.Orm.Where("template_id = ? and var_name = ?", p.TemplateId, p.VarName).Find(&reportTemplateModel.ReportTemplateVar{}).RowsAffected
	} else {
		row = conf.Orm.Where("var_name = ?", p.VarName).Find(&reportTemplateModel.ReportTemplateVar{}).RowsAffected
	}
	if row >= 1 {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "create", users.Name, users.Email, "report template var exist!", nil)
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: nil,
				Msg:  "变量已经存在，请检查！",
			},
			ResCode: serializer.REPORT_TEMPLATE_VAR_EXIST,
		}
	}
	reportTemplateVarModel := &reportTemplateModel.ReportTemplateVar{
		BaseModel:      model.BaseModel{},
		TemplateId:     p.TemplateId,
		VarName:        p.VarName,
		VarUrl:         p.VarUrl,
		VarHeader:      p.VarHeader,
		VarResultField: p.VarResultField,
		VarType:        p.VarType,
		VarDesc:        p.VarDesc,
		Creator:        users.Name,
	}
	err := conf.Orm.Create(&reportTemplateVarModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "create", users.Name, users.Email, "report template var create error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "新增变量注册失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_VAR_CREATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "create", users.Name, users.Email, "report template var create success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportTemplateVarModel,
			Msg:  "新增变量注册成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_VAR_CREATE_SUCCESS,
	}
}

func GetReportTemplateVarList(users *util.UserCookie) serializer.SsopaResponse {
	var reportTemplateVarModel []reportTemplateModel.ReportTemplateVar
	err := conf.Orm.Order("created_at desc").Find(&reportTemplateVarModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "get", users.Name, users.Email, "get report template var list error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取变量列表失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_VAR_GET_LIST_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "get", users.Name, users.Email, "report template var list success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportTemplateVarModel,
			Msg:  "获取变量列表成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_VAR_GET_LIST_SUCCESS,
	}
}

func DeleteReportTemplateVar(users *util.UserCookie, templateId string, varName string) serializer.SsopaResponse {
	var reportTemplateVarModel reportTemplateModel.ReportTemplateVar
	err := conf.Orm.Where("template_id = ? and var_name = ?", templateId,varName).Delete(&reportTemplateVarModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "delete", users.Name, users.Email, "delete report template var specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "删除变量失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_VAR_DELETE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "delete", users.Name, users.Email, "delete report template var specified success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "删除变量成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_VAR_DELETE_SUCCESS,
	}
}