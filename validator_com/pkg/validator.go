package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Accept-language
		locale := c.GetHeader("locale")
		uni := ut.New(zh.New(), en.New(), zh_Hant_TW.New())
		// 获取translator
		trans, _ := uni.GetTranslator(locale)
		// 绑定engine
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			//register translate
			// 注册翻译器
			switch locale {
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}

// 封装错误信息
type ValidError struct {
	Key     string
	Message string
}

// 封装多个validError
type ValidErrors []*ValidError

// 重写Error方法，相当于继承error，可以作为error使用
// 底层的validator.ValidationErrors也是error类型，后期需要转换成该类型，所以必须实现该方法
func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// 用于在api层进行封装数据，并检测是否符合验证规则
func BindAndValid(c *gin.Context, param interface{}) (bool, ValidErrors) {
	var errInfo ValidErrors
	// 将c中的参数绑定到v中
	err := c.ShouldBind(param)
	if err != nil {
		// 获取trans
		trans, _ := c.Value("trans").(ut.Translator)
		// 将err转换成validator.ValidationErrors类型
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errInfo
		}
		// 遍历err中的数据，添加到ValidaErrors中
		for k, v := range verrs.Translate(trans) {
			errInfo = append(errInfo, &ValidError{
				Key:     k,
				Message: v,
			})
		}
		return false, errInfo
	}
	return true, nil
}
