package relationtuple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/ory/herodot"

	"github.com/sirupsen/logrus"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

	"github.com/tidwall/sjson"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/x"

	"github.com/tidwall/gjson"
)

type (
	ManagerProvider interface {
		RelationTupleManager() Manager
	}
	Manager interface {
		GetRelationTuples(ctx context.Context, query *RelationQuery, options ...x.PaginationOptionSetter) ([]*InternalRelationTuple, string, error)
		WriteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error
		DeleteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error
		TransactRelationTuples(ctx context.Context, insert []*InternalRelationTuple, delete []*InternalRelationTuple) error
	}

	RelationCollection struct {
		protoRelations    []*acl.RelationTuple
		internalRelations []*InternalRelationTuple
	}
	SubjectID struct {
		ID string `json:"id"`
	}
)

type RelationQuery struct {
	// required: true
	Namespace string  `json:"namespace"`
	Object    string  `json:"object"`
	Relation  string  `json:"relation"`
	Subject   Subject `json:"subject"`
}

// swagger:ignore
type TupleData interface {
	// swagger:ignore
	GetSubject() *acl.Subject
	GetObject() string
	GetNamespace() string
	GetRelation() string
}

// swagger:model subject
type Subject interface {
	// swagger:ignore
	json.Marshaler

	// swagger:ignore
	String() string
	// swagger:ignore
	FromString(string) (Subject, error)
	// swagger:ignore
	Equals(interface{}) bool

	// swagger:ignore
	ToProto() *acl.Subject
}

// swagger:parameters getCheck deleteRelationTuple
type InternalRelationTuple struct {
	// Namespace of the Relation Tuple
	//
	// in: query
	// required: true
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	//
	// in: query
	// required: true
	Object string `json:"object"`

	// Relation of the Relation Tuple
	//
	// in: query
	// required: true
	Relation string `json:"relation"`

	// Subject of the Relation Tuple
	//
	// The subject follows the subject string encoding format.
	//
	// in: query
	// required: true
	Subject Subject `json:"subject"`
}

// swagger:model subject
// nolint:deadcode,unused
type stringEncodedSubject string

// swagger:parameters getExpand
type SubjectSet struct {
	// Namespace of the Relation Tuple
	//
	// in: query
	// required: true
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	//
	// in: query
	// required: true
	Object string `json:"object"`

	// Relation of the Relation Tuple
	//
	// in: query
	// required: true
	Relation string `json:"relation"`
}

var (
	_, _ Subject = &SubjectID{}, &SubjectSet{}

	ErrMalformedInput = herodot.ErrBadRequest.WithError("malformed string input")
	ErrNilSubject     = herodot.ErrBadRequest.WithError("subject is not allowed to be nil")
)

// swagger:enum patchAction
type patchAction string

const (
	ActionInsert patchAction = "insert"
	ActionDelete patchAction = "delete"
)

// The patch request payload
//
// swagger:parameters patchRelationTuples
// nolint:deadcode,unused
type patchPayload struct {
	// in:body
	Payload []*PatchDelta
}

type PatchDelta struct {
	Action        patchAction            `json:"action"`
	RelationTuple *InternalRelationTuple `json:"relation_tuple"`
}

func SubjectFromString(s string) (Subject, error) {
	if strings.Contains(s, "#") {
		return (&SubjectSet{}).FromString(s)
	}
	return (&SubjectID{}).FromString(s)
}

// swagger:ignore
func SubjectFromProto(gs *acl.Subject) (Subject, error) {
	switch s := gs.GetRef().(type) {
	case nil:
		return nil, errors.WithStack(ErrNilSubject)
	case *acl.Subject_Id:
		return &SubjectID{
			ID: s.Id,
		}, nil
	case *acl.Subject_Set:
		return &SubjectSet{
			Namespace: s.Set.Namespace,
			Object:    s.Set.Object,
			Relation:  s.Set.Relation,
		}, nil
	}
	return nil, errors.WithStack(ErrNilSubject)
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

// swagger:ignore
func (s *SubjectID) ToProto() *acl.Subject {
	return &acl.Subject{
		Ref: &acl.Subject_Id{
			Id: s.ID,
		},
	}
}

// swagger:ignore
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

func (r *InternalRelationTuple) FromString(s string) (*InternalRelationTuple, error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return nil, errors.Wrap(ErrMalformedInput, "expected input to contain ':'")
	}
	r.Namespace = parts[0]

	parts = strings.SplitN(parts[1], "#", 2)
	if len(parts) != 2 {
		return nil, errors.Wrap(ErrMalformedInput, "expected input to contain '#'")
	}
	r.Object = parts[0]

	parts = strings.SplitN(parts[1], "@", 2)
	if len(parts) != 2 {
		return nil, errors.Wrap(ErrMalformedInput, "expected input to contain '@'")
	}
	r.Relation = parts[0]

	// remove optional brackets around the subject set
	sub := strings.Trim(parts[1], "()")

	var err error
	r.Subject, err = SubjectFromString(sub)
	if err != nil {
		return nil, err
	}

	return r, nil
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

func (r *InternalRelationTuple) FromDataProvider(d TupleData) (*InternalRelationTuple, error) {
	var err error
	r.Subject, err = SubjectFromProto(d.GetSubject())
	if err != nil {
		return nil, err
	}

	r.Object = d.GetObject()
	r.Namespace = d.GetNamespace()
	r.Relation = d.GetRelation()

	return r, nil
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
	vals := url.Values{
		"namespace": []string{r.Namespace},
		"object":    []string{r.Object},
		"relation":  []string{r.Relation},
	}
	if r.Subject != nil {
		vals.Set("subject", r.Subject.String())
	}
	return vals
}

func (r *InternalRelationTuple) ToLoggerFields() logrus.Fields {
	return logrus.Fields{
		"namespace": r.Namespace,
		"object":    r.Object,
		"relation":  r.Relation,
		"subject":   r.Subject,
	}
}

func (q *RelationQuery) FromProto(query *acl.ListRelationTuplesRequest_Query) (*RelationQuery, error) {
	r, err := (&InternalRelationTuple{}).FromDataProvider(query)
	if err != nil {
		return nil, err
	}

	*q = RelationQuery(*r)
	return q, nil
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

func NewProtoRelationCollection(rels []*acl.RelationTuple) *RelationCollection {
	return &RelationCollection{
		protoRelations: rels,
	}
}

func NewRelationCollection(rels []*InternalRelationTuple) *RelationCollection {
	return &RelationCollection{
		internalRelations: rels,
	}
}

func (r *RelationCollection) Header() []string {
	return []string{
		"NAMESPACE",
		"OBJECT",
		"RELATION NAME",
		"SUBJECT",
	}
}

func (r *RelationCollection) Table() [][]string {
	ir, err := r.ToInternal()
	if err != nil {
		return [][]string{{fmt.Sprintf("%+v", err)}}
	}

	data := make([][]string, len(ir))
	for i, rel := range ir {
		data[i] = []string{rel.Namespace, rel.Object, rel.Relation, rel.Subject.String()}
	}

	return data
}

func (r *RelationCollection) ToInternal() ([]*InternalRelationTuple, error) {
	if r.internalRelations == nil {
		r.internalRelations = make([]*InternalRelationTuple, len(r.protoRelations))
		for i, rel := range r.protoRelations {
			ir, err := (&InternalRelationTuple{}).FromDataProvider(rel)
			if err != nil {
				return nil, err
			}
			r.internalRelations[i] = ir
		}
	}
	return r.internalRelations, nil
}

func (r *RelationCollection) Interface() interface{} {
	ir, err := r.ToInternal()
	if err != nil {
		return err
	}
	return ir
}

func (r *RelationCollection) MarshalJSON() ([]byte, error) {
	ir, err := r.ToInternal()
	if err != nil {
		return nil, err
	}
	return json.Marshal(ir)
}

func (r *RelationCollection) UnmarshalJSON(raw []byte) error {
	return json.Unmarshal(raw, &r.internalRelations)
}

func (r *RelationCollection) Len() int {
	if ir := len(r.internalRelations); ir > 0 {
		return ir
	}
	return len(r.protoRelations)
}

func (r *RelationCollection) IDs() []string {
	rts, err := r.ToInternal()
	if err != nil {
		// fmt.Sprintf to include the stacktrace
		return []string{fmt.Sprintf("%+v", err)}
	}
	ids := make([]string, len(rts))
	for i, rt := range rts {
		ids[i] = rt.String()
	}
	return ids
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
	t.RequestedPages = append(t.RequestedPages, opts.Token)
	return t.Reg.RelationTupleManager().GetRelationTuples(ctx, query, append(t.PageOpts, options...)...)
}

func (t *ManagerWrapper) WriteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error {
	return t.Reg.RelationTupleManager().WriteRelationTuples(ctx, rs...)
}

func (t *ManagerWrapper) DeleteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error {
	return t.Reg.RelationTupleManager().DeleteRelationTuples(ctx, rs...)
}

func (t *ManagerWrapper) TransactRelationTuples(ctx context.Context, insert []*InternalRelationTuple, delete []*InternalRelationTuple) error {
	return t.Reg.RelationTupleManager().TransactRelationTuples(ctx, insert, delete)
}

func (t *ManagerWrapper) RelationTupleManager() Manager {
	return t
}
