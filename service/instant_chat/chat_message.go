package instant_chat

import (
	"github.com/google/uuid"
	"net/http"
	"ssopa/conf"
	"ssopa/middleware"
	"ssopa/model"
	"ssopa/model/instant_chat"
	"ssopa/serializer"
	"ssopa/util"
	"strconv"
	"time"
)

type CreateChatMessageParams struct {
	Receiver      string `form:"receiver" json:"receiver" binding:"required"`             // 接收人
	ReceiverEmail string `form:"receiver_email" json:"receiver_email" binding:"required"` // 接收人邮件
	Content       string `form:"content" json:"content" binding:"required"`               // 消息内容
}

func (p *CreateChatMessageParams) CreateChatMessage(users *util.UserCookie) serializer.SsopaResponse {
	if users.Email == p.ReceiverEmail {
		middleware.CustomOutPutLog(serializer.INSTNAT_CHAT_MESSAGE_MODULE, "create", users.Name, users.Email, "can't message yourself", nil)
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: nil,
				Msg:  "不能发消息给自己！",
			},
			ResCode: serializer.INSTNAT_CHAT_MESSAGE_CREATE_ERROR,
		}
	}
	chatMessageId := "nt-ch-" + util.RandStringRunes(5) + strconv.FormatInt(time.Now().Unix(), 10)
	roomId, _ := uuid.NewUUID()
	instantChatMessageModel := &instant_chat.InstantChatMessage{
		BaseModel:     model.BaseModel{},
		ChatMessageId: chatMessageId,
		RoomId:        roomId.String(),
		Sender:        users.Name,
		SenderEmail:   users.Email,
		Receiver:      p.Receiver,
		ReceiverEmail: p.ReceiverEmail,
		Content:       p.Content,
		Type:          "PrivateMessage",
	}
	err := conf.Orm.Create(&instantChatMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.INSTNAT_CHAT_MESSAGE_MODULE, "create", users.Name, users.Email, "instant chat message create error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "创建私信失败！",
			},
			ResCode: serializer.INSTNAT_CHAT_MESSAGE_CREATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.INSTNAT_CHAT_MESSAGE_MODULE, "create", users.Name, users.Email, "instant chat message create success!", instantChatMessageModel)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: instantChatMessageModel,
			Msg:  "创建私信成功！",
		},
		ResCode: serializer.INSTNAT_CHAT_MESSAGE_CREATE_SUCCESS,
	}
}

func GetInstantChatMessageList(users *util.UserCookie) serializer.SsopaResponse {
	var instantChatMessageModel []instant_chat.InstantChatMessage
	err := conf.Orm.Where("sender_email = ? or receiver_email = ?", users.Email, users.Email).Order("created_at desc").Group("room_id").Find(&instantChatMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.INSTNAT_CHAT_MESSAGE_MODULE, "get", users.Name, users.Email, "get instant chat message list error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取私信消息列表失败！",
			},
			ResCode: serializer.INSTNAT_CHAT_MESSAGE_GET_LIST_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.INSTNAT_CHAT_MESSAGE_MODULE, "get", users.Name, users.Email, "get instant chat message list success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: instantChatMessageModel,
			Msg:  "获取私信消息列表成功！",
		},
		ResCode: serializer.INSTNAT_CHAT_MESSAGE_GET_LIST_SUCCESS,
	}
}
