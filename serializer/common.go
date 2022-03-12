package serializer

// Response 基础序列化器
// omitempty如果返回的数据这个字段为空，则序列化出来的数据没有这个字段
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

type SsopaResponse struct {
	Response
	ResCode int `json:"rescode"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// user相关
	USER_MODULE                                         = "user"
	CREATE_USER_SUCCESS                                 = 10000 // 创建用户成功
	LOGIN_SUCCESS                                       = 10001 // 登录成功
	USER_EXISTS                                         = 10002 // 用户存在
	CREATE_USER_ERROR                                   = 10003 // 创建用户失败
	USER_NOT_EXISTS                                     = 10004 // 用户不存在
	PASSWORD_ERROR                                      = 10005 // 密码错误
	SETCOOKIE_FAILURE                                   = 10006 // 设置cookie失败
	LOGOUT_SUCCESS                                      = 10007 // 退出登录成功
	USER_GET_LIST_ERROR                                 = 10008
	USER_GET_LIST_SUCCESS                               = 10009
	REPORT_TEMPLATE_MODULE                              = "report_template"
	REPORT_TEMPLATE_CREATE_SUCCESS                      = 20001
	REPORT_TEMPLATE_CREATE_ERROR                        = 20002
	REPORT_TEMPLATE_EXIST                               = 20003
	REPORT_TEMPLATE_GET_LIST_ERROR                      = 20004
	REPORT_TEMPLATE_GET_LIST_SUCCESS                    = 20005
	REPORT_TEMPLATE_GET_ERROR                           = 20006
	REPORT_TEMPLATE_GET_SUCCESS                         = 20007
	REPORT_TEMPLATE_UPDATE_ERROR                        = 20008
	REPORT_TEMPLATE_UPDATE_SUCCESS                      = 20009
	REPORT_TEMPLATE_DELETE_ERROR                        = 20010
	REPORT_TEMPLATE_DELETE_SUCCESS                      = 20011
	REPORT_TEMPLATE_VAR_MODULE                          = "report_template_var"
	REPORT_TEMPLATE_VAR_CREATE_SUCCESS                  = 30001
	REPORT_TEMPLATE_VAR_CREATE_ERROR                    = 30002
	REPORT_TEMPLATE_VAR_EXIST                           = 30003
	REPORT_TEMPLATE_VAR_GET_LIST_ERROR                  = 30004
	REPORT_TEMPLATE_VAR_GET_LIST_SUCCESS                = 30005
	REPORT_TEMPLATE_VAR_GET_ERROR                       = 30006
	REPORT_TEMPLATE_VAR_GET_SUCCESS                     = 30007
	REPORT_TEMPLATE_VAR_UPDATE_ERROR                    = 30008
	REPORT_TEMPLATE_VAR_UPDATE_SUCCESS                  = 30009
	REPORT_TEMPLATE_VAR_DELETE_ERROR                    = 30010
	REPORT_TEMPLATE_VAR_DELETE_SUCCESS                  = 30011
	REPORT_TEMPLATE_VAR_RENDER_RECORD_GET_LIST_ERROR    = 30012
	REPORT_TEMPLATE_VAR_RENDER_RECORD_GET_LIST_SUCCESS  = 30013
	REPORT_TEMPLATE_VAR_RENDER_PROGRESS_GET_ERROR       = 30014
	REPORT_TEMPLATE_VAR_RENDER_PROGRESS_GET_SUCCESS     = 30015
	REPORT_TEMPLATE_VAR_MERGE_ERROR                     = 30016
	REPORT_TEMPLATE_VAR_MERGE_SUCCESS                   = 30017
	REPORT_TEMPLATE_SLOT_MODULE                         = "report_template_slot"
	REPORT_TEMPLATE_SLOT_CREATE_SUCCESS                 = 40001
	REPORT_TEMPLATE_SLOT_CREATE_ERROR                   = 40002
	REPORT_TEMPLATE_SLOT_EXIST                          = 40003
	REPORT_TEMPLATE_SLOT_GET_LIST_ERROR                 = 40004
	REPORT_TEMPLATE_SLOT_GET_LIST_SUCCESS               = 40005
	REPORT_TEMPLATE_SLOT_GET_ERROR                      = 40006
	REPORT_TEMPLATE_SLOT_GET_SUCCESS                    = 40007
	REPORT_TEMPLATE_SLOT_UPDATE_ERROR                   = 40008
	REPORT_TEMPLATE_SLOT_UPDATE_SUCCESS                 = 40009
	REPORT_TEMPLATE_SLOT_DELETE_ERROR                   = 40010
	REPORT_TEMPLATE_SLOT_DELETE_SUCCESS                 = 40011
	REPORT_TEMPLATE_SLOT_RENDER_RECORD_GET_LIST_ERROR   = 40012
	REPORT_TEMPLATE_SLOT_RENDER_RECORD_GET_LIST_SUCCESS = 40013
	REPORT_TEMPLATE_SLOT_RENDER_PROGRESS_GET_ERROR      = 40014
	REPORT_TEMPLATE_SLOT_RENDER_PROGRESS_GET_SUCCESS    = 40015
	REPORT_TEMPLATE_SLOT_MERGE_ERROR                    = 40016
	REPORT_TEMPLATE_SLOT_MERGE_SUCCESS                  = 40017
	REPORT_AUTHORITY_NOTICE_CHANNEL_MODULE              = "authority_notice_channel"
	REPORT_AUTHORITY_NOTICE_CHANNEL_CREATE_SUCCESS      = 50001
	REPORT_AUTHORITY_NOTICE_CHANNEL_CREATE_ERROR        = 50002
	REPORT_AUTHORITY_NOTICE_CHANNEL_EXIST               = 50003
	REPORT_AUTHORITY_NOTICE_CHANNEL_GET_LIST_ERROR      = 50004
	REPORT_AUTHORITY_NOTICE_CHANNEL_GET_LIST_SUCCESS    = 50005
	REPORT_AUTHORITY_NOTICE_CHANNEL_GET_ERROR           = 50006
	REPORT_AUTHORITY_NOTICE_CHANNEL_GET_SUCCESS         = 50007
	REPORT_AUTHORITY_NOTICE_CHANNEL_UPDATE_ERROR        = 50008
	REPORT_AUTHORITY_NOTICE_CHANNEL_UPDATE_SUCCESS      = 50009
	REPORT_AUTHORITY_NOTICE_CHANNEL_DELETE_ERROR        = 50010
	REPORT_AUTHORITY_NOTICE_CHANNEL_DELETE_SUCCESS      = 50011
	REPORT_AUTHORITY_MESSAGE_MODULE                     = "authority_message"
	REPORT_AUTHORITY_MESSAGE_CREATE_SUCCESS             = 60001
	REPORT_AUTHORITY_MESSAGE_CREATE_ERROR               = 60002
	REPORT_AUTHORITY_MESSAGE_EXIST                      = 60003
	REPORT_AUTHORITY_MESSAGE_GET_LIST_ERROR             = 60004
	REPORT_AUTHORITY_MESSAGE_GET_LIST_SUCCESS           = 60005
	REPORT_AUTHORITY_MESSAGE_GET_ERROR                  = 60006
	REPORT_AUTHORITY_MESSAGE_GET_SUCCESS                = 60007
	REPORT_AUTHORITY_MESSAGE_UPDATE_ERROR               = 60008
	REPORT_AUTHORITY_MESSAGE_UPDATE_SUCCESS             = 60009
	REPORT_AUTHORITY_MESSAGE_DELETE_ERROR               = 60010
	REPORT_AUTHORITY_MESSAGE_DELETE_SUCCESS             = 60011
	REPORT_AUTHORITY_MESSAGE_MERGE_ERROR                = 60012
	REPORT_AUTHORITY_MESSAGE_MERGE_SUCCESS              = 60013
	REPORT_AUTHORITY_MESSAGE_SEND_SUCCESS               = 60014
	REPORT_AUTHORITY_MESSAGE_AUDIT_ERROR                = 60015
	REPORT_AUTHORITY_MESSAGE_AUDIT_SUCCESS              = 60016
	REPORT_MODULE                                       = "report"
	REPORT_CREATE_SUCCESS                               = 70001
	REPORT_CREATE_ERROR                                 = 70002
	REPORT_EXIST                                        = 70003
	REPORT_GET_LIST_ERROR                               = 70004
	REPORT_GET_LIST_SUCCESS                             = 70005
	REPORT_GET_ERROR                                    = 70006
	REPORT_GET_SUCCESS                                  = 70007
	REPORT_UPDATE_ERROR                                 = 70008
	REPORT_UPDATE_SUCCESS                               = 70009
	REPORT_DELETE_ERROR                                 = 70010
	REPORT_DELETE_SUCCESS                               = 70011
	REPORT_MERGE_ERROR                                  = 70012
	REPORT_MERGE_SUCCESS                                = 70013
	REPORT_RENDER_SUCCESS                               = 70014
	REPORT_FINISH_SUCCESS                               = 70015
	MESSAGE_MODULE                                      = "message"
	EVENT_MODULE                                        = "event"
	INSTNAT_CHAT_MESSAGE_MODULE                         = "instant_chat_message"
	INSTNAT_CHAT_MESSAGE_CREATE_SUCCESS                 = 80001
	INSTNAT_CHAT_MESSAGE_CREATE_ERROR                   = 80002
	INSTNAT_CHAT_MESSAGE_EXIST                          = 80003
	INSTNAT_CHAT_MESSAGE_GET_LIST_ERROR                 = 80004
	INSTNAT_CHAT_MESSAGE_GET_LIST_SUCCESS               = 80005
	INSTNAT_CHAT_MESSAGE_GET_ERROR                      = 80006
	INSTNAT_CHAT_MESSAGE_GET_SUCCESS                    = 80007
	INSTNAT_CHAT_MESSAGE_UPDATE_ERROR                   = 80008
	INSTNAT_CHAT_MESSAGE_UPDATE_SUCCESS                 = 80009
	INSTNAT_CHAT_MESSAGE_DELETE_ERROR                   = 80010
	INSTNAT_CHAT_MESSAGE_DELETE_SUCCESS                 = 80011
	INSTNAT_CHAT_MESSAGE_MERGE_ERROR                    = 80012
	INSTNAT_CHAT_MESSAGE_MERGE_SUCCESS                  = 80013
	INSTNAT_CHAT_MESSAGE_RENDER_SUCCESS                 = 80014
	INSTNAT_CHAT_MESSAGE_FINISH_SUCCESS                 = 80015
	PARAMS_ERROR                                        = 40000 // 参数错误
	ALL_SUCCESS                                         = 20000 // 成功
)
