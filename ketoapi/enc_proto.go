package ketoapi

import (
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/x"
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
		r.SubjectID = x.Ptr(s.Id)
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

func (q *RelationQuery) FromDataProvider(d queryData) *RelationQuery {
	q.Namespace = d.GetNamespace()
	q.Object = d.GetObject()
	q.Relation = d.GetRelation()
	q.SubjectID = nil
	q.SubjectSet = nil

	if s := d.GetSubject(); s != nil {
		switch sub := s.Ref.(type) {
		case *rts.Subject_Id:
			q.SubjectID = x.Ptr(sub.Id)
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

func (t *ExpandTree) ToProto() *rts.SubjectTree {
	res := &rts.SubjectTree{
		NodeType: t.Type.ToProto(),
		Children: make([]*rts.SubjectTree, len(t.Children)),
	}
	if t.SubjectID != nil {
		res.Subject = rts.NewSubjectID(*t.SubjectID)
	} else {
		res.Subject = rts.NewSubjectSet(t.SubjectSet.Namespace, t.SubjectSet.Object, t.SubjectSet.Relation)
	}
	for i := range t.Children {
		res.Children[i] = t.Children[i].ToProto()
	}
	return res
}

func (t *ExpandTree) FromProto(pt *rts.SubjectTree) *ExpandTree {
	t.Type = ExpandNodeType("").FromProto(pt.NodeType)

	switch sub := pt.Subject.Ref.(type) {
	case *rts.Subject_Id:
		t.SubjectID = x.Ptr(sub.Id)
	case *rts.Subject_Set:
		t.SubjectSet = &SubjectSet{
			Namespace: sub.Set.Namespace,
			Object:    sub.Set.Object,
			Relation:  sub.Set.Relation,
		}
	}

	t.Children = make([]*ExpandTree, len(pt.Children))
	for i := range pt.Children {
		t.Children[i] = (&ExpandTree{}).FromProto(pt.Children[i])
	}

	return t
}
