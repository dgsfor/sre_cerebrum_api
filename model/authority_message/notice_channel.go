package authority_message

import "ssopa/model"

// 通知渠道表
type NoticeChannel struct {
	model.BaseModel
	NoticeChannelId string `json:"notice_channel_id"`                // 渠道id
	Name            string `json:"name"`                             // 名称
	Type            int    `json:"type"`                             // 类型， 1(企业微信机器人)、2(企业微信app)
	ShareType       int    `json:"share_type"`                       // 是否公开，1(公开的)、2(私有的)
	RobotKey        string `json:"robot_key" gorm:"default:null"`    // 机器人key
	CorpId          string `json:"corp_id"  gorm:"default:null"`     // 企业id
	CorpSecret      string `json:"corp_secret"  gorm:"default:null"` // 企业微信应用秘钥
	AgentId         string `json:"agent_id" gorm:"default:null"`     // 企业微信应用id
	Creator         string `json:"creator"`                          // 创建人
}

func (NoticeChannel) TableName() string {
	return "authority_notice_channel"
}
