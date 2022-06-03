package relationtuple

import "github.com/ory/keto/ketoapi"

var (
	_ = (*relationTupleWithRequired)(nil)
	_ = (*patchPayload)(nil)
)

// swagger:model InternalRelationTuple
type relationTupleWithRequired struct {
	// Namespace of the Relation Tuple
	//
	// required: true
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	//
	// required: true
	Object string `json:"object"`

	// Relation of the Relation Tuple
	//
	// required: true
	Relation string `json:"relation"`

	// SubjectID of the Relation Tuple
	//
	// Either SubjectSet or SubjectID are required.
	SubjectID *string `json:"subject_id,omitempty"`
	// SubjectSet of the Relation Tuple
	//
	// Either SubjectSet or SubjectID are required.
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
}

// The patch request payload
//
// swagger:parameters patchRelationTuples
type patchPayload struct {
	// in:body
	Payload []*ketoapi.PatchDelta
}
