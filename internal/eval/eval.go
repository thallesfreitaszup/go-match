package eval

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
	"strconv"
	"strings"
)

func EvalExpression(exp string, mapParameters map[string]interface{}) (bool, error) {
	functions := map[string]govaluate.ExpressionFunction{
		"len": func(args ...interface{}) (interface{}, error) {
			inputVar := args[0].([]int)
			return len(inputVar), nil
		},
		"contains": func(args ...interface{}) (interface{}, error) {
			reference := args[0].(string)
			compared := args[1].(string)
			return strings.Contains(reference, compared), nil
		},
		"equal": equal,
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

func equal(args ...interface{}) (interface{}, error) {

	reference := convertString(args[0])
	compared := convertString(args[1])
	return strings.EqualFold(reference, compared), nil
}
func convertString(data interface{}) string {
	dataType := reflect.TypeOf(data)
	switch dataType.Kind() {
	case reflect.Int:
		int := data.(int)
		return strconv.Itoa(int)
	case reflect.Float64:
		float64 := data.(float64)
		return fmt.Sprintf("%f", float64)
	case reflect.Float32:
		float32 := data.(float32)
		return fmt.Sprintf("%f", float32)
	case reflect.Bool:
		boolean := data.(bool)
		return strconv.FormatBool(boolean)
	case reflect.Int32:
		int32 := data.(int32)
		return strconv.Itoa(int(int32))
	case reflect.Int64:
		int64 := data.(int64)
		return strconv.Itoa(int(int64))
	case reflect.String:
		return data.(string)
	default:
		return ""
	}
}
