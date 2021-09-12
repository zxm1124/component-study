package api

import (
	"github.com/gin-gonic/gin"
)
type PlatForm int

const (
	// 认证客户端ID
	CLIENT_ID_GITHUB = "a444d99337aaf494a875"
	// 认证客户端Secrets
	CLIENT_SECRETS_GITHUB = "55fb71eb893e8aa26e5e1b91097544b7ef69863d"

	// 用于判断oauth服务提供平台类型
	Type_QQ PlatForm = iota
	Type_Wechat
	Type_Github
	Type_Sina
)

func Oauth(c *gin.Context){
	// 获取获取授权码
	// 根据授权码获取token
	// 根据token获取用户资源
	// 封装用户信息

}

// 获取token
func GetToken() {
	// 获取类型
}
// 判断oauth服务提供平台类型
func TokenUrl(p PlatForm,code string) string{
	switch p {
	case Type_QQ:
		return "qq"
	case Type_Wechat:
		return "wechat"
	case Type_Github:
		return "https://github.com/login/oauth/access_token?client_id="+CLIENT_ID_GITHUB+
			"&client_secret="+CLIENT_SECRETS_GITHUB+"&code="+code
	case Type_Sina:
		return "sina"
	default:
		return ""
	}
}