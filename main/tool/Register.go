package tool

import (
	"github.com/gogf/gf/net/ghttp"
	"reflect"
)

func BindObjectReflect(pattern string, object interface{}, s *ghttp.Server) {
	v := reflect.ValueOf(object).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		if !value.CanSet() {
			continue
		}
		tag := t.Field(i).Tag
		method := tag.Get("method")
		address := tag.Get("address")
		s.BindHandler(method+":"+pattern+address, value.Interface().(func(request *ghttp.Request)))
	}
}
