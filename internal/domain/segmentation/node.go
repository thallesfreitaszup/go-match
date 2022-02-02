package segmentation

import "fmt"

const (
	SpaceString = " "
)

type Node struct {
	Clauses         []Node          `json:"clauses"`
	Type            NodeType        `json:"type"`
	Content         Content         `json:"content"`
	LogicalOperator LogicalOperator `json:"logicalOperator"`
}

type NodeType string

const (
	Rule   NodeType = "RULE"
	Clause NodeType = "CLAUSE"
)

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
