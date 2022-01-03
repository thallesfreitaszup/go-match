package segmentation

import (
	"fmt"
	"strconv"
)

const (
	SpaceString = " "
)

type Segmentation struct {
	Node        Node   `json:"segmentation"`
	WorkspaceID string `json:"workspaceId"`
	CircleID    string `json:"circleId"`
	Name        string `json:"name"`
}

type Content struct {
	Key       string    `json:"key"`
	Condition Condition `json:"condition"`
	Value     string    `json:"value"`
}

type NodeType string
type LogicalOperator string

func (o LogicalOperator) Expression() string {
	switch o {
	case AND:
		return "&&"
	default:
		return "||"
	}
}

const (
	Rule   NodeType = "RULE"
	Clause NodeType = "CLAUSE"
)

const (
	AND LogicalOperator = "AND"
	OR  LogicalOperator = "OR"
)

type Node struct {
	Clauses         []Node          `json:"clauses"`
	Type            NodeType        `json:"type"`
	Content         Content         `json:"content"`
	LogicalOperator LogicalOperator `json:"logicalOperator"`
}

func (n Node) Expression() string {
	expression := ""
	if n.Type == Clause {
		expression += "("
		for _, clause := range n.Clauses {
			expression += clause.Expression()
			expression = expression + n.LogicalOperator.Expression() + SpaceString
		}

		expression = expression[:len(expression)-4]
		expression += ")" + SpaceString
		return expression
	} else {
		return fmt.Sprintf("%s%s", n.Content.Condition.Expression(n.Content.Key, n.Content.Value), SpaceString)
	}
}

type SegmentationType string

const (
	SimpleKV SegmentationType = "SIMPLE_KV"
	Regular  SegmentationType = "REGULAR"
)

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
