package report_template

import "ssopa/model"

/**
变量表
1.变量默认使用get请求方式获取
2.变量分为：内置变量、自定义变量、图片变量
3.内置变量可以共用，不需要重复定义
4.图片变量目前只支持使用grafana
*/
type ReportTemplateVar struct {
	model.BaseModel
	TemplateId     string `json:"template_id" gorm:"default:null"` // 模板id
	VarName        string `json:"var_name"`                        // 变量名称
	VarUrl         string `json:"var_url" gorm:"type:TEXT"`        // 变量获取地址
	VarHeader      string `json:"var_header" gorm:"default:null"`  // 请求的header头
	VarResultField string `json:"var_result_field"`                // 变量结果获取字段
	VarType        string `json:"var_type"`                        // 变量类型 ,内置变量、自定义变量、图片变量，inner_var、custom_var、img_var
	VarDesc        string `json:"var_desc" gorm:"default:null"`    // 变量备注，变量说明
	Creator        string `json:"creator" gorm:"default:null"`     // 创建人
}

func (ReportTemplateVar) TableName() string {
	return "report_template_var"
}

/**
变量渲染状态表
该表记录每一个报告或者权威消息所有的变量的状态信息
每一个报告有多条变量数据
每一个权威消息有多条变量数据
未渲染：0
渲染成功：1
渲染失败：2
未注册：3
*/
type VarRenderedRecord struct {
	model.BaseModel
	ResourceId   string `json:"resource_id"`                              // 资源Id，可以是报告id、权威消息id
	VarName      string `json:"var_name"`                                 // 变量名称
	RenderStatus int    `json:"render_status" gorm:"default:0"`           // 渲染状态，未渲染、渲染成功、渲染失败
	VarResult    string `json:"var_result" gorm:"type:TEXT;default:null"` // 变量获取到的结果
	RenderId     string `json:"render_id" gorm:"default:null"`            // 渲染ID
	RenderTime   string `json:"render_time" gorm:"default:null"`          // 渲染时间
}

func (VarRenderedRecord) TableName() string {
	return "var_rendered_record"
}

/**
批注相关
*/
// 批注插槽注册表
type ReportTemplateSlot struct {
	model.BaseModel
	TemplateId    string `json:"template_id"`     // 模板id
	SlotName      string `json:"slot_name"`       // 批注插槽名称
	SlotNameAlias string `json:"slot_name_alias"` // 批注插槽别名
}

func (ReportTemplateSlot) TableName() string {
	return "report_template_slot"
}

// 批注表
type SlotAnnotate struct {
	model.BaseModel
	ReportId        string `json:"report_id"`                            // 报告id
	SlotName        string `json:"slot_name"`                            // 批注插槽名称
	SlotNameAlias   string `json:"slot_name_alias"`                      // 批注插槽别名
	AnnotateContent string `json:"annotate_content" gorm:"default:null"` // 插槽内容
	AnnotateTag     string `json:"annotate_tag" gorm:"default:null"`     // 批注标签
	Creator         string `json:"creator" gorm:"default:null"`          // 批注人
	SlotStatus      int64   `json:"slot_status" gorm:"default: 1"`    // 插槽状态，1 未注册，2 未批注，3 批注完成
}

func (SlotAnnotate) TableName() string {
	return "slot_annotate"
}
