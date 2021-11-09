package authority_message

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ssopa/serializer"
	"ssopa/service/authority_message"
	"ssopa/util"
)

// 创建权威消息
func CreateAuthorityMessage(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	svc := authority_message.CreateAuthorityMessageParams{}
	if err := c.ShouldBind(&svc); err != nil {
		c.JSON(http.StatusBadRequest, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusBadRequest,
				Data: err.Error(),
				Msg:  "参数错误，请检查！",
			},
			ResCode: 40000,
		})
		return
	}
	result := svc.CreateAuthorityMessage(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 获取所有权威消息列表
func GetAuthorityMessageList(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	result := authority_message.GetAuthorityMessageList(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 获取指定权威消息
func GetAuthorityMessage(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	messageId := c.Param("message_id")
	result := authority_message.GetAuthorityMessage(users, messageId)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 更新权威消息
func UpdateAuthorityMessage(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	svc := authority_message.UpdateAuthorityMessageParams{}
	if err := c.ShouldBind(&svc); err != nil {
		c.JSON(http.StatusBadRequest, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusBadRequest,
				Data: err.Error(),
				Msg:  "参数错误，请检查！",
			},
			ResCode: 40000,
		})
		return
	}
	result := svc.UpdateAuthorityMessage(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 发送权威消息
func AuthorityMessageSend(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	messageId := c.Param("message_id")
	sendType := c.Param("send_type")
	result := authority_message.AuthorityMessageSend(users, messageId, sendType)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 获取发送历史
func GetAuthorityMessageSendHistory(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	messageId := c.Param("message_id")
	result := authority_message.GetAuthorityMessageSendHistory(users, messageId)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 获取发送历史
func AuditAuthorityMessage(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	messageId := c.Param("message_id")
	auditStatus := c.Param("audit_status")
	result := authority_message.AuditAuthorityMessage(users, messageId, auditStatus)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}
