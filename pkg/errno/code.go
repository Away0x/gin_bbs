package errno

var (
	// 100x 通用类型
	OK                  = &Errno{Code: 0, Message: "成功"}
	ParamsError         = &Errno{Code: 1000, Message: "参数错误"}
	AuthError           = &Errno{Code: 1001, Message: "禁止访问"}
	InternalServerError = &Errno{Code: 1002, Message: "服务器错误"}
	ErrDatabase         = &Errno{Code: 1003, Message: "数据库错误"}
)
