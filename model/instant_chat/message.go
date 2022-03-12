package instant_chat

import "ssopa/model"

/**
暂时不做多人对话。
那么当A用户登录后，就获取sender = A or receiver = A的消息列表即可
*/

type InstantChatMessage struct {
	model.BaseModel
	ChatMessageId string `json:"chat_message_id"`                       // 即时通信消息id
	RoomId        string `json:"room_id"`                               // 房间id，用于关联一次聊天消息的上下文，roomid相同的可以认为是一次完整地消息对话，可用uuid来做
	Sender        string `json:"sender" gorm:"default:null"`            // 发送人
	SenderEmail   string `json:"sender_email" gorm:"default:null"`      // 发送人邮件
	Receiver      string `json:"receiver"`                              // 接收人
	ReceiverEmail string `json:"receiver_email"`                        // 接收人邮件
	Content       string `json:"content" gorm:"type:TEXT;default:null"` // 消息内容
	Annex         string `json:"annex" gorm:"default:null"`             // 附件，比如图片之类的消息
	Type          string `json:"type" gorm:"default:'InSiteMessage'"`   // 消息类型，站内消息(InSiteMessage)、私信(PrivateMessage)
	Status        string `json:"status" gorm:"default:'Unread'"`        // 状态，已读(Read)、未读(Unread)
}

func (InstantChatMessage) TableName() string {
	return "instant_chat_message"
}
