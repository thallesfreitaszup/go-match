package segmentation

import (
	"fmt"
	"strconv"
)

type Content struct {
	Key       string    `json:"key"`
	Condition Condition `json:"condition"`
	Value     string    `json:"value"`
}

type Condition string

const (
	Equal       Condition = "EQUAL"
	NotEqual    Condition = "NOT_EQUALS"
	Contains    Condition = "CONTAINS"
	LowerThan   Condition = "LOWER_THAN"
	GreaterThan Condition = "GREATER_THAN"
)

func (c Condition) Expression(key, value string) string {
	switch c {
	case Equal:
		if _, err := strconv.ParseFloat(value, 64); err == nil {
			return fmt.Sprintf("toFloat(%s) == toFloat(%s)", key, value)
		}
		return fmt.Sprintf("equal(%s,'%s')", key, value)
	case NotEqual:
		if _, err := strconv.ParseFloat(value, 64); err == nil {
			return fmt.Sprintf("toFloat(%s) != toFloat(%s)", key, value)
		}
		return fmt.Sprintf("!equal(%s,'%s')", key, value)
	case Contains:
		return fmt.Sprintf("contains(%s, %s)", key, value)
	default:
		return ""
	}
}
