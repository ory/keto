package expand

import (
	"encoding/json"
	"fmt"
	"strings"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

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

func (t NodeType) ToProto() acl.NodeType {
	switch t {
	case Leaf:
		return acl.NodeType_NODE_TYPE_LEAF
	case Union:
		return acl.NodeType_NODE_TYPE_UNION
	case Exclusion:
		return acl.NodeType_NODE_TYPE_EXCLUSION
	case Intersection:
		return acl.NodeType_NODE_TYPE_INTERSECTION
	}
	return acl.NodeType_NODE_TYPE_UNSPECIFIED
}

func NodeTypeFromProto(t acl.NodeType) NodeType {
	switch t {
	case acl.NodeType_NODE_TYPE_LEAF:
		return Leaf
	case acl.NodeType_NODE_TYPE_UNION:
		return Union
	case acl.NodeType_NODE_TYPE_EXCLUSION:
		return Exclusion
	case acl.NodeType_NODE_TYPE_INTERSECTION:
		return Intersection
	}
	return Leaf
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

func (t *Tree) ToProto() *acl.SubjectTree {
	if t.Type == Leaf {
		return &acl.SubjectTree{
			NodeType: acl.NodeType_NODE_TYPE_LEAF,
			Subject:  t.Subject.ToProto(),
		}
	}

	children := make([]*acl.SubjectTree, len(t.Children))
	for i, c := range t.Children {
		children[i] = c.ToProto()
	}

	return &acl.SubjectTree{
		NodeType: t.Type.ToProto(),
		Subject:  t.Subject.ToProto(),
		Children: children,
	}
}

func TreeFromProto(t *acl.SubjectTree) *Tree {
	if t.NodeType == acl.NodeType_NODE_TYPE_LEAF {
		return &Tree{
			Type:    Leaf,
			Subject: relationtuple.SubjectFromProto(t.Subject),
		}
	}

	children := make([]*Tree, len(t.Children))
	for i, c := range t.Children {
		children[i] = TreeFromProto(c)
	}

	return &Tree{
		Type:     NodeTypeFromProto(t.NodeType),
		Subject:  relationtuple.SubjectFromProto(t.Subject),
		Children: children,
	}
}

func (t *Tree) String() string {
	sub := t.Subject.String()

	if t.Type == Leaf {
		return fmt.Sprintf("☘ %s️", sub)
	}

	children := make([]string, len(t.Children))
	for i, c := range t.Children {
		children[i] = strings.Join(strings.Split(c.String(), "\n"), "\n│  ")
	}

	return fmt.Sprintf("∪ %s\n├─ %s", sub, strings.Join(children, "\n├─ "))
}
