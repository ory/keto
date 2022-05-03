package relationtuple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/x/pointerx"

	"github.com/ory/herodot"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/x"
)

type (
	ManagerProvider interface {
		RelationTupleManager() Manager
		UUIDMappingManager() UUIDMappingManager
	}
	Manager interface {
		GetRelationTuples(ctx context.Context, query *RelationQuery, options ...x.PaginationOptionSetter) ([]*InternalRelationTuple, string, error)
		WriteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error
		DeleteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error
		DeleteAllRelationTuples(ctx context.Context, query *RelationQuery) error
		TransactRelationTuples(ctx context.Context, insert []*InternalRelationTuple, delete []*InternalRelationTuple) error
	}

	RelationCollection struct {
		protoRelations    []*rts.RelationTuple
		internalRelations []*InternalRelationTuple
	}
	SubjectID struct {
		ID string `json:"id"`
	}
)

type RelationQuery struct {
	// Namespace of the Relation Tuple
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	Object string `json:"object"`

	// Relation of the Relation Tuple
	Relation string `json:"relation"`

	// SubjectID of the Relation Tuple
	//
	// Either SubjectSet or SubjectID can be provided.
	SubjectID *string `json:"subject_id,omitempty"`
	// SubjectSet of the Relation Tuple
	//
	// Either SubjectSet or SubjectID can be provided.
	//
	// swagger:allOf
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
}

// swagger:ignore
type TupleData interface {
	// swagger:ignore
	GetSubject() *rts.Subject
	GetObject() string
	GetNamespace() string
	GetRelation() string
}

// swagger:model subject
type Subject interface {
	// swagger:ignore
	String() string
	// swagger:ignore
	FromString(string) (Subject, error)
	// swagger:ignore
	Equals(interface{}) bool
	// swagger:ignore
	SubjectID() *string
	// swagger:ignore
	SubjectSet() *SubjectSet

	// swagger:ignore
	ToProto() *rts.Subject

	// swagger:ignore
	UUIDMappable
}

// swagger:ignore
type InternalRelationTuple struct {
	Namespace string  `json:"namespace"`
	Object    string  `json:"object"`
	Relation  string  `json:"relation"`
	Subject   Subject `json:"subject"`
}
type InternalRelationTuples []*InternalRelationTuple

func (rt *InternalRelationTuple) UUIDMappableFields() []*string {
	return append([]*string{&rt.Object}, rt.Subject.UUIDMappableFields()...)
}

func (rtt InternalRelationTuples) UUIDMappableFields() (fields []*string) {
	for _, rt := range rtt {
		fields = append(fields, rt.UUIDMappableFields()...)
	}
	return fields
}

// swagger:parameters getExpand
type SubjectSet struct {
	// Namespace of the Subject Set
	//
	// required: true
	Namespace string `json:"namespace"`

	// Object of the Subject Set
	//
	// required: true
	Object string `json:"object"`

	// Relation of the Subject Set
	//
	// required: true
	Relation string `json:"relation"`
}

var (
	_, _ Subject = &SubjectID{}, &SubjectSet{}

	ErrMalformedInput    = herodot.ErrBadRequest.WithError("malformed string input")
	ErrNilSubject        = herodot.ErrBadRequest.WithError("subject is not allowed to be nil")
	ErrDuplicateSubject  = herodot.ErrBadRequest.WithError("exactly one of subject_set or subject_id has to be provided")
	ErrDroppedSubjectKey = herodot.ErrBadRequest.WithDebug(`provide "subject_id" or "subject_set.*"; support for "subject" was dropped`)
	ErrIncompleteSubject = herodot.ErrBadRequest.WithError(`incomplete subject, provide "subject_id" or a complete "subject_set.*"`)
)

// swagger:enum patchAction
type patchAction string

const (
	ActionInsert patchAction = "insert"
	ActionDelete patchAction = "delete"
)

func SubjectFromString(s string) (Subject, error) {
	if strings.Contains(s, "#") {
		return (&SubjectSet{}).FromString(s)
	}
	return (&SubjectID{}).FromString(s)
}

// swagger:ignore
func (s *SubjectID) UUIDMappableFields() []*string {
	return []*string{&s.ID}
}

// swagger:ignore
func (s *SubjectSet) UUIDMappableFields() []*string {
	return []*string{&s.Object}
}

// swagger:ignore
func SubjectFromProto(gs *rts.Subject) (Subject, error) {
	switch s := gs.GetRef().(type) {
	case nil:
		return nil, errors.WithStack(ErrNilSubject)
	case *rts.Subject_Id:
		return &SubjectID{
			ID: s.Id,
		}, nil
	case *rts.Subject_Set:
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

func (s *SubjectSet) SubjectID() *string {
	return nil
}

func (s *SubjectSet) SubjectSet() *SubjectSet {
	return s
}

func (s *SubjectID) SubjectID() *string {
	return &s.ID
}

func (s *SubjectID) SubjectSet() *SubjectSet {
	return nil
}

// swagger:ignore
func (s *SubjectID) ToProto() *rts.Subject {
	return &rts.Subject{
		Ref: &rts.Subject_Id{
			Id: s.ID,
		},
	}
}

// swagger:ignore
func (s *SubjectSet) ToProto() *rts.Subject {
	return &rts.Subject{
		Ref: &rts.Subject_Set{
			Set: &rts.SubjectSet{
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
	return json.Marshal(s.ID)
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
	var rq RelationQuery
	if err := json.Unmarshal(raw, &rq); err != nil {
		return errors.WithStack(err)
	}
	if rq.SubjectID != nil && rq.SubjectSet != nil {
		return errors.WithStack(ErrDuplicateSubject)
	} else if rq.SubjectID == nil && rq.SubjectSet == nil {
		return errors.WithStack(ErrNilSubject)
	}

	r.Namespace = rq.Namespace
	r.Object = rq.Object
	r.Relation = rq.Relation

	// validation was done before already
	if rq.SubjectID == nil {
		r.Subject = rq.SubjectSet
	} else {
		r.Subject = &SubjectID{ID: *rq.SubjectID}
	}

	return nil
}

func (r *InternalRelationTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.ToQuery())
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

func (r *InternalRelationTuple) ToProto() *rts.RelationTuple {
	return &rts.RelationTuple{
		Namespace: r.Namespace,
		Object:    r.Object,
		Relation:  r.Relation,
		Subject:   r.Subject.ToProto(),
	}
}

func (r *InternalRelationTuple) ToQuery() *RelationQuery {
	return &RelationQuery{
		Namespace:  r.Namespace,
		Object:     r.Object,
		Relation:   r.Relation,
		SubjectID:  r.Subject.SubjectID(),
		SubjectSet: r.Subject.SubjectSet(),
	}
}

func (r *InternalRelationTuple) FromURLQuery(query url.Values) (*InternalRelationTuple, error) {
	q, err := (&RelationQuery{}).FromURLQuery(query)
	if err != nil {
		return nil, err
	}

	if s := q.Subject(); s == nil {
		return nil, errors.WithStack(ErrNilSubject)
	} else {
		r.Subject = s
	}

	r.Namespace = q.Namespace
	r.Object = q.Object
	r.Relation = q.Relation

	return r, nil
}

func (r *InternalRelationTuple) ToURLQuery() (url.Values, error) {
	vals := url.Values{
		"namespace": []string{r.Namespace},
		"object":    []string{r.Object},
		"relation":  []string{r.Relation},
	}
	switch s := r.Subject.(type) {
	case *SubjectID:
		vals.Set(subjectIDKey, s.ID)
	case *SubjectSet:
		vals.Set(subjectSetNamespaceKey, s.Namespace)
		vals.Set(subjectSetObjectKey, s.Object)
		vals.Set(subjectSetRelationKey, s.Relation)
	case nil:
		return nil, errors.WithStack(ErrNilSubject)
	}
	return vals, nil
}

func (r *InternalRelationTuple) ToLoggerFields() logrus.Fields {
	return logrus.Fields{
		"namespace": r.Namespace,
		"object":    r.Object,
		"relation":  r.Relation,
		"subject":   r.Subject.String(),
	}
}

func (q *RelationQuery) FromProto(query TupleData) (*RelationQuery, error) {
	q.Namespace = query.GetNamespace()
	q.Object = query.GetObject()
	q.Relation = query.GetRelation()
	// reset subject
	q.SubjectID = nil
	q.SubjectSet = nil

	if query.GetSubject() != nil {
		switch s := query.GetSubject().Ref.(type) {
		case *rts.Subject_Id:
			q.SubjectID = &s.Id
		case *rts.Subject_Set:
			q.SubjectSet = &SubjectSet{
				Namespace: s.Set.Namespace,
				Object:    s.Set.Object,
				Relation:  s.Set.Relation,
			}
		case nil:
			return nil, errors.WithStack(ErrNilSubject)
		}
	}

	return q, nil
}

func (q *RelationQuery) UUIDMappableFields() []*string {
	res := []*string{&q.Object}
	if q.SubjectID != nil {
		res = append(res, q.SubjectID)
	}
	if q.SubjectSet != nil {
		res = append(res, q.SubjectSet.UUIDMappableFields()...)
	}
	return res
}

const (
	subjectIDKey           = "subject_id"
	subjectSetNamespaceKey = "subject_set.namespace"
	subjectSetObjectKey    = "subject_set.object"
	subjectSetRelationKey  = "subject_set.relation"
)

func (q *RelationQuery) FromURLQuery(query url.Values) (*RelationQuery, error) {
	if q == nil {
		q = &RelationQuery{}
	}

	if query.Has("subject") {
		return nil, errors.WithStack(ErrDroppedSubjectKey)
	}

	// reset subject
	q.SubjectID = nil
	q.SubjectSet = nil

	switch {
	case !query.Has(subjectIDKey) && !query.Has(subjectSetNamespaceKey) && !query.Has(subjectSetObjectKey) && !query.Has(subjectSetRelationKey):
		// was not queried for the subject
	case query.Has(subjectIDKey) && query.Has(subjectSetNamespaceKey) && query.Has(subjectSetObjectKey) && query.Has(subjectSetRelationKey):
		return nil, errors.WithStack(ErrDuplicateSubject)
	case query.Has(subjectIDKey):
		q.SubjectID = pointerx.String(query.Get(subjectIDKey))
	case query.Has(subjectSetNamespaceKey) && query.Has(subjectSetObjectKey) && query.Has(subjectSetRelationKey):
		q.SubjectSet = &SubjectSet{
			Namespace: query.Get(subjectSetNamespaceKey),
			Object:    query.Get(subjectSetObjectKey),
			Relation:  query.Get(subjectSetRelationKey),
		}
	default:
		return nil, errors.WithStack(ErrIncompleteSubject)
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
	if q.SubjectID != nil {
		v.Add(subjectIDKey, *q.SubjectID)
	} else if q.SubjectSet != nil {
		v.Add(subjectSetNamespaceKey, q.SubjectSet.Namespace)
		v.Add(subjectSetObjectKey, q.SubjectSet.Object)
		v.Add(subjectSetRelationKey, q.SubjectSet.Relation)
	}

	return v
}

func (q *RelationQuery) Subject() Subject {
	if q.SubjectID != nil {
		return &SubjectID{ID: *q.SubjectID}
	} else if q.SubjectSet != nil {
		return q.SubjectSet
	}
	return nil
}

func (q *RelationQuery) String() string {
	if q.SubjectID != nil {
		return fmt.Sprintf("namespace: %s; object: %s; relation: %s; subject: %s", q.Namespace, q.Object, q.Relation, *q.SubjectID)
	}
	return fmt.Sprintf("namespace: %s; object: %s; relation: %s; subject: %v", q.Namespace, q.Object, q.Relation, q.SubjectSet)
}

func (r *InternalRelationTuple) Header() []string {
	return []string{
		"NAMESPACE",
		"OBJECT ID",
		"RELATION NAME",
		"SUBJECT",
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

func NewProtoRelationCollection(rels []*rts.RelationTuple) *RelationCollection {
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

func (t *ManagerWrapper) DeleteAllRelationTuples(ctx context.Context, query *RelationQuery) error {
	return t.Reg.RelationTupleManager().DeleteAllRelationTuples(ctx, query)
}

func (t *ManagerWrapper) TransactRelationTuples(ctx context.Context, insert []*InternalRelationTuple, delete []*InternalRelationTuple) error {
	return t.Reg.RelationTupleManager().TransactRelationTuples(ctx, insert, delete)
}

func (t *ManagerWrapper) RelationTupleManager() Manager {
	return t
}

func (t *ManagerWrapper) UUIDMappingManager() UUIDMappingManager {
	return t.Reg.UUIDMappingManager()
}
