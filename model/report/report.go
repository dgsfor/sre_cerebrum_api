package report

import "ssopa/model"

// 报告表
type Report struct {
	model.BaseModel
	ReportId    string `json:"report_id"`                               // 报告id
	Name        string `json:"name"`                                    // 报告名称
	TemplateId  string `json:"template_id"`                             // 模板id
	ReportType  string `json:"report_type"`                             // 报告类型，比如时报，日报，周报，月报，年报，巡检报(hourly,daily,weekly,monthly,yearly,inspection)
	Status      string `json:"status" gorm:"default:'NotRendered'"`     // 报告状态 未渲染、渲染中、待标注、已发布(NotRendered、RenderIng、ToBeLabeled、Published)
	Content     string `json:"content" gorm:"type:TEXT;default:null"`   // markdown内容
	VarList     string `json:"var_list" gorm:"type:TEXT;default:null"`  // 变量json串
	SlotList    string `json:"slot_list" gorm:"type:TEXT;default:null"` // 批注插槽json串
	StartTime   string `json:"start_time" gorm:"default:null"`          // 开始时间
	EndTime     string `json:"end_time" gorm:"default:null"`            // 结束时间
	PreviewHash string `json:"preview_hash"`                            // 预览hash链接，用做路由
	RenderId    string `json:"render_id" gorm:"default:null"`           // 渲染ID
	MergeStatus string `json:"merge_status" gorm:"default:'NotMerge'"`  // 合并状态，未合并(NotMerge)，已合并(MergeD)，合并失败(MergeError)
	Creator     string `json:"creator"`                                 // 创建人
}

func (Report) TableName() string {
	return "report"
}

// 作为report的一个副本，只保存内容
type Replica struct {
	model.BaseModel
	ReportId string `json:"report_id"`                             // 报告id
	Content  string `json:"content" gorm:"type:TEXT;default:null"` // markdown内容
}

func (Replica) TableName() string {
	return "report_replica"
}
