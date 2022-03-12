package instant_chat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ssopa/serializer"
	"ssopa/service/instant_chat"
	"ssopa/util"
)

// CreateChatMessage 新增私信消息
func CreateChatMessage(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	svc := instant_chat.CreateChatMessageParams{}
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
	result := svc.CreateChatMessage(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// GetNoticeChannelList 获取所有私信列表
func GetInstantChatMessageList(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	result := instant_chat.GetInstantChatMessageList(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}
