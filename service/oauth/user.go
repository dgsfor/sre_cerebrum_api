package oauth

import (
	"net/http"
	"ssopa/conf"
	"ssopa/middleware"
	"ssopa/model/auth"
	"ssopa/serializer"
	"ssopa/util"
)

type LoginParams struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"` // 用户名
	Password string `form:"password" json:"password" binding:"required"`   // 密码
}

type RegisterParams struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"` // 用户名
	Password string `form:"password" json:"password" binding:"required"`   // 密码
	Email    string `form:"email" json:"email" binding:"required"`         // 邮箱
}

func (p *LoginParams) Login() serializer.SsopaResponse {
	var SsoPaUserModel auth.SsoPaUsers
	err := conf.Orm.Where("user_name = ?", p.UserName).Find(&SsoPaUserModel).RowsAffected
	if err == 0 {
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err,
				Msg:  "登录失败，用户未找到，请注册！",
			},
			ResCode: serializer.USER_NOT_EXISTS,
		}
	}
	validatePassStatus := util.ValidatePasswords(SsoPaUserModel.Password, []byte(p.Password))
	if !validatePassStatus {
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: nil,
				Msg:  "登录失败，请检查用户密码！",
			},
			ResCode: serializer.PASSWORD_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.USER_MODULE, "login", p.UserName, "nil", "login success",nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: SsoPaUserModel.Email,
			Msg:  "登录成功",
		},
		ResCode: serializer.LOGIN_SUCCESS,
	}
}

func (p *RegisterParams) Register() serializer.SsopaResponse {
	var SsoPaUserModel auth.SsoPaUsers
	err := conf.Orm.Where("user_name = ?", p.UserName).Find(&SsoPaUserModel).RowsAffected
	if err >= 1 {
		middleware.CustomOutPutLog(serializer.USER_MODULE, "register", p.UserName, p.Email, "user exist",nil)
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: err,
				Msg:  "用户已经存在，请直接登录！",
			},
			ResCode: serializer.USER_EXISTS,
		}
	}
	SsoPaUserModel.UserName = p.UserName
	SsoPaUserModel.Password = p.Password
	SsoPaUserModel.Email = p.Email
	createErr := conf.Orm.Create(&SsoPaUserModel).Error
	if createErr != nil {
		middleware.CustomOutPutLog(serializer.USER_MODULE, "register", p.UserName, p.Email, "register error",createErr.Error())
		return serializer.SsopaResponse{
			Response: serializer.Response{
				Code: http.StatusInternalServerError,
				Data: createErr.Error(),
				Msg:  "创建用户失败！",
			},
			ResCode: serializer.CREATE_USER_ERROR,
		}
	}
	middleware.CustomOutPutLog(serializer.USER_MODULE, "register", p.UserName, p.Email, "register success",nil)
	return serializer.SsopaResponse{
		Response: serializer.Response{
			Code: http.StatusOK,
			Data: nil,
			Msg:  "注册成功，请登录!",
		},
		ResCode: serializer.CREATE_USER_SUCCESS,
	}
}
