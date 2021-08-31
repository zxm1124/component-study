package errorcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int
	msg     string
	details []string // 错误详细信息
}

// 防止状态码过多，导致错误添加重复的状态码
// 可在编译时报错提醒进行修改
var codes = map[int]string{}

func NewCode(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		// 不可以创建新的code码
		panic(fmt.Sprintf("错误码：%d已经存在，请重新更换", code))
	}
	// 将该错误码存放到map中
	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.code, e.msg)
}
func (e *Error) Code() int {
	return e.code
}
func (e *Error) Msg() string {
	return e.msg
}
func (e *Error) Details() []string {
	return e.details
}

// 链式编程，返回Error
func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	// 遍历要插入的信息
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}

// 获取错误码对应的http状态码
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedTokenError.Code():
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
