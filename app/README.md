# 封装响应处理和分页处理
## 响应处理
### 实现思路
#### 1.包含普通数据响应，列表数据响应，错误数据响应
#### 2.在路由方法中使用
   ```go
   func hello(c *gin.Context){
       resp := app.NewResponse(c)
       // TODO
       resp.ToResponse(data)
   }
   ```

## 分页处理
### 实现思路
    由于后端获取的数据均为string类型，则需要将获得的参数转换成int类型
#### 1.封装Pager，存储page、pageSize、total等信息
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
