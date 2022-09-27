package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"reflect"
	"strings"

	"akcasbin/global"
	"akcasbin/utils"
)

func init() {
	err := utils.RegisterRule("PageVerify", utils.Rules{
		"Page":     {utils.NotEmpty()},
		"PageSize": {utils.NotEmpty()},
	})

	err = utils.RegisterRule("IdVerify",
		utils.Rules{
			"Id": {utils.NotEmpty()},
		},
	)
	if err != nil {
		global.CASBIN_LOG.Error("初始化自定义规则失败：", zap.Error(err))
	}
}

func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator 引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.FORM_TRANS, ok = uni.GetTranslator(locale)
		if !ok {
			global.CASBIN_LOG.Error("初始化 form translator 失败")
			return fmt.Errorf("初始化 form translator: %v 失败", locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.FORM_TRANS)
		case "zh":
			err = zh_translations.RegisterDefaultTranslations(v, global.FORM_TRANS)
			if err != nil {
				return err
			}
		default:
			en_translations.RegisterDefaultTranslations(v, global.FORM_TRANS)
		}
		return
	}
	return
}
