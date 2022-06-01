package expand

import (
	"encoding/json"
	"fmt"
	"strings"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/relationtuple"
)

// swagger:enum NodeType
type NodeType string

const (
	Union           NodeType = "union"
	Exclusion       NodeType = "exclusion"
	Intersection    NodeType = "intersection"
	TupeToUserset   NodeType = "tuple_to_userset"
	ComputedUserset NodeType = "computed_userset"
	Leaf            NodeType = "leaf"
)

// swagger:ignore
type Tree struct {
	Type     NodeType
	Subject  relationtuple.Subject
	Tuple    *relationtuple.InternalRelationTuple
	Children []*Tree
}

func (t *Tree) GetSubject() relationtuple.Subject {
	if t.Subject != nil {
		return t.Subject
	}
	return t.Tuple.Subject
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

func (t NodeType) ToProto() rts.NodeType {
	switch t {
	case Leaf:
		return rts.NodeType_NODE_TYPE_LEAF
	case Union:
		return rts.NodeType_NODE_TYPE_UNION
	case Exclusion:
		return rts.NodeType_NODE_TYPE_EXCLUSION
	case Intersection:
		return rts.NodeType_NODE_TYPE_INTERSECTION
	}
	return rts.NodeType_NODE_TYPE_UNSPECIFIED
}

func NodeTypeFromProto(t rts.NodeType) NodeType {
	switch t {
	case rts.NodeType_NODE_TYPE_LEAF:
		return Leaf
	case rts.NodeType_NODE_TYPE_UNION:
		return Union
	case rts.NodeType_NODE_TYPE_EXCLUSION:
		return Exclusion
	case rts.NodeType_NODE_TYPE_INTERSECTION:
		return Intersection
	}
	return Leaf
}

// swagger:model expandTree
type node struct {
	// required: true
	Type       NodeType                  `json:"type"`
	Children   []*node                   `json:"children,omitempty"`
	SubjectID  *string                   `json:"subject_id,omitempty"`
	SubjectSet *relationtuple.SubjectSet `json:"subject_set,omitempty"`
}

func (n *node) toTree() (*Tree, error) {
	t := &Tree{}
	if n.SubjectID == nil && n.SubjectSet == nil {
		return nil, errors.WithStack(relationtuple.ErrNilSubject)
	} else if n.SubjectID != nil && n.SubjectSet != nil {
		return nil, errors.WithStack(relationtuple.ErrDuplicateSubject)
	}

	if n.SubjectID != nil {
		t.Subject = &relationtuple.SubjectID{ID: *n.SubjectID}
	} else {
		t.Subject = n.SubjectSet
	}

	t.Type = n.Type

	if n.Children != nil {
		t.Children = make([]*Tree, len(n.Children))
		for i := range n.Children {
			var err error
			t.Children[i], err = n.Children[i].toTree()
			if err != nil {
				return nil, err
			}
		}
	}

	return t, nil
}

func (n *node) fromTree(t *Tree) error {
	n.Type = t.Type
	n.SubjectID = t.GetSubject().SubjectID()
	n.SubjectSet = t.GetSubject().SubjectSet()

	if t.Children != nil {
		n.Children = make([]*node, len(t.Children))
		for i := range t.Children {
			n.Children[i] = &node{}
			if err := n.Children[i].fromTree(t.Children[i]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Tree) UnmarshalJSON(v []byte) error {
	var n node
	if err := json.Unmarshal(v, &n); err != nil {
		return errors.WithStack(err)
	}

	tt, err := (&n).toTree()
	if err != nil {
		return err
	}

	*t = *tt
	return nil
}

func (t *Tree) MarshalJSON() ([]byte, error) {
	var n node
	if err := n.fromTree(t); err != nil {
		return nil, err
	}
	return json.Marshal(n)
}

// swagger:ignore
func (t *Tree) ToProto() *rts.SubjectTree {
	if t == nil {
		return nil
	}

	if t.Type == Leaf {
		return &rts.SubjectTree{
			NodeType: rts.NodeType_NODE_TYPE_LEAF,
			Subject:  t.GetSubject().ToProto(),
		}
	}

	children := make([]*rts.SubjectTree, len(t.Children))
	for i, c := range t.Children {
		children[i] = c.ToProto()
	}

	return &rts.SubjectTree{
		NodeType: t.Type.ToProto(),
		Subject:  t.GetSubject().ToProto(),
		Children: children,
	}
}

// swagger:ignore
func TreeFromProto(t *rts.SubjectTree) (*Tree, error) {
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

	if t.NodeType != rts.NodeType_NODE_TYPE_LEAF {
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

func (t *Tree) Label() string {
	if t == nil {
		return ""
	}

	if t.Subject != nil {
		return t.Subject.String()
	} else if t.Tuple != nil {
		return t.Tuple.String()
	}
	return ""
}

func (t *Tree) String() string {
	if t == nil {
		return ""
	}

	nodeLabel := t.Label()

	if t.Type == Leaf {
		return fmt.Sprintf("∋ %s️", nodeLabel)
	}

	children := make([]string, len(t.Children))
	for i, c := range t.Children {
		var indent string
		if i == len(t.Children)-1 {
			indent = "   "
		} else {
			indent = "│  "
		}
		children[i] = strings.Join(strings.Split(c.String(), "\n"), "\n"+indent)
	}

	setOperation := ""
	switch t.Type {
	case Intersection:
		setOperation = "⋂"
	case Union:
		setOperation = "⋃"
	case Exclusion:
		setOperation = `\`
	case TupeToUserset:
		setOperation = "┐ tuple to userset"
	case ComputedUserset:
		setOperation = "┐ computed userset"
	}

	boxSymbol := "├"
	if len(children) == 1 {
		boxSymbol = "└"
	}
	return fmt.Sprintf("%s %s\n%s──%s", setOperation, nodeLabel, boxSymbol, strings.Join(children, "\n└──"))
}
