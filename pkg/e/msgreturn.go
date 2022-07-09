package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token超时",
	ErrorAuth:                  "Token生成失败",
	ErrorNotCompare:            "Token错误",
	ErrorDatabase:              "数据库操作出错，请重试",
}

//GetMsg 获取参数化的msg值
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
