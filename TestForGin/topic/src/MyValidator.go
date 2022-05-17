package src

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// func TopicUrl(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value,
// 	filedType reflect.Type, fieldKind reflect.Kind, param string) bool {
// 	fmt.Println(topStruct)
// 	fmt.Println(topStruct.Interface())
// 	return false
// }

var TopicUrlNew validator.Func = func(fl validator.FieldLevel) bool {
	//正则校验
	url, ok := fl.Field().Interface().(string)
	if ok {
		if matched, _ := regexp.MatchString(`^\w{4,10}$`, url); matched {
			return true
		}
	}

	fmt.Printf("topicurl 校验不通过%s :", url)
	return false
}
