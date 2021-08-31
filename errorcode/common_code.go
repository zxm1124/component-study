package errorcode

var (
	Success       = NewCode(200, "请求成功")
	ServerError   = NewCode(10000000, "服务器内部错误")
	InvalidParams = NewCode(10000001, "无效参数")

	UnauthorizedTokenError = NewCode(10000005, "鉴权失败，token错误")

	// ...
)
