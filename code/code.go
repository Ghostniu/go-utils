package code

type MyCode int64

const (
	Success           MyCode = 0
	InvalidParams     MyCode = 1001
	UserExist         MyCode = 1002
	UserNotExist      MyCode = 1003
	InvalidPassword   MyCode = 1004
	ServerBusy        MyCode = 1005
	InvalidUrl        MyCode = 1006
	InvalidToken      MyCode = 1007
	InvalidAuthFormat MyCode = 1008
	NotLogin          MyCode = 1009
	UnKnow            MyCode = 1010
)

var msgFlags = map[MyCode]string{
	Success:           "success",
	InvalidParams:     "请求参数错误",
	UserExist:         "用户名重复",
	UserNotExist:      "用户不存在",
	InvalidPassword:   "用户名或密码错误",
	ServerBusy:        "服务繁忙",
	InvalidUrl:        "请求路径错误",
	InvalidToken:      "无效的Token",
	InvalidAuthFormat: "认证格式有误",
	NotLogin:          "未登录",
	UnKnow:            "未知错误",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[ServerBusy]
}
