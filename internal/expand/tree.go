package expand

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/relationtuple"
)

type (
	NodeType int
	Tree     struct {
		Type     NodeType              `json:"type"`
		Subject  relationtuple.Subject `json:"subject"`
		Children []*Tree               `json:"children,omitempty"`
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
	case `"union"`:
		*t = Union
	case `"exclusion"`:
		*t = Exclusion
	case `"intersection"`:
		*t = Intersection
	case `"leaf"`:
		*t = Leaf
	default:
		return ErrUnknownNodeType
	}
	return nil
}

func (t *Tree) UnmarshalJSON(v []byte) error {
	type node struct {
		Type     NodeType `json:"type"`
		Children []*Tree  `json:"children,omitempty"`
		Subject  string   `json:"subject"`
	}

	n := &node{}
	if err := json.Unmarshal(v, n); err != nil {
		return errors.WithStack(err)
	}

	var err error
	t.Subject, err = relationtuple.SubjectFromString(n.Subject)
	if err != nil {
		return err
	}

	t.Type = n.Type
	t.Children = n.Children

	return nil
}
