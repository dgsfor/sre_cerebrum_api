package authority_message

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"ssopa/conf"
	"ssopa/middleware"
	"ssopa/model"
	authorityMessage "ssopa/model/authority_message"
	reportTemplate "ssopa/model/report_template"
	"ssopa/serializer"
	"ssopa/service/report_template"
	"ssopa/util"
	"strconv"
	"strings"
	"time"
)

type CreateAuthorityMessageParams struct {
	Name               string `form:"name" json:"name" binding:"required"`                           // 消息名称
	TemplateId         string `form:"template_id" json:"template_id" binding:"required"`             // 消息模板
	NoticeChannelId    string `form:"notice_channel_id" json:"notice_channel_id" binding:"required"` // 发送渠道
	MessageType        string `form:"message_type" json:"message_type"`                              // 消息类型
	AuditRequired      string `form:"audit_required" json:"audit_required"`                          // 是否审核
	MessageContentType string `form:"message_content_type" json:"message_content_type"`              // 内容格式
}

type UpdateAuthorityMessageParams struct {
	MessageId string `form:"message_id" json:"message_id" binding:"required"` //  消息id
	Content   string `form:"content" json:"content" binding:"required"`       // markdown内容
}

// 创建消息
func (p *CreateAuthorityMessageParams) CreateAuthorityMessage(users *util.UserCookie) serializer.SsopaResponse {
	var reportTemplateModel reportTemplate.ReportTemplate
	err := conf.Orm.Where("template_id = ?", p.TemplateId).Find(&reportTemplateModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "create", users.Name, users.Email, "get report template error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusOK,
				Data: err.Error(),
				Msg:  "所用模板可能不存在，请检查！",
			},
			ResCode: serializer.REPORT_TEMPLATE_GET_ERROR,
		}
	}
	renderId, _ := uuid.NewRandom()
	authorityMessageId := "msg-" + util.RandStringRunes(5) + strconv.FormatInt(time.Now().Unix(), 10)
	var AuditStatus string
	if p.AuditRequired == "false" {
		AuditStatus = "NoNeedAudit"
	} else if p.AuditRequired == "true" {
		AuditStatus = "NotApproved"
	}
	authorityMessageModel := &authorityMessage.AuthorityMessage{
		BaseModel:          model.BaseModel{},
		MessageId:          authorityMessageId,
		TemplateId:         p.TemplateId,
		Name:               p.Name,
		Creator:            users.Name,
		Content:            reportTemplateModel.Content,
		AuditStatus:        AuditStatus,
		NoticeChannelId:    p.NoticeChannelId,
		MessageType:        p.MessageType,
		AuditRequired:      p.AuditRequired,
		VarList:            reportTemplateModel.VarList,
		MessageContentType: p.MessageContentType,
		RenderId:           renderId.String(),
	}
	err = conf.Orm.Create(&authorityMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "create", users.Name, users.Email, "create authority message error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "创建权威消息失败!",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_CREATE_ERROR,
		}
	}
	if reportTemplateModel.VarList != "" {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "create", users.Name, users.Email, "start to create var rendered record", nil)
		report_template.GenerateVarRenderedRecord(authorityMessageId, reportTemplateModel.VarList)
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "create", users.Name, users.Email, "start to rendered var ", nil)
		report_template.RenderedVar(renderId.String(), serializer.REPORT_AUTHORITY_MESSAGE_MODULE, authorityMessageId, p.TemplateId, reportTemplateModel.VarList)
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "create", users.Name, users.Email, "create authority message success!", nil)
	}
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "创建权威消息成功",
		},
		ResCode: serializer.REPORT_AUTHORITY_MESSAGE_CREATE_SUCCESS,
	}
}

// 权威消息列表
func GetAuthorityMessageList(users *util.UserCookie) serializer.SsopaResponse {
	var authorityMessageModel []authorityMessage.AuthorityMessage
	err := conf.Orm.Order("created_at desc").Find(&authorityMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "get", users.Name, users.Email, "get authority message list error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取权威消息列表失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_GET_LIST_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "get", users.Name, users.Email, "get authority message list success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: authorityMessageModel,
			Msg:  "获取权威消息列表成功！",
		},
		ResCode: serializer.REPORT_AUTHORITY_MESSAGE_GET_LIST_SUCCESS,
	}
}

// 获取权威消息
func GetAuthorityMessage(users *util.UserCookie, messageId string) serializer.SsopaResponse {
	var authorityMessageModel authorityMessage.AuthorityMessage
	err := conf.Orm.Where("message_id = ?", messageId).Find(&authorityMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "get", users.Name, users.Email, "get authority message specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取权威消息详情失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_GET_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "get", users.Name, users.Email, "report authority message specified success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: authorityMessageModel,
			Msg:  "获取权威消息详情成功！",
		},
		ResCode: serializer.REPORT_AUTHORITY_MESSAGE_GET_SUCCESS,
	}
}

// 编辑权威消息
func (p *UpdateAuthorityMessageParams) UpdateAuthorityMessage(users *util.UserCookie) serializer.SsopaResponse {
	var authorityMessageModel authorityMessage.AuthorityMessage
	err := conf.Orm.Where("message_id = ?", p.MessageId).Find(&authorityMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "update", users.Name, users.Email, "get authority message specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取权威消息详情失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_GET_ERROR,
		}
	}
	authorityMessageModel.Content = p.Content
	err = conf.Orm.Save(&authorityMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "update", users.Name, users.Email, "update report template error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "更新权威消息失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_UPDATE_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "update", users.Name, users.Email, "update report template success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "更新权威消息成功！",
		},
		ResCode: serializer.REPORT_AUTHORITY_MESSAGE_UPDATE_SUCCESS,
	}
}

// 发送消息
func AuthorityMessageSend(users *util.UserCookie, messageId string, sendType string) serializer.SsopaResponse {
	var authorityMessageModel authorityMessage.AuthorityMessage
	err := conf.Orm.Where("message_id = ?", messageId).Find(&authorityMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "update", users.Name, users.Email, "get authority message specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取权威消息详情失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_GET_ERROR,
		}
	}
	var noticeChannelModel []authorityMessage.NoticeChannel
	err = conf.Orm.Where("notice_channel_id in (?)", strings.Split(authorityMessageModel.NoticeChannelId, ",")).Find(&noticeChannelModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_MODULE, "get", users.Name, users.Email, "get notice channel error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取渠道信息失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_NOTICE_CHANNEL_GET_LIST_ERROR,
		}
	}
	var status bool
	if sendType == "test" {
		robotKey := conf.GetConfig("notice::test_send_robot").String()
		content := authorityMessageModel.Content + "\n" + fmt.Sprintf("消息类型：测试，发送人：%s(正式发送不会有这一行)", users.Name)
		status, _ = util.NoticeToQywxRobot(authorityMessageModel.MessageContentType, content, robotKey)
	} else if sendType == "release" {
		for _, noticeChannel := range noticeChannelModel {
			if noticeChannel.Type == 1 {
				robotKey := noticeChannel.RobotKey
				content := authorityMessageModel.Content
				status, _ = util.NoticeToQywxRobot(authorityMessageModel.MessageContentType, content, robotKey)
			} else if noticeChannel.Type == 2 {
				corpId := noticeChannel.CorpId
				corpSecret := noticeChannel.CorpSecret
				agentId := noticeChannel.AgentId
				content := authorityMessageModel.Content
				accessToken := util.GetQywxAppAccessToken(corpId, corpSecret)
				status, _ = util.NoticeToQywxApp(accessToken, agentId, authorityMessageModel.MessageContentType, content, "@all")
			}
		}
	}
	if !status {
		authorityMessageModel.SendStatus = "Fail"
	} else {
		authorityMessageModel.SendStatus = "Sent"
	}
	if sendType == "release" {
		_ = conf.Orm.Save(&authorityMessageModel).Error
	}
	/**
	创建发送记录
	*/
	authorityMessageSendHistoryModel := &authorityMessage.AuthorityMessageSendHistory{
		BaseModel:  model.BaseModel{},
		MessageId:  messageId,
		SendStatus: authorityMessageModel.SendStatus,
		SendType:   sendType,
		Creator:    users.Name,
	}
	err = conf.Orm.Create(&authorityMessageSendHistoryModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "get", users.Name, users.Email, "crate send record error", err.Error())
	}
	middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "sent", users.Name, users.Email, "send message success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "发送消息成功！",
		},
		ResCode: serializer.REPORT_AUTHORITY_MESSAGE_SEND_SUCCESS,
	}
}

// 查看消息发送历史
func GetAuthorityMessageSendHistory(users *util.UserCookie, messageId string) serializer.SsopaResponse {
	var authorityMessageSendHistoryModel []authorityMessage.AuthorityMessageSendHistory
	err := conf.Orm.Where("message_id = ?", messageId).Find(&authorityMessageSendHistoryModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "update", users.Name, users.Email, "get authority message send history error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取发送历史失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_GET_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "update", users.Name, users.Email, "update report template success!", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: authorityMessageSendHistoryModel,
			Msg:  "获取发送历史成功！",
		},
		ResCode: serializer.REPORT_AUTHORITY_MESSAGE_GET_ERROR,
	}
}

// 审核消息
func AuditAuthorityMessage(users *util.UserCookie, messageId string,auditStatus string) serializer.SsopaResponse {
	var authorityMessageModel authorityMessage.AuthorityMessage
	err := conf.Orm.Where("message_id = ?", messageId).Find(&authorityMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "audit", users.Name, users.Email, "audit authority message specified error!", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "获取权威消息详情失败！",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_AUDIT_ERROR,
		}
	}
	if authorityMessageModel.AuditStatus != "NotApproved" {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "audit", users.Name, users.Email, "when audit, authority message auditStatus not 'NotApproved'", nil)
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: "权威消息审核状态不是未审核！",
				Msg:  "审核权威消息失败",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_AUDIT_ERROR,
		}
	}
	if auditStatus == "Disuse" {
		goto AuditFunc
	} else {
		reviewer := strings.Split(conf.GetConfig("notice::message_audit_member").String(), ",")
		isReviewer := false
		for key := range reviewer {
			if reviewer[key] == users.Email {
				isReviewer = true
			}
		}
		if !isReviewer {
			middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "audit", users.Name, users.Email, "not audit member", nil)
			return serializer.SsopaResponse{
				Response: serializer.Response{
					Code: http.StatusInternalServerError,
					Data: "你不是审核人员！",
					Msg:  "审核权威消息失败",
				},
				ResCode: serializer.REPORT_AUTHORITY_MESSAGE_AUDIT_ERROR,
			}
		}
		goto AuditFunc
	}
AuditFunc:
	authorityMessageModel.AuditStatus = auditStatus
	authorityMessageModel.Reviewer = users.Name
	err = conf.Orm.Save(&authorityMessageModel).Error
	if err != nil {
		middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "audit", users.Name, users.Email, "when audit,update message error", err.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "审核权威消息失败",
			},
			ResCode: serializer.REPORT_AUTHORITY_MESSAGE_AUDIT_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.REPORT_AUTHORITY_MESSAGE_MODULE, "audit", users.Name, users.Email, "audit authority message success", nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "审核权威消息成功",
		},
		ResCode: serializer.REPORT_AUTHORITY_MESSAGE_AUDIT_SUCCESS,
	}
}