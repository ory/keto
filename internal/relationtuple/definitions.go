package relationtuple

import (
	"context"
	"fmt"
	"strings"

	"github.com/ory/keto/internal/x"

	"github.com/tidwall/gjson"

	"github.com/ory/x/cmdx"
)

type (
	ManagerProvider interface {
		RelationTupleManager() Manager
	}
	Manager interface {
		GetRelationTuples(ctx context.Context, query *RelationQuery, options ...x.PaginationOptionSetter) ([]*InternalRelationTuple, error)
		WriteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error
	}

	relationCollection struct {
		grpcRelations     []*RelationTuple
		internalRelations []*InternalRelationTuple
	}
	Object struct {
		ID        string `json:"id"`
		Namespace string `json:"namespace"`
	}
	Subject interface {
		String() string
		FromString(string) Subject
		Equals(interface{}) bool
	}
	UserID struct {
		ID string `json:"id"`
	}
	UserSet struct {
		Object   *Object `json:"object"`
		Relation string  `json:"relation"`
	}
	InternalRelationTuple struct {
		Object   *Object `json:"object"`
		Relation string  `json:"relation"`
		Subject  Subject `json:"subject"`
	}
	RelationQuery struct {
		Object   *Object `json:"object"`
		Relation string  `json:"relation"`
		Subject  Subject `json:"subject"`
	}
)

var _, _ Subject = &UserID{}, &UserSet{}

func SubjectFromString(s string) Subject {
	if strings.Contains(s, "#") {
		return (&UserSet{}).FromString(s)
	}
	return (&UserID{}).FromString(s)
}

func (o *Object) String() string {
	return fmt.Sprintf("%s:%s", o.Namespace, o.ID)
}

func (o *Object) Equals(v interface{}) bool {
	ov, ok := v.(*Object)
	if !ok {
		return false
	}
	return ov.ID == o.ID && ov.Namespace == o.Namespace
}

func (o *Object) UnmarshalJSON(raw []byte) error {
	o.FromString(string(raw))
	return nil
}

func (o *Object) FromString(s string) *Object {
	parts := strings.Split(s, ":")
	if len(parts) == 2 {
		o.Namespace, o.ID = parts[0], parts[1]
	}
	return o
}

func (x *RelationObject) FromString(s string) *RelationObject {
	o := (&Object{}).FromString(s)
	x.Namespace, x.ObjectId = o.Namespace, o.ID
	return x
}

func (u *UserID) String() string {
	return u.ID
}

func (u *UserSet) String() string {
	return fmt.Sprintf("%s#%s", u.Object, u.Relation)
}

func (u *UserID) FromString(s string) Subject {
	u.ID = s
	return u
}

func (u *UserSet) FromString(s string) Subject {
	parts := strings.Split(s, "#")
	if len(parts) == 2 {
		u.Object.FromString(parts[0])
		u.Relation = parts[1]
	}
	return u
}

func (u *UserID) Equals(v interface{}) bool {
	uv, ok := v.(*UserID)
	if !ok {
		return false
	}
	return uv.ID == u.ID
}

func (u *UserSet) Equals(v interface{}) bool {
	uv, ok := v.(*UserSet)
	if !ok {
		return false
	}
	return uv.Relation == u.Relation && uv.Object.Equals(u.Object)
}

func (r *InternalRelationTuple) String() string {
	return fmt.Sprintf("%s#%s@%s", r.Object, r.Relation, r.Subject)
}

func (r *InternalRelationTuple) DeriveSubject() Subject {
	return &UserSet{
		// TODO check if this should be copied
		Object:   r.Object,
		Relation: r.Relation,
	}
}

func (r *InternalRelationTuple) UnmarshalJSON(raw []byte) error {
	subject := gjson.GetBytes(raw, "subject").Str
	r.Subject = SubjectFromString(subject)
	object := gjson.GetBytes(raw, "object").Str
	r.Object = (&Object{}).FromString(object)
	r.Relation = gjson.GetBytes(raw, "relation").Str

	return nil
}

func (r *InternalRelationTuple) FromGRPC(gr *RelationTuple) *InternalRelationTuple {
	var subject Subject
	switch gr.Subject.(type) {
	case *RelationTuple_UserId:
		subject = &UserID{
			ID: gr.GetUserId(),
		}
	case *RelationTuple_UserSet:
		subject = &UserSet{
			Object: &Object{
				gr.GetUserSet().Object.Namespace,
				gr.GetUserSet().Object.ObjectId,
			},
			Relation: gr.GetUserSet().Relation,
		}
	}

	if gr.Object != nil {
		r.Object = &Object{
			gr.Object.Namespace,
			gr.Object.ObjectId,
		}
	}
	r.Relation = gr.Relation
	r.Subject = subject

	return r
}

func (x *RelationTuple) FromInternal(r *InternalRelationTuple) *RelationTuple {
	var subject isRelationTuple_Subject
	switch s := r.Subject.(type) {
	case *UserID:
		subject = &RelationTuple_UserId{
			s.ID,
		}
	case *UserSet:
		subject = &RelationTuple_UserSet{
			UserSet: &RelationUserSet{
				Object: &RelationObject{
					ObjectId:  s.Object.ID,
					Namespace: s.Object.Namespace,
				},
			},
		}
	}

	x.Object = &RelationObject{
		ObjectId:  r.Object.ID,
		Namespace: r.Object.Namespace,
	}
	x.Relation = r.Relation
	x.Subject = subject

	return x
}

func (rq *RelationQuery) FromGRPC(query *ReadRelationTuplesRequest_Query) *RelationQuery {
	var subject Subject
	switch query.Subject.(type) {
	case *ReadRelationTuplesRequest_Query_UserId:
		subject = &UserID{
			ID: query.GetUserId(),
		}
	case *ReadRelationTuplesRequest_Query_UserSet:
		subject = &UserSet{
			Object: &Object{
				query.GetUserSet().Object.Namespace,
				query.GetUserSet().Object.ObjectId,
			},
			Relation: query.GetUserSet().Relation,
		}
	}

	rq.Object = &Object{
		query.Object.Namespace,
		query.Object.ObjectId,
	}
	rq.Relation = query.Relation
	rq.Subject = subject

	return rq
}

func (r *InternalRelationTuple) Header() []string {
	return []string{
		"RELATION NAME",
		"SUBJECT ID",
		"OBJECT ID",
	}
}

func (r *InternalRelationTuple) Fields() []string {
	return []string{
		r.Relation,
		r.Subject.String(),
		r.Object.String(),
	}
}

func (r *InternalRelationTuple) Interface() interface{} {
	return r
}

func NewGRPCRelationCollection(rels []*RelationTuple) cmdx.OutputCollection {
	return &relationCollection{
		grpcRelations: rels,
	}
}

func NewRelationCollection(rels []*InternalRelationTuple) cmdx.OutputCollection {
	return &relationCollection{
		internalRelations: rels,
	}
}

func (r *relationCollection) Header() []string {
	return []string{
		"RELATION NAME",
		"SUBJECT",
		"OBJECT",
	}
}

func (r *relationCollection) Table() [][]string {
	if r.internalRelations == nil {
		for _, rel := range r.grpcRelations {
			r.internalRelations = append(r.internalRelations, (&InternalRelationTuple{}).FromGRPC(rel))
		}
	}

	data := make([][]string, len(r.internalRelations))
	for i, rel := range r.internalRelations {
		data[i] = []string{rel.Relation, cmdx.None, cmdx.None}
		if rel.Subject != nil {
			data[i][1] = rel.Subject.String()
		}
		if rel.Object != nil {
			data[i][2] = rel.Object.String()
		}
	}

	return data
}

func (r *relationCollection) Interface() interface{} {
	if r.internalRelations == nil {
		return r.grpcRelations
	}
	return r.internalRelations
}

func (r *relationCollection) Len() int {
	// one of them is zero so the sum is always correct
	return len(r.grpcRelations) + len(r.internalRelations)
}
