package errno

var (
	// 100x 通用类型
	OK                  = &Errno{Code: 0, Message: "成功"}
	ParamsError         = &Errno{Code: 1000, Message: "参数错误"}
	AuthError           = &Errno{Code: 1001, Message: "禁止访问"}
	InternalServerError = &Errno{Code: 1002, Message: "服务器错误"}
	DatabaseError       = &Errno{Code: 1003, Message: "数据库错误"}
	TooManyRequestError = &Errno{Code: 1004, Message: "发送了太多请求"}
	SessionError        = &Errno{Code: 1005, Message: "您的 Session 已过期"}
	NotFoundError       = &Errno{Code: 1006, Message: "Not Found"}

	// 500x 第三方错误
	SmsError = &Errno{Code: 5000, Message: "短信发送异常"}
)
