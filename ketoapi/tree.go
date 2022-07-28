package ketoapi

import (
	"fmt"
	"strings"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

// swagger:enum TreeNodeType
type TreeNodeType string

const (
	TreeNodeUnion           TreeNodeType = "union"
	TreeNodeExclusion       TreeNodeType = "exclusion"
	TreeNodeIntersection    TreeNodeType = "intersection"
	TreeNodeLeaf            TreeNodeType = "leaf"
	TreeNodeTupeToUserset   TreeNodeType = "tuple_to_userset"
	TreeNodeComputedUserset TreeNodeType = "computed_userset"
	TreeNodeNot             TreeNodeType = "not"
	TreeNodeUnspecified     TreeNodeType = "unspecified"
)

type Tuple[T any] interface {
	fmt.Stringer
	ToProto() *rts.RelationTuple
	FromProto(*rts.RelationTuple) T
}

// Tree is a generic tree of either internal relation tuples (with UUIDs for
// objects, etc.) or API relation tuples (with strings for objects, etc.).
type Tree[T Tuple[T]] struct {
	// Propagate all struct changes to `SwaggerOnlyExpandTree` as well.

	// The type of the node.
	//
	// required: true
	Type TreeNodeType `json:"type"`

	// The children of the node, possibly none.
	Children []*Tree[T] `json:"children,omitempty"`

	// The relation tuple the node represents.
	Tuple T `json:"tuple"`
}

// IMPORTANT: We need a manual instantiation of the generic Tree[T] for the
// OpenAPI spec, since go-swagger does not understand generics :(.
// This can be fixed by using grpc-gateway.
//
// swagger:model expandTree
type SwaggerOnlyExpandTree struct { // nolint
	// The type of the node.
	//
	// required: true
	Type TreeNodeType `json:"type"`

	// The children of the node, possibly none.
	Children []*SwaggerOnlyExpandTree `json:"children,omitempty"`

	// The relation tuple the node represents.
	Tuple *RelationTuple `json:"tuple"`
}

func (t TreeNodeType) String() string {
	return string(t)
}

func (t *TreeNodeType) UnmarshalJSON(v []byte) error {
	switch string(v) {
	case `"union"`:
		*t = TreeNodeUnion
	case `"exclusion"`:
		*t = TreeNodeExclusion
	case `"intersection"`:
		*t = TreeNodeIntersection
	case `"leaf"`:
		*t = TreeNodeLeaf
	case `"tuple_to_userset"`:
		*t = TreeNodeTupeToUserset
	case `"computed_userset"`:
		*t = TreeNodeComputedUserset
	case `"not"`:
		*t = TreeNodeNot
	case `"unspecified"`:
		*t = TreeNodeUnspecified
	default:
		return ErrUnknownNodeType
	}
	return nil
}

func (t TreeNodeType) ToProto() rts.NodeType {
	switch t {
	case TreeNodeLeaf:
		return rts.NodeType_NODE_TYPE_LEAF
	case TreeNodeUnion:
		return rts.NodeType_NODE_TYPE_UNION
	case TreeNodeExclusion:
		return rts.NodeType_NODE_TYPE_EXCLUSION
	case TreeNodeIntersection:
		return rts.NodeType_NODE_TYPE_INTERSECTION
	}
	return rts.NodeType_NODE_TYPE_UNSPECIFIED
}

func (TreeNodeType) FromProto(pt rts.NodeType) TreeNodeType {
	switch pt {
	case rts.NodeType_NODE_TYPE_LEAF:
		return TreeNodeLeaf
	case rts.NodeType_NODE_TYPE_UNION:
		return TreeNodeUnion
	case rts.NodeType_NODE_TYPE_EXCLUSION:
		return TreeNodeExclusion
	case rts.NodeType_NODE_TYPE_INTERSECTION:
		return TreeNodeIntersection
	}
	return TreeNodeUnspecified
}

func (t *Tree[NodeT]) Label() string {
	if t == nil {
		return ""
	}

	return t.Tuple.String()
}

func (t *Tree[NodeT]) String() string {
	if t == nil {
		return ""
	}

	nodeLabel := t.Label()

	if t.Type == TreeNodeLeaf {
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
	case TreeNodeIntersection:
		setOperation = "and"
	case TreeNodeUnion:
		setOperation = "or"
	case TreeNodeExclusion:
		setOperation = `\`
	case TreeNodeNot:
		setOperation = `not`
	case TreeNodeTupeToUserset:
		setOperation = "┐ tuple to userset"
	case TreeNodeComputedUserset:
		setOperation = "┐ computed userset"
	}

	boxSymbol := "├"
	if len(children) == 1 {
		boxSymbol = "└"
	}
	return fmt.Sprintf("%s %s\n%s──%s", setOperation, nodeLabel, boxSymbol, strings.Join(children, "\n└──"))
}
