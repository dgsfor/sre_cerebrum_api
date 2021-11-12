package report_template

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"math"
	"net/http"
	"net/url"
	"ssopa/conf"
	"ssopa/middleware"
	"ssopa/model"
	authorityMessage "ssopa/model/authority_message"
	"ssopa/model/report"
	reportTemplateModel "ssopa/model/report_template"
	"ssopa/serializer"
	"ssopa/util"
	"strings"
	"time"
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

type CreateReportTemplateSlotParams struct {
	TemplateId    string `form:"template_id" json:"template_id"`                            // 模板id
	SlotName      string `form:"slot_name" json:"slot_name" binding:"required"`             // 插槽名称
	SlotNameAlias string `form:"slot_name_alias" json:"slot_name_alias" binding:"required"` // 插槽别名
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
	err := conf.Orm.Where("template_id = ? and var_name = ?", templateId, varName).Delete(&reportTemplateVarModel).Error
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

func (p *CreateReportTemplateSlotParams) CreateReportTemplateSlot(users *util.UserCookie) serializer.SsopaResponse {
	row := conf.Orm.Where("template_id = ? and slot_name = ?", p.TemplateId, p.SlotName).Find(&reportTemplateModel.ReportTemplateSlot{}).RowsAffected
	if row >= 1 {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "create", users.Name, users.Email, "report template var exist!", nil)
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: nil,
				Msg:  "插槽已经存在，请检查！",
			},
			ResCode: serializer.REPORT_TEMPLATE_SLOT_EXIST,
		}
	}
	reportTemplateSlotModel := &reportTemplateModel.ReportTemplateSlot{
		BaseModel:     model.BaseModel{},
		TemplateId:    p.TemplateId,
		SlotName:      p.SlotName,
		SlotNameAlias: p.SlotNameAlias,
	}
	err := conf.Orm.Create(&reportTemplateSlotModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "create", users.Name, users.Email, "report template slot create error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "新增插槽注册失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_SLOT_CREATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "create", users.Name, users.Email, "report template var create success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportTemplateSlotModel,
			Msg:  "新增插槽注册成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_SLOT_CREATE_SUCCESS,
	}
}

func GetReportTemplateSlotList(users *util.UserCookie) serializer.SsopaResponse {
	var reportTemplateSlotModel []reportTemplateModel.ReportTemplateSlot
	err := conf.Orm.Order("created_at desc").Find(&reportTemplateSlotModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "get", users.Name, users.Email, "get report template slot list error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取插槽列表失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_SLOT_GET_LIST_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "get", users.Name, users.Email, "report template slot list success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: reportTemplateSlotModel,
			Msg:  "获取插槽列表成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_SLOT_GET_LIST_SUCCESS,
	}
}

func DeleteReportTemplateSlot(users *util.UserCookie, templateId string, slotName string) serializer.SsopaResponse {
	var reportTemplateSlotModel reportTemplateModel.ReportTemplateSlot
	err := conf.Orm.Where("template_id = ? and slot_name = ?", templateId, slotName).Delete(&reportTemplateSlotModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "delete", users.Name, users.Email, "delete report template slot specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "删除插槽失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_SLOT_DELETE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "delete", users.Name, users.Email, "delete report template slot specified success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "删除插槽成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_SLOT_DELETE_SUCCESS,
	}
}


// 批量创建变量渲染记录
func GenerateVarRenderedRecord(resourceId string, varListString string) {
	varList := strings.Split(varListString, ",")
	for _, varName := range varList {
		// 如果这里使用go来做，可能会造成有些变量还没有创建出来，然后就已经到了渲染步骤，最终渲染失败
		GenerateVarRenderedRecordService(resourceId, varName)
	}
}
func GenerateVarRenderedRecordService(resourceId string, varName string) {
	varRenderRecordModel := &reportTemplateModel.VarRenderedRecord{
		BaseModel:  model.BaseModel{},
		ResourceId: resourceId,
		VarName:    varName,
	}
	err := conf.Orm.Create(&varRenderRecordModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "create", "auto", "auto", "create var rendered record error!", resourceId+":"+err.Error())
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "create", "auto", "auto", "create var rendered record success!", resourceId+":"+varName)
}
// 批量创建插槽
func GenerateSlot(templateId string, resourceId string, slotListString string) {
	slotList := strings.Split(slotListString, ",")
	for _, slotName := range slotList {
		// 如果这里使用go来做，可能会造成有些变量还没有创建出来，然后就已经到了渲染步骤，最终渲染失败
		GenerateSlotService(templateId, resourceId, slotName)
	}
}
func GenerateSlotService(templateId string, resourceId string, slotName string) {
	var slotStatus int64
	var reportTemplateSlotModel reportTemplateModel.ReportTemplateSlot
	row := conf.Orm.Where("template_id = ? and slot_name = ?",templateId, slotName).Find(&reportTemplateSlotModel).RowsAffected
	if row < 1 {
		slotStatus = 1
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "create", "auto", "auto", "get slot register error", resourceId+":"+slotName)
	} else {
		slotStatus = 2
	}
	slotAnnotateModel := &reportTemplateModel.SlotAnnotate{
		BaseModel:       model.BaseModel{},
		ReportId:        resourceId,
		SlotName:        slotName,
		SlotNameAlias:   reportTemplateSlotModel.SlotNameAlias,
		SlotStatus:      slotStatus,
	}
	err := conf.Orm.Create(&slotAnnotateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "create", "auto", "auto", "create slot error!", resourceId+":"+err.Error())
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_SLOT_MODULE, "create", "auto", "auto", "create slot success!", resourceId+":"+slotName)
}

/**
渲染变量
1.循环出变量
2.批量获取到变量值，然后修改变量渲染记录中的VarResult字段
3.这时候就可以实时获取渲染的进度
4.所有的编辑动作、发送消息动作、预览报告动作都会预先检查渲染进度
5.如果100%，那么就更新content
6.如果不是100%，可以选择强制更新content
*/
func RenderedVar(renderId string, resourceType string, resourceId string, templateId string, varListString string) {
	varList := strings.Split(varListString, ",")
	for _, varName := range varList {
		go RenderedVarService(renderId, resourceType, resourceId, templateId, varName)
	}
}
func RenderedVarService(renderId string, resourceType string, resourceId string, templateId string, varName string) {
	renderStartTime := time.Now()
	var varRenderRecordModel reportTemplateModel.VarRenderedRecord
	_ = conf.Orm.Where("resource_id = ? and var_name = ?", resourceId, varName).Find(&varRenderRecordModel).Error
	var reportTemplateVarModel reportTemplateModel.ReportTemplateVar
	row := conf.Orm.Where("template_id = ? and var_name = ?", templateId, varName).Find(&reportTemplateVarModel).RowsAffected
	if row == 0 {
		// 找不到变量注册，可能是内置变量
		middleware.RenderOutPutLog(renderId, serializer.REPORT_TEMPLATE_VAR_MODULE, "render", "auto", "auto", "render var error,can't not found var register,maybe inner var!",
			"resourceId: "+resourceId+",vaName: "+varName)
		row = conf.Orm.Where("var_name = ?", varName).Find(&reportTemplateVarModel).RowsAffected
		if row == 0 {
			// 内置变量也没有，直接退出不进行渲染
			middleware.RenderOutPutLog(renderId, serializer.REPORT_TEMPLATE_VAR_MODULE, "render", "auto", "auto", "render var error,can't not found var register,not inner var!",
				"resourceId: "+resourceId+",vaName: "+varName)
			varRenderRecordModel.RenderStatus = 3
			varRenderRecordModel.RenderId = renderId
			_ = conf.Orm.Save(&varRenderRecordModel).Error
			return
		} else {
			goto RenderFunc
		}
	}
	// get var data and write VarRenderRecord
	goto RenderFunc
RenderFunc:
	getStatus, dataResult := RenderVarServiceGetVarData(resourceId, resourceType, reportTemplateVarModel)
	if getStatus {
		middleware.RenderOutPutLog(renderId, serializer.REPORT_TEMPLATE_VAR_MODULE, "render", "auto", "auto", "render var , get var data success",
			"resourceId: "+resourceId+",vaName: "+varName+",newData:"+dataResult)
		varRenderRecordModel.RenderStatus = 1
	} else {
		middleware.RenderOutPutLog(renderId, serializer.REPORT_TEMPLATE_VAR_MODULE, "render", "auto", "auto", "render var , get var data failure",
			"resourceId: "+resourceId+",vaName: "+varName+",error:"+dataResult)
		varRenderRecordModel.RenderStatus = 2
	}
	varRenderRecordModel.VarResult = dataResult
	varRenderRecordModel.RenderId = renderId
	renderEndTime := time.Now()
	renderTime := renderEndTime.Sub(renderStartTime)
	varRenderRecordModel.RenderTime = renderTime.String()
	err := conf.Orm.Save(&varRenderRecordModel).Error
	if err != nil {
		middleware.RenderOutPutLog(renderId, serializer.REPORT_TEMPLATE_VAR_MODULE, "render", "auto", "auto", "render var , save result to varRenderRecord failure",
			"resourceId: "+resourceId+",vaName: "+varName+",error:"+err.Error())
	}
}
func RenderVarServiceGetVarData(resourceId string, resourceType string, p reportTemplateModel.ReportTemplateVar) (bool, string) {
	//inner_var、custom_var、img_var
	// 增加header
	headerMap := make(map[string]string)
	headerMap["content-type"] = "application/json"
	if p.VarHeader != "" {
		varHeaderList := strings.Split(p.VarHeader, ",")
		for _, singleHeader := range varHeaderList {
			varHeaderKey := strings.Split(singleHeader, ":")[0]
			varHeaderValue := strings.Split(singleHeader, ":")[1]
			headerMap[varHeaderKey] = varHeaderValue
		}
	}
	if p.VarType == "inner_var" || p.VarType == "custom_var" {
		var requestUrl string
		if resourceType == serializer.REPORT_AUTHORITY_MESSAGE_MODULE {
			requestUrl = fmt.Sprintf("%s", p.VarUrl)
		}
		if resourceType == serializer.REPORT_MODULE {
			var reportModel report.Report
			_ = conf.Orm.Where("report_id = ?", resourceId).Find(&reportModel).Error
			requestUrl = fmt.Sprintf("%s?start_time=%s&end_time=%s", p.VarUrl, url.QueryEscape(reportModel.StartTime), url.QueryEscape(reportModel.EndTime))
		}
		response, err := util.HandlerRequest("GET", requestUrl, headerMap, nil)
		jsonBody, errNewJson := simplejson.NewJson(response)
		if err != nil || errNewJson != nil || jsonBody == nil {
			return false, "get var data error"
		}
		stringValue := jsonBody.Get(p.VarResultField).MustString()
		floatValue := jsonBody.Get(p.VarResultField).MustFloat64()
		if stringValue == "" {
			return true, fmt.Sprintf("%f", floatValue)
		} else {
			return true, stringValue
		}
	} else if p.VarType == "img_var" {
		return true, "this is image url!"
	} else {
		return false, "not found var type!"
	}
}

/**
获取渲染进度
*/
func GetRenderProgress(renderId string, resourceId string, resourceType string, users *util.UserCookie) serializer.SsopaResponse {
	var varRenderRecordModel []reportTemplateModel.VarRenderedRecord
	var varRenderRecordSuccessCount int64
	var varNum int64
	result := make(map[string]interface{})
	row := conf.Orm.Where("render_id = ?", renderId).Find(&varRenderRecordModel).RowsAffected
	if row <= 0 {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "get", users.Name, users.Email, "when get render progress,get var render record error", nil)
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: "未找到render_id对应的记录",
				Msg:  "获取变量渲染记录列表失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_VAR_RENDER_RECORD_GET_LIST_ERROR,
		}
	}
	result["render_record"] = varRenderRecordModel
	if resourceType == serializer.REPORT_AUTHORITY_MESSAGE_MODULE {
		var authorityMessageModel authorityMessage.AuthorityMessage
		_ = conf.Orm.Where("message_id = ?", resourceId).Find(&authorityMessageModel).Error
		varNum = int64(len(strings.Split(authorityMessageModel.VarList, ",")))
	}
	if resourceType == serializer.REPORT_MODULE {
		var reportModel report.Report
		_ = conf.Orm.Where("report_id = ?", resourceId).Find(&reportModel).Error
		varNum = int64(len(strings.Split(reportModel.VarList, ",")))
	}
	err := conf.Orm.Where("render_id = ? and render_status = ?", renderId, 1).Find(&varRenderRecordModel).Count(&varRenderRecordSuccessCount).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "get", users.Name, users.Email, "when get render progress,get var render record error", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取变量渲染记录列表失败！",
			},
			ResCode: serializer.REPORT_TEMPLATE_VAR_RENDER_RECORD_GET_LIST_ERROR,
		}
	}
	per := (float64(varRenderRecordSuccessCount) / float64(varNum)) * 100
	result["render_progress"] = math.Ceil(per)
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "get", users.Name, users.Email, "get render progress success", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: result,
			Msg:  "获取渲染进度成功！",
		},
		ResCode: serializer.REPORT_TEMPLATE_VAR_RENDER_PROGRESS_GET_SUCCESS,
	}

}

/**
合并渲染记录到内容主体
*/
type MergeRenderRecordToContentParams struct {
	ResourceId   string `form:"resource_id" json:"resource_id" binding:"required"`     // 资源id
	ResourceType string `form:"resource_type" json:"resource_type" binding:"required"` // 资源类型
	RenderId     string `form:"render_id" json:"render_id" binding:"required"`         // 渲染id
}

func (p *MergeRenderRecordToContentParams) MergeRenderRecordToContent(users *util.UserCookie) serializer.SsopaResponse {
	var varRenderRecordModel []reportTemplateModel.VarRenderedRecord
	err := conf.Orm.Where("render_id = ? and resource_id = ?", p.RenderId, p.ResourceId).Find(&varRenderRecordModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "merge", users.Name, users.Email, "when merge var render record to content, get record error", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "合并渲染记录到主体失败",
			},
			ResCode: serializer.REPORT_TEMPLATE_VAR_MERGE_ERROR,
		}
	}
	if p.ResourceType == serializer.REPORT_AUTHORITY_MESSAGE_MODULE {
		var authorityMessageModel authorityMessage.AuthorityMessage
		err := conf.Orm.Where("message_id = ?", p.ResourceId).Find(&authorityMessageModel).Error
		if authorityMessageModel.MergeStatus == "Merged" {
			goto SuccessMergeFunc
		}
		if err != nil {
			middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "merge", users.Name, users.Email, "when merge var render record to content, get authority record error", err.Error())
			return serializer.SsopaResponse{
				Response: serializer.Response{
					Code: http.StatusInternalServerError,
					Data: err.Error(),
					Msg:  "合并渲染记录到主体失败",
				},
				ResCode: serializer.REPORT_AUTHORITY_MESSAGE_MERGE_ERROR,
			}
		}
		for _, record := range varRenderRecordModel {
			authorityMessageModel.Content = RenderVarServiceReplaceVarData(authorityMessageModel.Content, record.VarName, record.VarResult, 1)
		}
		authorityMessageModel.MergeStatus = "MergeD"
		err = conf.Orm.Save(&authorityMessageModel).Error
		if err != nil {
			return serializer.SsopaResponse{
				Response: serializer.Response{
					Code: http.StatusInternalServerError,
					Data: err.Error(),
					Msg:  "合并渲染记录到主体失败",
				},
				ResCode: serializer.REPORT_AUTHORITY_MESSAGE_MERGE_ERROR,
			}
		}

	}
	if p.ResourceType == serializer.REPORT_MODULE {
		var reportModel report.Report
		err := conf.Orm.Where("report_id = ?", p.ResourceId).Find(&reportModel).Error
		if reportModel.MergeStatus == "Merged" {
			goto SuccessMergeFunc
		}
		if err != nil {
			middleware.CustomOutPutLog(serializer.REPORT_MODULE, "merge", users.Name, users.Email, "when merge var render record to content, get report record error", err.Error())
			return serializer.SsopaResponse{
				Response: serializer.Response{
					Code: http.StatusInternalServerError,
					Data: err.Error(),
					Msg:  "合并渲染记录到主体失败",
				},
				ResCode: serializer.REPORT_MERGE_ERROR,
			}
		}
		for _, record := range varRenderRecordModel {
			reportModel.Content = RenderVarServiceReplaceVarData(reportModel.Content, record.VarName, record.VarResult, 1)
		}
		reportModel.MergeStatus = "MergeD"
		reportModel.Status = "ToBeLabeled"
		err = conf.Orm.Save(&reportModel).Error
		if err != nil {
			return serializer.SsopaResponse{
				Response: serializer.Response{
					Code: http.StatusInternalServerError,
					Data: err.Error(),
					Msg:  "合并渲染记录到主体失败",
				},
				ResCode: serializer.REPORT_MERGE_ERROR,
			}
		}
	}
SuccessMergeFunc:
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_VAR_MODULE, "merge", users.Name, users.Email, "when merge var render record to content, get record error", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: "ok",
			Msg:  "合并渲染记录到主体成功",
		},
		ResCode: serializer.REPORT_TEMPLATE_VAR_MERGE_SUCCESS,
	}
}
func RenderVarServiceReplaceVarData(content string, oldStr string, newStr string, count int) string {
	return strings.Replace(content, oldStr, newStr, count)
}
