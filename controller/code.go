package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota //下面会自增
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeInvalidPassword: "用户名或密码错误",
	CodeUserNotExist:    "用户不存在",
	CodeServerBusy:      "服务繁忙",
	CodeUserExist:       "用户已存在",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
