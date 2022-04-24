package tool

import (
	"errors"
	"reflect"
)

func Analyse(obj interface{}) error {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		childrenType := t.Field(i)
		notNull, ok := childrenType.Tag.Lookup("NotNull")
		if !ok {
			continue
		} else {
			switch childrenType.Type.Kind().String() {
			case "int":
			case "int32":
			case "int64":
			case "float32":
			case "float64":
				if v.Field(i).IsZero() {
					return errors.New(notNull)
				}
			case "string":
				if v.Field(i).String() == "" {
					return errors.New(notNull)
				}
			}
		}
	}
	return nil
}
