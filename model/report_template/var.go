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
