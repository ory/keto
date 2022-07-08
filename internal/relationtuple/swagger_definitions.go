package relationtuple

import (
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

var (
	_ = (*patchPayload)(nil)
	_ = (*getRelationsParams)(nil)
	_ = (*bodyRelationTuple)(nil)
	_ = (*queryRelationTuple)(nil)
)

// The patch request payload
//
// swagger:parameters patchRelationTuples
type patchPayload struct {
	// in:body
	Payload []*ketoapi.PatchDelta
}

// swagger:parameters getRelationTuples
type getRelationsParams struct {
	// Namespace of the Relation Tuple
	//
	// in: query
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	//
	// in: query
	Object string `json:"object"`

	// Relation of the Relation Tuple
	//
	// in: query
	Relation string `json:"relation"`

	// SubjectID of the Relation Tuple
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SubjectID string `json:"subject_id"`

	// Namespace of the Subject Set
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SNamespace string `json:"subject_set.namespace"`

	// Object of the Subject Set
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SObject string `json:"subject_set.object"`

	// Relation of the Subject Set
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SRelation string `json:"subject_set.relation"`

	// swagger:allOf
	x.PaginationOptions
}

// The basic ACL relation tuple
//
// swagger:parameters postCheck createRelationTuple
type bodyRelationTuple struct {
	// in: body
	Payload ketoapi.RelationQuery
}

// The basic ACL relation tuple
//
// swagger:parameters getCheck deleteRelationTuples
type queryRelationTuple struct {
	// Namespace of the Relation Tuple
	//
	// in: query
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	//
	// in: query
	Object string `json:"object"`

	// Relation of the Relation Tuple
	//
	// in: query
	Relation string `json:"relation"`

	// SubjectID of the Relation Tuple
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SubjectID string `json:"subject_id"`

	// Namespace of the Subject Set
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SNamespace string `json:"subject_set.namespace"`

	// Object of the Subject Set
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SObject string `json:"subject_set.object"`

	// Relation of the Subject Set
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SRelation string `json:"subject_set.relation"`
}
