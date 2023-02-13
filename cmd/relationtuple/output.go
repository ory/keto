// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"encoding/json"

	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	Collection struct {
		apiRelations []*ketoapi.RelationTuple
	}
	OutputTuple struct {
		*ketoapi.RelationTuple
	}
)

func NewProtoCollection(rels []*rts.RelationTuple) (*Collection, error) {
	r := &Collection{apiRelations: make([]*ketoapi.RelationTuple, len(rels))}
	for i, rel := range rels {
		var err error
		r.apiRelations[i], err = (&ketoapi.RelationTuple{}).FromDataProvider(rel)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func MustNewProtoCollection(rels []*rts.RelationTuple) *Collection {
	c, err := NewProtoCollection(rels)
	if err != nil {
		panic(err)
	}
	return c
}

func NewAPICollection(rels []*ketoapi.RelationTuple) *Collection {
	return &Collection{apiRelations: rels}
}

func (r *Collection) Header() []string {
	return []string{
		"NAMESPACE",
		"OBJECT",
		"RELATION NAME",
		"SUBJECT",
	}
}

func (r *Collection) Table() [][]string {
	ir := r.apiRelations

	data := make([][]string, len(ir))
	for i, rel := range ir {
		var sub string
		if rel.SubjectID != nil {
			sub = *rel.SubjectID
		} else {
			sub = rel.SubjectSet.String()
		}

		data[i] = []string{rel.Namespace, rel.Object, rel.Relation, sub}
	}

	return data
}

func (r *Collection) Interface() interface{} {
	return r.apiRelations
}

func (r *Collection) MarshalJSON() ([]byte, error) {
	ir := r.apiRelations
	return json.Marshal(ir)
}

func (r *Collection) UnmarshalJSON(raw []byte) error {
	return json.Unmarshal(raw, &r.apiRelations)
}

func (r *Collection) Len() int {
	return len(r.apiRelations)
}

func (r *Collection) IDs() []string {
	ts := r.apiRelations
	ids := make([]string, len(ts))
	for i, rt := range ts {
		ids[i] = rt.String()
	}
	return ids
}

func (r *OutputTuple) Header() []string {
	return []string{
		"NAMESPACE",
		"OBJECT ID",
		"RELATION NAME",
		"SUBJECT",
	}
}

func (r *OutputTuple) Columns() []string {
	return []string{
		r.Namespace,
		r.Object,
		r.Relation,
		outputSubject(r.RelationTuple),
	}
}

func outputSubject(r *ketoapi.RelationTuple) string {
	if r.SubjectID != nil {
		return *r.SubjectID
	}
	return r.SubjectSet.String()
}
