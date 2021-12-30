package eval

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
	"strconv"
	"strings"
)

func Expression(exp string, mapParameters map[string]interface{}) (bool, error) {
	functions := map[string]govaluate.ExpressionFunction{
		"len": func(args ...interface{}) (interface{}, error) {
			inputVar := args[0].([]int)
			return len(inputVar), nil
		},
		"contains": func(args ...interface{}) (interface{}, error) {
			reference := convertString(args[0])
			compared := convertString(args[1])
			return strings.Contains(reference, compared), nil
		},
		"equal":   equal,
		"toFloat": toFloat,
	}
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(exp, functions)
	if err != nil {
		return false, err
	}
	result, err := expression.Evaluate(mapParameters)
	if err != nil {
		return false, err
	}
	resultBool, ok := result.(bool)
	if !ok {
		return false, errors.New("error converting expression")
	}
	return resultBool, nil
}

func toFloat(arguments ...interface{}) (interface{}, error) {
	dataType := reflect.TypeOf(arguments[0])
	switch dataType.Kind() {
	case reflect.Int32:
		int32 := arguments[0].(int32)
		return float64(int32), nil
	case reflect.Int64:
		int64 := arguments[0].(int64)
		return float64(int64), nil
	case reflect.Int:
		int := arguments[0].(int)
		return float64(int), nil
	case reflect.String:
		text := arguments[0].(string)
		return strconv.ParseFloat(text, 64)
	case reflect.Float64:
		float64 := arguments[0].(float64)
		return float64, nil
	case reflect.Float32:
		float32 := arguments[0].(float32)
		return float32, nil
	default:
		return nil, fmt.Errorf("type not supported: %s", dataType.Kind())
	}
}

func equal(args ...interface{}) (interface{}, error) {

	reference := convertString(args[0])
	compared := convertString(args[1])
	return strings.EqualFold(reference, compared), nil
}
func convertString(data interface{}) string {
	dataType := reflect.TypeOf(data)
	switch dataType.Kind() {
	case reflect.Bool:
		boolean := data.(bool)
		return strconv.FormatBool(boolean)
	case reflect.Int32:
		int32 := data.(int32)
		return strconv.Itoa(int(int32))
	case reflect.Int64:
		int64 := data.(int64)
		return strconv.Itoa(int(int64))
	case reflect.Float64:
		float64 := data.(float64)
		return strconv.Itoa(int(float64))

	case reflect.String:
		return data.(string)
	default:
		return ""
	}
}
