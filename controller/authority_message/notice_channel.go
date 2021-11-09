package authority_message

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ssopa/serializer"
	"ssopa/service/authority_message"
	"ssopa/util"
)

// 新增通知渠道
func CreateNoticeChannel(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	svc := authority_message.CreateNoticeChannelParams{}
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
	result := svc.CreateNoticeChannel(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// 获取所有通知渠道列表
func GetNoticeChannelList(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	result := authority_message.GetNoticeChannelList(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}
