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
	USER_MODULE                      = "user"
	CREATE_USER_SUCCESS              = 10000 // 创建用户成功
	LOGIN_SUCCESS                    = 10001 // 登录成功
	USER_EXISTS                      = 10002 // 用户存在
	CREATE_USER_ERROR                = 10003 // 创建用户失败
	USER_NOT_EXISTS                  = 10004 // 用户不存在
	PASSWORD_ERROR                   = 10005 // 密码错误
	SETCOOKIE_FAILURE                = 10006 // 设置cookie失败
	LOGOUT_SUCCESS                   = 10007 // 退出登录成功
	REPORT_TEMPLATE_MODULE           = "report_template"
	REPORT_TEMPLATE_CREATE_SUCCESS   = 20001
	REPORT_TEMPLATE_CREATE_ERROR     = 20002
	REPORT_TEMPLATE_EXIST            = 20003
	REPORT_TEMPLATE_GET_LIST_ERROR   = 20004
	REPORT_TEMPLATE_GET_LIST_SUCCESS = 20005
	REPORT_TEMPLATE_GET_ERROR        = 20006
	REPORT_TEMPLATE_GET_SUCCESS      = 20007
	REPORT_TEMPLATE_UPDATE_ERROR     = 20008
	REPORT_TEMPLATE_UPDATE_SUCCESS   = 20009
	REPORT_TEMPLATE_DELETE_ERROR     = 20010
	REPORT_TEMPLATE_DELETE_SUCCESS   = 20011
	REPORT_MODULE                    = "report"
	MESSAGE_MODULE                   = "message"
	EVENT_MODULE                     = "event"

	PARAMS_ERROR = 40000 // 参数错误
	ALL_SUCCESS  = 20000 // 成功
)
