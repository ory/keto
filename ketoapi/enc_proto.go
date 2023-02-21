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
		GetObject() string
		GetNamespace() string
		GetRelation() string
		GetSubject() *rts.Subject
	}

	queryData interface {
		GetSubject() *rts.Subject
		GetObject() *string
		GetNamespace() *string
		GetRelation() *string
	}

	openAPIFields interface {
		GetObject() string
		GetNamespace() string
		GetRelation() string
		GetSubject() interface {
			GetSubjectId() string
			GetSubjectSet() *rts.SubjectSet
		}
	}

	OpenAPITupleData struct {
		Wrapped interface {
			GetObject() string
			GetRelation() string
			GetNamespace() string
			GetSubjectSet() *rts.SubjectSet
			GetSubjectId() string
		}
	}
)

func (q *OpenAPITupleData) GetObject() string    { return q.Wrapped.GetObject() }
func (q *OpenAPITupleData) GetNamespace() string { return q.Wrapped.GetNamespace() }
func (q *OpenAPITupleData) GetRelation() string  { return q.Wrapped.GetRelation() }
func (q *OpenAPITupleData) GetSubject() *rts.Subject {
	if sub, ok := q.Wrapped.(interface{ GetSubject() *rts.Subject }); ok && sub.GetSubject() != nil {
		return sub.GetSubject()
	}
	if set := q.Wrapped.GetSubjectSet(); set != nil {
		return rts.NewSubjectSet(set.Namespace, set.Object, set.Relation)
	}
	return rts.NewSubjectID(q.Wrapped.GetSubjectId())
}

func (r *RelationTuple) FromOpenAPIFields(f openAPIFields) (*RelationTuple, error) {
	subject := f.GetSubject()
	if subject == nil {
		return nil, errors.WithStack(ErrNilSubject)
	}
	if subjectSet := subject.GetSubjectSet(); subjectSet != nil {
		r.SubjectSet = &SubjectSet{
			Namespace: subjectSet.Namespace,
			Object:    subjectSet.Object,
			Relation:  subjectSet.Relation,
		}
	} else {
		r.SubjectID = pointerx.Ptr(subject.GetSubjectId())
	}

	r.Object = f.GetObject()
	r.Namespace = f.GetNamespace()
	r.Relation = f.GetRelation()

	return r, nil
}

func (r *RelationTuple) FromDataProvider(d TupleData) (*RelationTuple, error) {
	switch s := d.GetSubject().GetRef().(type) {
	case *rts.Subject_Set:
		r.SubjectSet = &SubjectSet{
			Namespace: s.Set.Namespace,
			Object:    s.Set.Object,
			Relation:  s.Set.Relation,
		}
	case *rts.Subject_Id:
		r.SubjectID = pointerx.Ptr(s.Id)
	default:
		return nil, errors.WithStack(ErrNilSubject)
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
		res.RestApiSubject = &rts.RelationTuple_SubjectId{SubjectId: *r.SubjectID}
	} else {
		res.Subject = rts.NewSubjectSet(r.SubjectSet.Namespace, r.SubjectSet.Object, r.SubjectSet.Relation)
		res.RestApiSubject = &rts.RelationTuple_SubjectSet{
			SubjectSet: &rts.SubjectSet{
				Namespace: r.SubjectSet.Namespace,
				Object:    r.SubjectSet.Object,
				Relation:  r.SubjectSet.Relation,
			},
		}
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

func (r *RelationTuple) FromCheckRequest(proto *rts.CheckRequest) *RelationTuple {
	if proto.Tuple != nil {
		return r.FromProto(proto.Tuple)
	}

	r = &RelationTuple{
		Namespace: proto.Namespace,
		Object:    proto.Object,
		Relation:  proto.Relation,
	}

	if proto.Subject != nil {
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
	} else {
		switch subject := proto.RestApiSubject.(type) {
		case *rts.CheckRequest_SubjectId:
			r.SubjectID = pointerx.Ptr(subject.SubjectId)
		case *rts.CheckRequest_SubjectSet:
			r.SubjectSet = &SubjectSet{
				Namespace: subject.SubjectSet.Namespace,
				Object:    subject.SubjectSet.Object,
				Relation:  subject.SubjectSet.Relation,
			}
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
	// nolint - fill deprecated field
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
		// nolint
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
