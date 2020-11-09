package expand

import (
	"errors"

	"github.com/ory/keto/models"
)

type (
	NodeType int
	Tree     struct {
		Type     NodeType       `json:"type"`
		Subject  models.Subject `json:"subject"`
		Children []*Tree        `json:"children"`
	}
)

const (
	Union NodeType = iota
	Exclusion
	Intersection
	Leaf
)

var (
	ErrUnknownNodeType = errors.New("unknown node type")
)

func (t NodeType) String() string {
	switch t {
	case Union:
		return "union"
	case Exclusion:
		return "exclusion"
	case Intersection:
		return "intersection"
	case Leaf:
		return "leaf"
	}
	return ""
}

func (t NodeType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t *NodeType) UnmarshalJSON(v []byte) error {
	switch string(v) {
	case "union":
		*t = Union
	case "exclusion":
		*t = Exclusion
	case "intersection":
		*t = Intersection
	case "leaf":
		*t = Leaf
	default:
		return ErrUnknownNodeType
	}
	return nil
}
