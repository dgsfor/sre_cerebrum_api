package report_template

import "ssopa/model"

// 模板表
type ReportTemplate struct {
	model.BaseModel
	TemplateId string `json:"template_id"`                             // 模板id
	Name       string `json:"name"`                                    // 模板名称
	Creator    string `json:"creator"`                                 // 创建人
	Status     string `json:"status"`                                  // 模板状态 禁用，可用 enable disable
	Type       string `json:"type"`                                    // 模板类型，开放模板(open_temp)，周期模板(cron_temp)，权威消息模板(authoritative_temp)
	Content    string `json:"content" gorm:"type:TEXT;default:null"`   // markdown内容
	VarList    string `json:"var_list" gorm:"type:TEXT;default:null"`  // 变量json串
	SlotList   string `json:"slot_list" gorm:"type:TEXT;default:null"` // 批注插槽json串
}

func (ReportTemplate) TableName() string {
	return "report_template"
}
