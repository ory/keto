// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoapi

import (
	"github.com/ory/x/pointerx"
	"github.com/pkg/errors"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	TupleData interface {
		GetSubject() *rts.Subject
		GetObject() string
		GetNamespace() string
		GetRelation() string
	}
	queryData interface {
		GetSubject() *rts.Subject
		GetObject() *string
		GetNamespace() *string
		GetRelation() *string
	}
)

func (r *RelationTuple) FromDataProvider(d TupleData) (*RelationTuple, error) {
	switch s := d.GetSubject().GetRef().(type) {
	case nil:
		return nil, errors.WithStack(ErrNilSubject)
	case *rts.Subject_Set:
		r.SubjectSet = &SubjectSet{
			Namespace: s.Set.Namespace,
			Object:    s.Set.Object,
			Relation:  s.Set.Relation,
		}
	case *rts.Subject_Id:
		r.SubjectID = pointerx.Ptr(s.Id)
	}

	r.Object = d.GetObject()
	r.Namespace = d.GetNamespace()
	r.Relation = d.GetRelation()

	return r, nil
}

func (r *RelationTuple) ToProto() *rts.RelationTuple {
	res := &rts.RelationTuple{
		Namespace: r.Namespace,
		Object:    r.Object,
		Relation:  r.Relation,
	}
	if r.SubjectID != nil {
		res.Subject = rts.NewSubjectID(*r.SubjectID)
	} else {
		res.Subject = rts.NewSubjectSet(r.SubjectSet.Namespace, r.SubjectSet.Object, r.SubjectSet.Relation)
	}
	return res
}

func (r *RelationTuple) FromProto(proto *rts.RelationTuple) *RelationTuple {
	r = &RelationTuple{
		Namespace: proto.Namespace,
		Object:    proto.Object,
		Relation:  proto.Relation,
	}
	switch subject := proto.Subject.Ref.(type) {
	case *rts.Subject_Id:
		r.SubjectID = pointerx.Ptr(subject.Id)
	case *rts.Subject_Set:
		r.SubjectSet = &SubjectSet{
			Namespace: subject.Set.Namespace,
			Object:    subject.Set.Object,
			Relation:  subject.Set.Relation,
		}
	}

	return r
}

func (q *RelationQuery) FromDataProvider(d queryData) *RelationQuery {
	q.Namespace = d.GetNamespace()
	q.Object = d.GetObject()
	q.Relation = d.GetRelation()
	q.SubjectID = nil
	q.SubjectSet = nil

	if s := d.GetSubject(); s != nil {
		switch sub := s.Ref.(type) {
		case *rts.Subject_Id:
			q.SubjectID = pointerx.Ptr(sub.Id)
		case *rts.Subject_Set:
			q.SubjectSet = &SubjectSet{
				Namespace: sub.Set.Namespace,
				Object:    sub.Set.Object,
				Relation:  sub.Set.Relation,
			}
		}
	}
	return q
}

func (q *RelationQuery) ToProto() *rts.RelationQuery {
	res := &rts.RelationQuery{
		Namespace: q.Namespace,
		Object:    q.Object,
		Relation:  q.Relation,
	}
	if q.SubjectID != nil {
		res.Subject = rts.NewSubjectID(*q.SubjectID)
	} else if q.SubjectSet != nil {
		res.Subject = rts.NewSubjectSet(q.SubjectSet.Namespace, q.SubjectSet.Object, q.SubjectSet.Relation)
	}
	return res
}

func (t *Tree[NodeT]) ToProto() *rts.SubjectTree {
	res := &rts.SubjectTree{
		NodeType: t.Type.ToProto(),
		Children: make([]*rts.SubjectTree, len(t.Children)),
	}
	res.Tuple = t.Tuple.ToProto()
	//lint:ignore SA1019 backwards compatibility
	//nolint:staticcheck
	res.Subject = res.Tuple.Subject
	for i := range t.Children {
		res.Children[i] = t.Children[i].ToProto()
	}
	return res
}

func TreeFromProto[T tuple[T]](pt *rts.SubjectTree) *Tree[T] {
	t := new(Tree[T])
	t.Type = TreeNodeType("").FromProto(pt.NodeType)

	var tuple T
	if pt.Tuple == nil {
		// legacy case: fetch from deprecated fields
		//lint:ignore SA1019 backwards compatibility
		//nolint:staticcheck
		switch sub := pt.Subject.Ref.(type) {
		case *rts.Subject_Id:
			pt.Tuple.Subject = rts.NewSubjectID(sub.Id)
		case *rts.Subject_Set:
			pt.Tuple.Subject = rts.NewSubjectSet(
				sub.Set.Namespace,
				sub.Set.Object,
				sub.Set.Relation,
			)
		}
	}
	t.Tuple = tuple.FromProto(pt.Tuple)

	t.Children = make([]*Tree[T], len(pt.Children))
	for i := range pt.Children {
		t.Children[i] = TreeFromProto[T](pt.Children[i])
	}

	return t
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
