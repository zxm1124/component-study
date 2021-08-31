# 封装响应处理和分页处理
## 响应处理
### 实现思路
#### 1.声明成员函数 包含普通数据响应，列表数据响应，错误数据响应
#### 2.在路由方法中使用
   ```go
   func hello(c *gin.Context){
       resp := app.NewResponse(c)
       // TODO
       resp.ToResponse(data)
   }

   func toErr(c *gin.Context){
          resp := app.NewResponse(c)
          // TODO：鉴权逻辑
          if err != nil{
                // TODO：日志记录
                resp.ToErrorResponse(errcode.UnauthorizedTokenError)
          }
      }
   ```

## 分页处理
### 实现思路
    由于后端获取的数据均为string类型，则需要将获得的参数转换成int类型
#### 1.封装Pager，存储page、pageSize、total等信息
```go
// 封装分页数据实体
type Pager struct {
	// 获取页
	Page int `json:"page"`
	// 每页显示数
	PageSize int `json:"page_size"`
	// 总条数
	TotalRows int `json:"total_rows"`
}
```
#### 2.定义转换参数的成员函数
     -  转换方法内部可使用strconv.AtoI方法
     -  或者自定义covert包，声明一些转换MustInt之类的方法
     ```go
     func (s StrTo) MustInt() int{
          	v, _ := s.Int()
          	return v
          }
          // ...
     ```
