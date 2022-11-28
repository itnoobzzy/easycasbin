package utils

import (
	"reflect"
)

func SetDefault(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() == reflect.Ptr {
		// 如果传入的结构体是指针，需要获取值
		t = t.Elem()
		v = v.Elem()
	} else {
		panic("struct must be ptr to struct")
	}

	for i := 0; i < t.NumField(); i++ {
		rtype := t.Field(i)
		rvalue := v.Field(i)

		defaultValue := rtype.Tag.Get("default")
		if isBlank(rvalue) {
			rvalue.Set(reflect.ValueOf(defaultValue))
		}
	}
}
