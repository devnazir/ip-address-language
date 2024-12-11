package utils

import (
	"reflect"
	"strconv"
	"strings"
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
		stringValue := strings.TrimSpace(v)

		if intVal, err := strconv.Atoi(stringValue); err == nil {
			return intVal, reflect.TypeOf(int(0)).Name()
		}

		if floatVal, err := strconv.ParseFloat(stringValue, 64); err == nil {
			return floatVal, reflect.TypeOf(float64(0)).Name()
		}

		if stringValue == "true" {
			return true, reflect.TypeOf(true).Name()
		} else if stringValue == "false" {
			return false, reflect.TypeOf(false).Name()
		}

		return stringValue, reflect.TypeOf(v).Name()
	case []interface{}:
		return v, reflect.Array.String()
	case map[string]interface{}:
		return v, reflect.Map.String()

	default:
		return nil, ""
	}
}

func InferDefaultValue(value interface{}) interface{} {
	switch value {
	case "int":
		return *new(int)
	case "bool":
		return *new(bool)
	case "float64":
		return *new(float64)
	case "string":
		return *new(string)
	default:
		return nil
	}
}
