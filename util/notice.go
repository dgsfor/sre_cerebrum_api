package util

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func NoticeToQywxRobot(msgtype string,content string, robotKey string) (bool,error) {
	headerMap := make(map[string]string)
	headerMap["content-type"] = "application/json"
	postData := make(map[string]interface{})
	postData["msgtype"] = msgtype
	ContentData := make(map[string]interface{})
	ContentData["content"] = content
	postData[msgtype] = ContentData
	wxUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", robotKey)
	response, err := HandlerRequest("POST",wxUrl,headerMap,postData)
	fmt.Println(postData)
	jsonBody, errNewJson := simplejson.NewJson(response)
	if err != nil || errNewJson != nil || jsonBody == nil {
		return false, err
	}
	return true,nil
}


// 企业微信应用
func GetQywxAppAccessToken(corpId string,corpSecret string) string {
	headerMap := make(map[string]string)
	headerMap["content-type"] = "application/json"
	wxUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpId,corpSecret)
	response, err := HandlerRequest("GET",wxUrl,headerMap,nil)
	jsonBody, errNewJson := simplejson.NewJson(response)
	if err != nil || errNewJson != nil || jsonBody == nil {
		return ""
	}
	accessToken,_ :=jsonBody.Get("access_token").String()
	return accessToken
}

func NoticeToQywxApp(accessToken string,corpId string,msgtype string,content string,touser string) (bool,error) {
	headerMap := make(map[string]string)
	headerMap["content-type"] = "application/json"
	postData := make(map[string]interface{})
	postData["msgtype"] = msgtype
	ContentData := make(map[string]interface{})
	ContentData["content"] = content
	postData[msgtype] = ContentData
	postData["agentid"] = corpId
	postData["touser"] = touser

	wxUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", accessToken)
	response, err := HandlerRequest("POST",wxUrl,headerMap,postData)
	fmt.Println(postData)
	jsonBody, errNewJson := simplejson.NewJson(response)
	if err != nil || errNewJson != nil || jsonBody == nil {
		return false, err
	}
	return true,nil
}