package utils

import (
	"reflect"
	"strconv"
)

func InferType(value interface{}) (interface{}, string) {
	switch v := value.(type) {
	case int:
		return v, reflect.TypeOf(v).Name()
	case bool:
		return v, reflect.TypeOf(v).Name()
	case float64:
		return v, reflect.TypeOf(v).Name()
	case string:

		if intVal, err := strconv.Atoi(v); err == nil {
			return intVal, reflect.TypeOf(int(0)).Name()
		}

		if floatVal, err := strconv.ParseFloat(v, 64); err == nil {
			return floatVal, reflect.TypeOf(float64(0)).Name()
		}

		if v == "true" {
			return true, reflect.TypeOf(true).Name()
		} else if v == "false" {
			return false, reflect.TypeOf(false).Name()
		}

		return v, reflect.TypeOf(v).Name()
	default:
		return nil, ""
	}
}
