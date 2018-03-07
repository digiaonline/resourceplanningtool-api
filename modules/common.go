package modules

import (
	"reflect"
	"errors"
)

func AggregateStructs(dest interface{}, src interface{}) error {
	val_of_dest := reflect.ValueOf(dest)
	val_of_src := reflect.ValueOf(src)
	if (val_of_dest.Kind() != reflect.Ptr || val_of_src.Kind() != reflect.Ptr) {
		return errors.New("'common.go': both of the input values must be pointers")
	}

	old_values := reflect.ValueOf(src).Elem()
	new_values := reflect.ValueOf(dest).Elem()

	for i := 0; i < new_values.NumField(); i++ {
		field := new_values.Field(i)
		if field.Interface() == "" || field.Interface() == nil || field.Interface() == 0 {
			if field.CanSet() {
				if field.Type().Kind() == reflect.String {
					new_val, _ := old_values.Field(i).Interface().(string)
					field.SetString(new_val)
				} else if field.Type().Kind() == reflect.Int {
					new_val := old_values.Field(i).Interface().(int)
					field.SetInt(int64(new_val))
				}
			}
		}
	}

	return nil
}
