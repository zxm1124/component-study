package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// 封装分页数据实体
type Pager struct {
	// 获取页
	Page int `json:"page"`
	// 每页显示数
	PageSize int `json:"page_size"`
	// 总条数
	TotalRows int `json:"total_rows"`
}
// 将获取到的param转换成int类型
func GetPage(c *gin.Context) int{
	page,err := strconv.Atoi(c.Query("page"))
	if err !=nil{
		// todo
	}
	return page
}

func GetPageSize(c *gin.Context) int{
	// TODO 转换成int数据
	return 0
}

func GetPageOffset(page ,pageSize int) int{
	offset := 0
	if page > 0 {
		return (page - 1) * pageSize
	}
	return offset 
}
