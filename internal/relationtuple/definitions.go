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
	Subject interface {
		String() string
		FromString(string) Subject
		Equals(interface{}) bool
	}
	UserID struct {
		ID string `json:"id"`
	}
	UserSet struct {
		Namespace string `json:"namespace"`
		ObjectID  string `json:"object_id"`
		Relation  string `json:"relation"`
	}
	InternalRelationTuple struct {
		Namespace string  `json:"namespace"`
		ObjectID  string  `json:"object_id"`
		Relation  string  `json:"relation"`
		Subject   Subject `json:"subject"`
	}
	RelationQuery struct {
		ObjectID  string  `json:"object_id"`
		Namespace string  `json:"namespace"`
		Relation  string  `json:"relation"`
		Subject   Subject `json:"subject"`
	}
)

var _, _ Subject = &UserID{}, &UserSet{}

func SubjectFromString(s string) Subject {
	if strings.Contains(s, "#") {
		return (&UserSet{}).FromString(s)
	}
	return (&UserID{}).FromString(s)
}

func (u *UserID) String() string {
	return u.ID
}

func (u *UserSet) String() string {
	return fmt.Sprintf("%s:%s#%s", u.Namespace, u.ObjectID, u.Relation)
}

func (u *UserID) FromString(s string) Subject {
	u.ID = s
	return u
}

func (u *UserSet) FromString(s string) Subject {
	parts := strings.Split(s, "#")
	if len(parts) == 2 {
		innerParts := strings.Split(parts[0], ":")
		if len(innerParts) == 2 {
			u.Namespace = innerParts[0]
			u.ObjectID = innerParts[1]
		}

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
	return uv.Relation == u.Relation && uv.ObjectID == u.ObjectID && uv.Namespace == u.Namespace
}

func (r *InternalRelationTuple) String() string {
	return fmt.Sprintf("%s:%s#%s@%s", r.Namespace, r.ObjectID, r.Relation, r.Subject)
}

func (r *InternalRelationTuple) DeriveSubject() Subject {
	return &UserSet{
		Namespace: r.Namespace,
		ObjectID:  r.ObjectID,
		Relation:  r.Relation,
	}
}

func (r *InternalRelationTuple) UnmarshalJSON(raw []byte) error {
	subject := gjson.GetBytes(raw, "subject").Str
	r.Subject = SubjectFromString(subject)
	r.ObjectID = gjson.GetBytes(raw, "object_id").Str
	r.Relation = gjson.GetBytes(raw, "relation").Str
	r.Namespace = gjson.GetBytes(raw, "namespace").Str

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
			Namespace: gr.GetUserSet().GetNamespace(),
			ObjectID:  gr.GetUserSet().GetObjectId(),
			Relation:  gr.GetUserSet().GetRelation(),
		}
	}

	r.ObjectID = gr.GetObjectId()
	r.Namespace = gr.GetNamespace()
	r.Relation = gr.GetRelation()
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
				ObjectId:  s.ObjectID,
				Namespace: s.Namespace,
			},
		}
	}

	x.ObjectId = r.ObjectID
	x.Namespace = r.Namespace
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
			ObjectID:  query.GetUserSet().GetObjectId(),
			Namespace: query.GetUserSet().GetNamespace(),
			Relation:  query.GetUserSet().GetRelation(),
		}
	}

	rq.ObjectID = query.ObjectId
	rq.Namespace = query.Namespace
	rq.Relation = query.Relation
	rq.Subject = subject

	return rq
}

func (r *InternalRelationTuple) Header() []string {
	return []string{
		"NAMESPACE",
		"OBJECT ID",
		"RELATION NAME",
		"SUBJECT ID",
	}
}

func (r *InternalRelationTuple) Fields() []string {
	return []string{
		r.Namespace,
		r.ObjectID,
		r.Relation,
		r.Subject.String(),
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
		"NAMESPACE",
		"OBJECT",
		"RELATION NAME",
		"SUBJECT",
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
		data[i] = []string{rel.Namespace, rel.ObjectID, rel.Relation, cmdx.None}
		if rel.Subject != nil {
			data[i][1] = rel.Subject.String()
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
