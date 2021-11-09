package authority_message

import (
	"net/http"
	"ssopa/conf"
	"ssopa/middleware"
	"ssopa/model"
	"ssopa/model/authority_message"
	"ssopa/serializer"
	"ssopa/util"
	"strconv"
	"time"
)

type CreateNoticeChannelParams struct {
	Name       string `form:"name" json:"name" binding:"required"`             // 名称
	Type       int    `form:"type" json:"type" binding:"required"`             // 类型 , 1(企业微信应用)、2(企业微信app)
	ShareType  int    `form:"share_type" json:"share_type" binding:"required"` // 是否公开, 1(公开的)、2(私有的)
	RobotKey   string `form:"robot_key" json:"robot_key"`                      // 机器人key
	CorpId     string `form:"corp_id" json:"corp_id"`                          // 企业id
	CorpSecret string `form:"corp_secret" json:"corp_secret"`                  // 企业微信应用秘钥
	AgentId    string `form:"agent_id" json:"agent_id"`                        // 企业微信应用id
}

func (p *CreateNoticeChannelParams) CreateNoticeChannel(users *util.UserCookie) serializer.SsopaResponse {
	noticeChannelId := "nt-ch-" + util.RandStringRunes(5) + strconv.FormatInt(time.Now().Unix(), 10)
	record := conf.Orm.Where("name = ? and type = ?", p.Name,p.Type).Find(&authority_message.NoticeChannel{}).RowsAffected
	if record >= 1 {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_MODULE, "create", users.Name, users.Email, "report authority notice channel exist!", nil)
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: nil,
				Msg:  "渠道名已经存在，请检查！",
			},
			ResCode: serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_EXIST,
		}
	}
	authorityNoticeChannelModel := &authority_message.NoticeChannel{
		BaseModel:       model.BaseModel{},
		NoticeChannelId: noticeChannelId,
		Name:            p.Name,
		Type:            p.Type,
		ShareType:       p.ShareType,
		RobotKey:        p.RobotKey,
		CorpId:          p.CorpId,
		CorpSecret:      p.CorpSecret,
		AgentId:         p.AgentId,
		Creator:         users.Name,
	}
	err := conf.Orm.Create(&authorityNoticeChannelModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_MODULE, "create", users.Name, users.Email, "report authority notice channel create error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "新增通知渠道失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_CREATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_MODULE, "create", users.Name, users.Email, "report authority notice channel create success!", authorityNoticeChannelModel)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: authorityNoticeChannelModel,
			Msg:  "新增通知渠道成功！",
		},
		ResCode: serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_CREATE_SUCCESS,
	}
}

func GetNoticeChannelList(users *util.UserCookie) serializer.SsopaResponse {
	var noticeChannelModel []authority_message.NoticeChannel
	err := conf.Orm.Order("created_at desc").Find(&noticeChannelModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_MODULE, "get", users.Name, users.Email, "get report authority notice channel list error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取通知渠道列表失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_GET_LIST_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_TEMPLATE_MODULE, "get", users.Name, users.Email, "report authority notice channel list success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: noticeChannelModel,
			Msg:  "获取通知渠道列表成功！",
		},
		ResCode: serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_GET_LIST_SUCCESS,
	}
}
