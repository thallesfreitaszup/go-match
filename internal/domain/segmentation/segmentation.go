package segmentation

type Segmentation struct {
	Node        Node   `json:"segmentation"`
	WorkspaceID string `json:"workspaceId"`
	CircleID    string `json:"circleId"`
	Name        string `json:"name"`
}

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
	AND LogicalOperator = "AND"
	OR  LogicalOperator = "OR"
)

type SegmentationType string

const (
	SimpleKV SegmentationType = "SIMPLE_KV"
	Regular  SegmentationType = "REGULAR"
)
