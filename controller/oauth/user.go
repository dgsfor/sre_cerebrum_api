package oauth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ssopa/serializer"
	oauth_service "ssopa/service/oauth"
	"ssopa/util"
)

// Login login
func Login(c *gin.Context) {
	svc := oauth_service.LoginParams{}
	if err := c.ShouldBind(&svc); err != nil {
		c.JSON(http.StatusBadRequest, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusBadRequest,
				Data: err.Error(),
				Msg:  "参数错误，请检查！",
			},
			ResCode: serializer.PARAMS_ERROR,
		})
		return
	}
	result := svc.Login()
	if result.ResCode != serializer.LOGIN_SUCCESS {
		c.JSON(result.Code, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: result.Code,
				Data: result.Data,
				Msg:  result.Msg,
			},
			ResCode: result.ResCode,
		})
		return
	}
	err := util.SetLoginCookies(c, svc.UserName, result.Data.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "set cookie failure",
			},
			ResCode: serializer.SETCOOKIE_FAILURE,
		})
		return
	}
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
	//c.Redirect(http.StatusFound,conf.GetConfig("SreStabilityOperatingPlatform::FrontendURL").String())
}

// Register register
func Register(c *gin.Context) {
	svc := oauth_service.RegisterParams{}
	if err := c.ShouldBind(&svc); err != nil {
		c.JSON(http.StatusBadRequest, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusBadRequest,
				Data: err.Error(),
				Msg:  "参数错误，请检查！",
			},
			ResCode: serializer.PARAMS_ERROR,
		})
		return
	}
	result := svc.Register()
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

// Logout logout
func Logout(c *gin.Context) {
	util.Logout(c)
	c.JSON(http.StatusOK, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "logout success",
		},
		ResCode: serializer.LOGOUT_SUCCESS,
	})
}

// GetUserInfo get user info
func GetUserInfo(c *gin.Context) {
	u, err := util.GetUserCookie(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err.Error(),
				Msg:  "get user cookie fail",
			},
			ResCode: serializer.SETCOOKIE_FAILURE,
		})
	}
	c.JSON(http.StatusOK, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: u,
			Msg:  "get user info success",
		},
		ResCode: serializer.ALL_SUCCESS,
	})
}

// GetUserList 获取所有用户列表
func GetUserList(c *gin.Context) {
	users, _ := util.GetUserCookie(c)
	result := oauth_service.GetUserList(users)
	c.JSON(result.Code, serializer.SsopaResponse{
		Response: serializer.Response{
			Code: result.Code,
			Data: result.Data,
			Msg:  result.Msg,
		},
		ResCode: result.ResCode,
	})
}

func CheckLogin(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}
