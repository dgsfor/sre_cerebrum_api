package event_orch

import "ssopa/model"

// 事件模板表
type OrchestrationTemplate struct {
	model.BaseModel
}

// 事件模板 -- 阶段表
type OrchestrationTemplateStage struct {
	model.BaseModel
}

// 事件模板 -- 任务表
type OrchestrationTemplateTask struct {
	model.BaseModel
}

func (OrchestrationTemplate) TableName() string {
	return "event_orchestration_template"
}
func (OrchestrationTemplateStage) TableName() string {
	return "event_orchestration_template_stage"
}
func (OrchestrationTemplateTask) TableName() string {
	return "event_orchestration_template_task"
}