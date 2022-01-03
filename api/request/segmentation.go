package request

import (
	"encoding/json"
	"go-match/internal/domain/segmentation"
)

type SegmentationRequest struct {
	Node        NodeRequest      `json:"node"`
	WorkspaceID string           `json:"workspaceId"`
	CircleID    string           `json:"circleId"`
	Name        string           `json:"name"`
	Type        SegmentationType `json:"type"`
}

type SegmentationType string

const (
	SimpleKV SegmentationType = "SIMPLE_KV"
	Regular  SegmentationType = "REGULAR"
)

type NodeRequest struct {
	Clauses         []NodeRequest   `json:"clauses"`
	Type            NodeType        `json:"type"`
	Content         Content         `json:"content"`
	LogicalOperator LogicalOperator `json:"logicalOperator"`
}

func (r NodeRequest) ToNode() (segmentation.Node, error) {
	nodeResponse := segmentation.Node{}
	nodeBytes, err := json.Marshal(r)
	if err != nil {
		return segmentation.Node{}, err
	}
	err = json.Unmarshal(nodeBytes, &nodeResponse)
	if err != nil {
		return segmentation.Node{}, err
	}
	return nodeResponse, nil
}

type Content struct {
	Key       string    `json:"key"`
	Condition Condition `json:"condition"`
	Value     string    `json:"value"`
}

type NodeType string
type LogicalOperator string

const (
	Rule   NodeType = "RULE"
	Clause NodeType = "CLAUSE"
)

const (
	AND LogicalOperator = "AND"
	OR  LogicalOperator = "OR"
)

type Condition string

const (
	Equals      Condition = "EQUAL"
	NotEquals   Condition = "NOT_EQUAL"
	Contains    Condition = "CONTAINS"
	LowerThan   Condition = "LOWER_THAN"
	GreaterThan Condition = "GREATER_THAN"
)

func (c Condition) Expression() string {
	switch c {
	case Equals:
		return `equal(x,y)`
	case NotEquals:
		return `toStr("x") != toStr("y")`
	case Contains:
		return `contains(x,y)`
	default:
		return ""
	}
}
