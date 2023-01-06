package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeInvalidToken
	CodeInvalidAuthFormat
	CodeNotLogin
	CodeQueryDatabase
	CodeRegistrationFailed
	CodeRefshToken
)

var msgFlags = map[ResCode]string{
	CodeSuccess:            "success",
	CodeInvalidParams:      "请求参数错误",
	CodeUserExist:          "用户名重复",
	CodeUserNotExist:       "用户不存在",
	CodeInvalidPassword:    "用户名或密码错误",
	CodeServerBusy:         "服务繁忙",
	CodeQueryDatabase:      "查询数据库错误",
	CodeRegistrationFailed: "注册失败",

	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",
	CodeRefshToken:        "刷新 token 错误",
}

func (c ResCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
