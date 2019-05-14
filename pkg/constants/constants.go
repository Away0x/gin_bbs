package constants

var (
	// UploadImageMimetypes 项目支持上传的文件格式
	UploadImageMimetypes = []string{"jpg", "jpeg", "bmp", "png", "gif"}
)

const (
	// UserNameRegex 验证用户名的正则
	UserNameRegex = `^[A-Za-z0-9\-\_]+$`
	// DateLayout -
	DateLayout = "2006-01-02"
	// DateTimeLayout -
	DateTimeLayout = "2006-01-02 15:04:05"
	// HeaderRequestedWith : Async Request Header
	HeaderRequestedWith = "X-Requested-With"
)
