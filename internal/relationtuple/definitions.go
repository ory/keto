package relationtuple

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"

	"github.com/ory/keto/internal/x"

	"github.com/tidwall/gjson"

	"github.com/ory/x/cmdx"
)

type (
	ManagerProvider interface {
		RelationTupleManager() Manager
	}
	Manager interface {
		GetRelationTuples(ctx context.Context, query *RelationQuery, options ...x.PaginationOptionSetter) ([]*InternalRelationTuple, string, error)
		WriteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error
	}

	relationCollection struct {
		grpcRelations     []*acl.RelationTuple
		internalRelations []*InternalRelationTuple
	}
	Subject interface {
		String() string
		FromString(string) (Subject, error)
		Equals(interface{}) bool
		ToGRPC() *acl.Subject
	}
	SubjectID struct {
		ID string `json:"id"`
	}
	SubjectSet struct {
		Namespace string `json:"namespace"`
		Object    string `json:"object"`
		Relation  string `json:"relation"`
	}
	InternalRelationTuple struct {
		Namespace string  `json:"namespace"`
		Object    string  `json:"object"`
		Relation  string  `json:"relation"`
		Subject   Subject `json:"subject"`
	}
	RelationQuery struct {
		Namespace string  `json:"namespace"`
		Object    string  `json:"object"`
		Relation  string  `json:"relation"`
		Subject   Subject `json:"subject"`
	}
)

var (
	_, _ Subject = &SubjectID{}, &SubjectSet{}

	ErrMalformedInput = errors.New("malformed string input")
)

func SubjectFromString(s string) (Subject, error) {
	if strings.Contains(s, "#") {
		return (&SubjectSet{}).FromString(s)
	}
	return (&SubjectID{}).FromString(s)
}

func SubjectFromGRPC(gs *acl.Subject) Subject {
	switch s := gs.GetRef().(type) {
	case *acl.Subject_Id:
		return &SubjectID{
			ID: s.Id,
		}
	case *acl.Subject_Set:
		return &SubjectSet{
			Namespace: s.Set.Namespace,
			Object:    s.Set.Object,
			Relation:  s.Set.Relation,
		}
	}
	return nil
}

func (s *SubjectID) String() string {
	return s.ID
}

func (s *SubjectSet) String() string {
	return fmt.Sprintf("%s:%s#%s", s.Namespace, s.Object, s.Relation)
}

func (s *SubjectID) FromString(str string) (Subject, error) {
	s.ID = str
	return s, nil
}

func (s *SubjectSet) FromString(str string) (Subject, error) {
	parts := strings.Split(str, "#")
	if len(parts) != 2 {
		return nil, errors.WithStack(ErrMalformedInput)
	}

	innerParts := strings.Split(parts[0], ":")
	if len(innerParts) != 2 {
		return nil, errors.WithStack(ErrMalformedInput)
	}

	s.Namespace = innerParts[0]
	s.Object = innerParts[1]
	s.Relation = parts[1]

	return s, nil
}

func (s *SubjectSet) FromURLQuery(values url.Values) *SubjectSet {
	if s == nil {
		s = &SubjectSet{}
	}

	s.Namespace = values.Get("namespace")
	s.Relation = values.Get("relation")
	s.Object = values.Get("object")

	return s
}

func (s *SubjectID) ToGRPC() *acl.Subject {
	return &acl.Subject{
		Ref: &acl.Subject_Id{
			Id: s.ID,
		},
	}
}

func (s *SubjectSet) ToGRPC() *acl.Subject {
	return &acl.Subject{
		Ref: &acl.Subject_Set{
			Set: &acl.SubjectSet{
				Namespace: s.Namespace,
				Object:    s.Object,
				Relation:  s.Relation,
			},
		},
	}
}

func (s *SubjectID) Equals(v interface{}) bool {
	uv, ok := v.(*SubjectID)
	if !ok {
		return false
	}
	return uv.ID == s.ID
}

func (s *SubjectSet) Equals(v interface{}) bool {
	uv, ok := v.(*SubjectSet)
	if !ok {
		return false
	}
	return uv.Relation == s.Relation && uv.Object == s.Object && uv.Namespace == s.Namespace
}

func (r *InternalRelationTuple) String() string {
	return fmt.Sprintf("%s:%s#%s@%s", r.Namespace, r.Object, r.Relation, r.Subject)
}

func (r *InternalRelationTuple) DeriveSubject() Subject {
	return &SubjectSet{
		Namespace: r.Namespace,
		Object:    r.Object,
		Relation:  r.Relation,
	}
}

func (r *InternalRelationTuple) UnmarshalJSON(raw []byte) error {
	subject := gjson.GetBytes(raw, "subject").Str

	var err error
	r.Subject, err = SubjectFromString(subject)
	if err != nil {
		return err
	}

	r.Namespace = gjson.GetBytes(raw, "namespace").Str
	r.Object = gjson.GetBytes(raw, "object").Str
	r.Relation = gjson.GetBytes(raw, "relation").Str

	return nil
}

func (r *InternalRelationTuple) FromGRPC(gr *acl.RelationTuple) *InternalRelationTuple {
	r.Subject = SubjectFromGRPC(gr.Subject)
	r.Object = gr.Object
	r.Namespace = gr.Namespace
	r.Relation = gr.Relation

	return r
}

func (r *InternalRelationTuple) ToGRPC() *acl.RelationTuple {
	return &acl.RelationTuple{
		Namespace: r.Namespace,
		Object:    r.Object,
		Relation:  r.Relation,
		Subject:   r.Subject.ToGRPC(),
	}
}

func (r *InternalRelationTuple) FromURLQuery(query url.Values) (*InternalRelationTuple, error) {
	if s := query.Get("subject"); s != "" {
		var err error
		r.Subject, err = SubjectFromString(s)
		if err != nil {
			return nil, err
		}
	}

	r.Object = query.Get("object")
	r.Relation = query.Get("relation")
	r.Namespace = query.Get("namespace")

	return r, nil
}

func (q *RelationQuery) FromGRPC(query *acl.ListRelationTuplesRequest_Query) *RelationQuery {
	return &RelationQuery{
		Namespace: query.Namespace,
		Object:    query.Object,
		Relation:  query.Relation,
		Subject:   SubjectFromGRPC(query.Subject),
	}
}

func (q *RelationQuery) FromURLQuery(query url.Values) (*RelationQuery, error) {
	if q == nil {
		q = &RelationQuery{}
	}

	if s := query.Get("subject"); s != "" {
		var err error
		q.Subject, err = SubjectFromString(s)
		if err != nil {
			return nil, err
		}
	}

	q.Object = query.Get("object")
	q.Relation = query.Get("relation")
	q.Namespace = query.Get("namespace")

	return q, nil
}

func (q *RelationQuery) String() string {
	return fmt.Sprintf("namespace: %s; object: %s; relation: %s; subject: %s", q.Namespace, q.Object, q.Relation, q.Subject)
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
		r.Object,
		r.Relation,
		r.Subject.String(),
	}
}

func (r *InternalRelationTuple) Interface() interface{} {
	return r
}

func NewGRPCRelationCollection(rels []*acl.RelationTuple) cmdx.OutputCollection {
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
		data[i] = []string{rel.Namespace, rel.Object, rel.Relation, cmdx.None}
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
