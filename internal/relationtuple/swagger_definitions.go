// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

var (
	_ = (*patchRelationships)(nil)
	_ = (*getRelationships)(nil)
	_ = (*relationshipInQuery)(nil)
)

// Patch Relationships Request Parameters
//
// swagger:parameters patchRelationships
type patchRelationships struct {
	// in:body
	Body []*ketoapi.PatchDelta
}

// Get Relationships Request Parameters
//
// swagger:parameters getRelationships
type getRelationships struct {
	// Namespace of the Relationship
	//
	// in: query
	Namespace string `json:"namespace"`

	// Object of the Relationship
	//
	// in: query
	Object string `json:"object"`

	// Relation of the Relationship
	//
	// in: query
	Relation string `json:"relation"`

	// SubjectID of the Relationship
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

// The relationship parameters in the URL query.
//
// swagger:parameters checkPermission checkPermissionOrError deleteRelationships
type relationshipInQuery struct {
	// Namespace of the Relationship
	//
	// in: query
	Namespace string `json:"namespace"`

	// Object of the Relationship
	//
	// in: query
	Object string `json:"object"`

	// Relation of the Relationship
	//
	// in: query
	Relation string `json:"relation"`

	// SubjectID of the Relationship
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
