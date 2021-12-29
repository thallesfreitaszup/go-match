package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type IdentifyRequest struct {
	RequestData map[string]interface{} `json:"requestData"`
}

func main() {
	app := newApp()
	app.start()
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

//
//func getValueFromContent(content Content, s string) (bool, error) {
//	functions := map[string]govaluate.ExpressionFunction{
//		"len": func(args ...interface{}) (interface{}, error) {
//			inputVar := args[0].([]int)
//			return len(inputVar), nil
//		},
//		"contains": func(args ...interface{}) (interface{}, error) {
//			reference := args[0].(string)
//			compared := args[1].(string)
//			return strings.Contains(reference, compared), nil
//		},
//		"equal": equal,
//	}
//	expr := content.Condition.Expression()
//	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, functions)
//	if err != nil {
//		return false, err
//	}
//	parameters := make(map[string]interface{}, 8)
//	parameters["x"] = content.Value
//	parameters["y"] = s
//	result, err := expression.Evaluate(parameters)
//	fmt.Println(result)
//	bool, ok := result.(bool)
//	if !ok {
//		return false, errors.New("error converting expression")
//	}
//	return bool, nil
//}
//
//func identify(key string, value interface{}) ([]SegmentationDB, error) {
//	var segmentation SegmentationDB
//	segmentationArray := make([]SegmentationDB, 0)
//	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
//	if err != nil {
//		return nil, err
//	}
//	collection := client.Database("matcher").Collection("node")
//	filter := bson.D{
//		{"key", key},
//		{"value", value},
//	}
//	find, err := collection.Find(context.TODO(), filter)
//	if err == mongo.ErrNoDocuments {
//		// Do something when no record was found
//		fmt.Println("record does not exist")
//	} else if err != nil {
//		return nil, err
//	}
//	for find.Next(context.TODO()) {
//		err := find.Decode(&segmentation)
//		if err != nil {
//			return nil, err
//		}
//		segmentationArray = append(segmentationArray, segmentation)
//	}
//	return segmentationArray, nil
//}

//func createSegmentation(nodes []Node, client *mongo.Client, request *SegmentationRequest) error {
//	collection := client.Database("matcher").Collection("node")
//	for _, node := range nodes {
//		_, err := collection.InsertOne(context.TODO(), bson.D{
//			{"key", node.Content.Key},
//			{"value", node.Content.Value},
//			{"circleId", request.CircleID},
//		})
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//func decomposeSegmentation(node Node, segmentations *[]Node) {
//	if node.Type == Clause && node.LogicalOperator == OR {
//		for _, clause := range node.Clauses {
//			decomposeSegmentation(clause, segmentations)
//		}
//	} else {
//		*segmentations = append(*segmentations, node)
//	}
//}
func equal(args ...interface{}) (interface{}, error) {

	reference := convertString(args[0])
	compared := convertString(args[1])
	return strings.EqualFold(reference, compared), nil
}
