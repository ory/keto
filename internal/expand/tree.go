package expand

import (
	"encoding/json"
	"fmt"
	"strings"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/relationtuple"
)

// swagger:enum NodeType
type NodeType string

const (
	Union        NodeType = "union"
	Exclusion    NodeType = "exclusion"
	Intersection NodeType = "intersection"
	Leaf         NodeType = "leaf"
)

// swagger:model expandTree
type Tree struct {
	// required: true
	Type NodeType `json:"type"`
	// required: true
	Subject  relationtuple.Subject `json:"subject"`
	Children []*Tree               `json:"children,omitempty"`
}

var (
	ErrUnknownNodeType = errors.New("unknown node type")
)

func (t NodeType) String() string {
	return string(t)
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
		Type     NodeType        `json:"type"`
		Children []*Tree         `json:"children,omitempty"`
		Subject  json.RawMessage `json:"subject"`
	}

	n := &node{}
	if err := json.Unmarshal(v, n); err != nil {
		return errors.WithStack(err)
	}

	var err error
	t.Subject, err = relationtuple.SubjectFromJSON(n.Subject)
	if err != nil {
		return err
	}

	t.Type = n.Type
	t.Children = n.Children

	return nil
}

// swagger:ignore
func (t *Tree) ToProto() *acl.SubjectTree {
	if t == nil {
		return nil
	}

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

// swagger:ignore
func TreeFromProto(t *acl.SubjectTree) (*Tree, error) {
	if t == nil {
		return nil, nil
	}

	sub, err := relationtuple.SubjectFromProto(t.Subject)
	if err != nil {
		return nil, err
	}
	self := &Tree{
		Type:    NodeTypeFromProto(t.NodeType),
		Subject: sub,
	}

	if t.NodeType != acl.NodeType_NODE_TYPE_LEAF {
		self.Children = make([]*Tree, len(t.Children))
		for i, c := range t.Children {
			var err error
			self.Children[i], err = TreeFromProto(c)
			if err != nil {
				return nil, err
			}
		}
	}

	return self, nil
}

func (t *Tree) String() string {
	if t == nil {
		return ""
	}

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
