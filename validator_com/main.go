package main

import (
	"component-study1/validator_com/model"
	"component-study1/validator_com/pkg"
	"github.com/gin-gonic/gin"

	"log"
)

func main() {
	r := gin.New()
	// 启用国际化中间件
	r.Use(pkg.Translations())

	r.GET("/login", func(c *gin.Context) {
		// 绑定数据
		param := model.LoginRequest{}
		valid, errs := pkg.BindAndValid(c, &param)
		if !valid {
			// TODO：验证未通过，回显错误信息
			log.Fatalf("errs: %v", errs.Errors())
		}
		// TODO：验证通过，业务处理
	})

	r.Run(":8181")

}
