package authority_message

import "ssopa/model"

// 权威消息表
type AuthorityMessage struct {
	model.BaseModel
	MessageId          string `json:"message_id"`                                     // 消息id
	TemplateId         string `json:"template_id"`                                    // 模板id
	Name               string `json:"name"`                                           // 消息名称
	SendStatus         string `json:"send_status" gorm:"default:'NotSent'"`           // 发送状态,NotSent(未发送)、Sent(已发送)、Fail(发送失败)
	AuditStatus        string `json:"audit_status" gorm:"default:'NotApproved'"`      // 审核状态,Reviewed(已审核)、NotApproved(未审核)、Refuse(拒绝)、Disuse(废弃)、无需审核(NoNeedAudit)
	Creator            string `json:"creator"`                                        // 创建人
	Content            string `json:"content" gorm:"type:TEXT;default:null"`          // 消息内容
	NoticeChannelId    string `json:"notice_channel_id"`                              // 使用的渠道id
	Reviewer           string `json:"reviewer" gorm:"default:null"`                   // 审核人员
	MessageType        string `json:"message_type" gorm:"default:'custom'"`           // 消息类型，custom、custom_oncall、holidays_oncall,srp_message,fault_message
	AuditRequired      string `json:"audit_required" gorm:"default:'Need'"`           // 是否需要审核，Need、DontNeed
	VarList            string `json:"var_list" gorm:"type:TEXT;default:null"`         // 变量json串
	MessageContentType string `json:"message_content_type" gorm:"default:'markdown'"` // 内容类型，markdown类型，text类型
	RenderId           string `json:"render_id" gorm:"default:null"`                  // 渲染ID
	MergeStatus        string `json:"merge_status" gorm:"default:'NotMerge'"`         // 合并状态，未合并(NotMerge)，已合并(MergeD)，合并失败(MergeError)，合并中(MergeIng)
}

/**
custom 普通类型消息
custom_oncall 日常SRE轮值
holidays_oncall 节假日值班
srp_message 稳定性运营报告消息
fault_message 故障同步消息
*/

func (AuthorityMessage) TableName() string {
	return "authority_message"
}

// 消息发送历史
type AuthorityMessageSendHistory struct {
	model.BaseModel
	MessageId  string `json:"message_id"`                           // 消息id
	SendStatus string `json:"send_status" gorm:"default:'NotSent'"` // 发送状态,NotSent(未发送)、Sent(已发送)、Fail(发送失败)
	SendType   string `json:"send_type"`                            // 发送类型，测试发送、正式发送
	Creator    string `json:"creator"`                              // 创建人
}

func (AuthorityMessageSendHistory) TableName() string {
	return "authority_message_send_history"
}