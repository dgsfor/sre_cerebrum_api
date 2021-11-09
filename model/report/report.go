package report

import "ssopa/model"

// 报告表
type Report struct {
	model.BaseModel
	ReportId    string `json:"report_id"`                               // 报告id
	ReportName  string `json:"report_name"`                             // 报告名称
	TemplateId  string `json:"template_id"`                             // 模板id
	Content     string `json:"content" gorm:"type:TEXT;default:null"`   // markdown内容
	VarList     string `json:"var_list" gorm:"type:TEXT;default:null"`  // 变量json串
	SlotList    string `json:"slot_list" gorm:"type:TEXT;default:null"` // 批注插槽json串
	StartTime   string `json:"start_time" gorm:"default:null"`          // 开始时间
	EndTime     string `json:"end_time" gorm:"default:null"`            // 结束时间
	RenderId    string `json:"render_id" gorm:"default:null"`           // 渲染ID
	MergeStatus string `json:"merge_status" gorm:"default:'NotMerge'"`  // 合并状态，未合并(NotMerge)，已合并(MergeD)，合并失败(MergeError)
}

func (Report) TableName() string {
	return "report"
}
