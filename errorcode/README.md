# 封装错误码
## 实现思路
1. 封装错误码实体，包括状态码，回显信息，详细信息
2. 创建New方法，用于实现创建错误码实体
3. 创建一个map集合，用于存储所有创建好的错误码实体，可用于检测是否有重复的状态码，以便于前期编译能够检查出来
4. 创建成员函数：
    - 重写Error方法，用于返回错误信息
    - 成员变量的get方法
    - 给错误码实体添加详细信息的方法
5. 创建一个常用的错误码全局变量，封装系统所需要的全部状态码

### 1.封装错误码实体
```go
type Error struct {
	code int
	msg string
	details []string // 错误详细信息
}
```

### 2.创建new方法
```go
func NewCode(code int, msg string) *Error{
	if _, ok := codes[code]; ok{
		// 不可以创建新的code码
		panic(fmt.Sprintf("错误码：%d已经存在，请重新更换",code))
	}
	// 将该错误码存放到map中
	codes[code] = msg
	return &Error{
		code: code,
		msg: msg,
	}
}
```

### 3.创建map集合
```go
// 防止状态码过多，导致错误添加重复的状态码
// 可在编译时报错提醒进行修改
var codes = map[int]string{}
```

### 4.创建相应的成员函数
```go
func(e *Error) Error() string{
	return fmt.Sprintf("错误码：%d，错误信息：%s",e.code, e.msg)
}
func(e *Error) Code() int{
	return e.code
}
func(e *Error) Msg() string{
	return e.msg
}
func(e *Error) Details() []string{
	return e.details
}
// 链式编程，返回Error
func(e *Error) WithDetails(details...string) *Error{
	e.details = []string{}
	// 遍历要插入的信息
	for _, d := range details{
		e.details = append(e.details, d)
	}
	return e
}
```

### 5.创建常用的状态码
```go
// 在common_code.go中创建
var(
	Success = NewCode(200, "请求成功")
	ServerError = NewCode(10000000, "服务器内部错误")
	InvalidParams = NewCode(10000001, "无效参数")

	UnauthorizedTokenError = NewCode(10000005, "鉴权失败，token错误")

	// ...
)
```