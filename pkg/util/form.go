/*
 * @Description:参数验证
 * @Autor: 小明～
 * @Date: 2021-09-15 17:49:56
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-15 10:23:55
 */
package util

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	val "github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

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

// func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
// 	var errs ValidErrors
// 	err := c.ShouldBind(v)
// 	if err != nil {
// 		v := c.Value("trans")
// 		trans, _ := v.(ut.Translator)
// 		verrs, ok := err.(val.ValidationErrors)
// 		if !ok {
// 			return false, errs
// 		}

// 		for key, value := range verrs.Translate(trans) {
// 			errs = append(errs, &ValidError{
// 				Key:     key,
// 				Message: value,
// 			})
// 		}

// 		return false, errs
// 	}

// 	return true, nil
// }
func GetTrans(c *gin.Context) ut.Translator {
	uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
	locale := c.GetHeader("locale")
	trans, _ := uni.GetTranslator(locale)
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}

	// if ok {
	// 	switch locale {
	// 	case "zh":
	// 		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	// 		break
	// 	case "en":
	// 		_ = en_translations.RegisterDefaultTranslations(v, trans)
	// 		break
	// 	default:
	// 		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	// 	}
	// }
	return trans
}

func BindAndValid(c *gin.Context, v interface{}) ValidErrors {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		trans := GetTrans(c)
		list, ok := err.(val.ValidationErrors)
		if !ok {
			return nil
		}
		for key, value := range list.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return errs
	}
	return nil
}
