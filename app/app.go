package app

import (
	"component-study1/errorcode"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 封装响应实体
type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response{
	return &Response{
		Ctx: ctx,
	}
}
// 正常响应数据
func (r *Response)ToResponse(data interface{}){
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}
// 响应列表信息
func (r *Response) ToResponseList(list interface{}, totalRows int){
	r.Ctx.JSON(http.StatusOK,gin.H{
		"list": list,
		"pager": Pager{
			Page: GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}
// 响应错误信息
func (r *Response) ToErrorResponse(err errorcode.Error){
	resp := gin.H{
		"code": err.Code(),
		"msg": err.Msg(),
	}
	if len(err.Details()) > 0 {
		resp["details"] = err.Details()
	}
	r.Ctx.JSON(err.StatusCode(), resp)
}

