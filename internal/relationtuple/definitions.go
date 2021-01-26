package relationtuple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/tidwall/sjson"

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
		protoRelations    []*acl.RelationTuple
		internalRelations []*InternalRelationTuple
	}
	Subject interface {
		json.Marshaler

		String() string
		FromString(string) (Subject, error)
		Equals(interface{}) bool
		ToProto() *acl.Subject
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
	TupleData interface {
		GetSubject() *acl.Subject
		GetObject() string
		GetNamespace() string
		GetRelation() string
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

func SubjectFromProto(gs *acl.Subject) Subject {
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

func (s *SubjectSet) ToURLQuery() url.Values {
	return url.Values{
		"namespace": []string{s.Namespace},
		"object":    []string{s.Object},
		"relation":  []string{s.Relation},
	}
}

func (s *SubjectID) ToProto() *acl.Subject {
	return &acl.Subject{
		Ref: &acl.Subject_Id{
			Id: s.ID,
		},
	}
}

func (s *SubjectSet) ToProto() *acl.Subject {
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

func (s SubjectID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func (s SubjectSet) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func (r *InternalRelationTuple) String() string {
	return fmt.Sprintf("%s:%s#%s@%s", r.Namespace, r.Object, r.Relation, r.Subject)
}

func (r *InternalRelationTuple) DeriveSubject() *SubjectSet {
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

func (r *InternalRelationTuple) MarshalJSON() ([]byte, error) {
	type t InternalRelationTuple

	enc, err := json.Marshal((*t)(r))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return sjson.SetBytes(enc, "subject", r.Subject.String())
}

func (r *InternalRelationTuple) FromDataProvider(d TupleData) *InternalRelationTuple {
	r.Subject = SubjectFromProto(d.GetSubject())
	r.Object = d.GetObject()
	r.Namespace = d.GetNamespace()
	r.Relation = d.GetRelation()

	return r
}

func (r *InternalRelationTuple) ToProto() *acl.RelationTuple {
	return &acl.RelationTuple{
		Namespace: r.Namespace,
		Object:    r.Object,
		Relation:  r.Relation,
		Subject:   r.Subject.ToProto(),
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

func (r *InternalRelationTuple) ToURLQuery() url.Values {
	return url.Values{
		"namespace": []string{r.Namespace},
		"object":    []string{r.Object},
		"relation":  []string{r.Relation},
		"subject":   []string{r.Subject.String()},
	}
}

func (q *RelationQuery) FromProto(query *acl.ListRelationTuplesRequest_Query) *RelationQuery {
	return &RelationQuery{
		Namespace: query.Namespace,
		Object:    query.Object,
		Relation:  query.Relation,
		Subject:   SubjectFromProto(query.Subject),
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

func (q *RelationQuery) ToURLQuery() url.Values {
	v := make(url.Values, 4)

	if q.Namespace != "" {
		v.Add("namespace", q.Namespace)
	}
	if q.Relation != "" {
		v.Add("relation", q.Relation)
	}
	if q.Object != "" {
		v.Add("object", q.Object)
	}
	if q.Subject != nil {
		v.Add("subject", q.Subject.String())
	}

	return v
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

func (r *InternalRelationTuple) Columns() []string {
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

func NewProtoRelationCollection(rels []*acl.RelationTuple) cmdx.Table {
	return &relationCollection{
		protoRelations: rels,
	}
}

func NewRelationCollection(rels []*InternalRelationTuple) cmdx.Table {
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
		for _, rel := range r.protoRelations {
			r.internalRelations = append(r.internalRelations, (&InternalRelationTuple{}).FromDataProvider(rel))
		}
	}

	data := make([][]string, len(r.internalRelations))
	for i, rel := range r.internalRelations {
		data[i] = []string{rel.Namespace, rel.Object, rel.Relation, cmdx.None}
		if rel.Subject != nil {
			data[i][3] = rel.Subject.String()
		}
	}

	return data
}

func (r *relationCollection) Interface() interface{} {
	if r.internalRelations == nil {
		r.internalRelations = make([]*InternalRelationTuple, len(r.protoRelations))
		for i, rel := range r.protoRelations {
			r.internalRelations[i] = (&InternalRelationTuple{}).FromDataProvider(rel)
		}
	}
	return r.internalRelations
}

func (r *relationCollection) Len() int {
	// one of them is zero so the sum is always correct
	return len(r.protoRelations) + len(r.internalRelations)
}

type ManagerWrapper struct {
	Reg            ManagerProvider
	PageOpts       []x.PaginationOptionSetter
	RequestedPages []string
}

var (
	_ Manager         = (*ManagerWrapper)(nil)
	_ ManagerProvider = (*ManagerWrapper)(nil)
)

func NewManagerWrapper(_ *testing.T, reg ManagerProvider, options ...x.PaginationOptionSetter) *ManagerWrapper {
	return &ManagerWrapper{
		Reg:      reg,
		PageOpts: options,
	}
}

func (t *ManagerWrapper) GetRelationTuples(ctx context.Context, query *RelationQuery, options ...x.PaginationOptionSetter) ([]*InternalRelationTuple, string, error) {
	opts := x.GetPaginationOptions(options...)
	if opts.Token != "" {
		t.RequestedPages = append(t.RequestedPages, opts.Token)
	}
	return t.Reg.RelationTupleManager().GetRelationTuples(ctx, query, append(t.PageOpts, options...)...)
}

func (t *ManagerWrapper) WriteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error {
	return t.Reg.RelationTupleManager().WriteRelationTuples(ctx, rs...)
}

func (t *ManagerWrapper) RelationTupleManager() Manager {
	return t
}
