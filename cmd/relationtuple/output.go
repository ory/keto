package relationtuple

import (
	"encoding/json"
	"fmt"
	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	Collection struct {
		protoRelations    []*rts.RelationTuple
		internalRelations []*ketoapi.RelationTuple
	}
	OutputTuple struct {
		*ketoapi.RelationTuple
	}
)

func NewProtoCollection(rels []*rts.RelationTuple) *Collection {
	return &Collection{
		protoRelations: rels,
	}
}

func NewAPICollection(rels []*ketoapi.RelationTuple) *Collection {
	return &Collection{
		internalRelations: rels,
	}
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
	ir, err := r.Normalize()
	if err != nil {
		return [][]string{{fmt.Sprintf("%+v", err)}}
	}

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

func (r *Collection) Normalize() ([]*ketoapi.RelationTuple, error) {
	if r.internalRelations == nil {
		r.internalRelations = make([]*ketoapi.RelationTuple, len(r.protoRelations))
		for i, rel := range r.protoRelations {
			ir, err := (&ketoapi.RelationTuple{}).FromDataProvider(rel)
			if err != nil {
				return nil, err
			}
			r.internalRelations[i] = ir
		}
	}
	return r.internalRelations, nil
}

func (r *Collection) Interface() interface{} {
	ir, err := r.Normalize()
	if err != nil {
		return err
	}
	return ir
}

func (r *Collection) MarshalJSON() ([]byte, error) {
	ir, err := r.Normalize()
	if err != nil {
		return nil, err
	}
	return json.Marshal(ir)
}

func (r *Collection) UnmarshalJSON(raw []byte) error {
	return json.Unmarshal(raw, &r.internalRelations)
}

func (r *Collection) Len() int {
	if ir := len(r.internalRelations); ir > 0 {
		return ir
	}
	return len(r.protoRelations)
}

func (r *Collection) IDs() []string {
	ts, err := r.Normalize()
	if err != nil {
		// fmt.Sprintf to include the stacktrace
		return []string{fmt.Sprintf("%+v", err)}
	}
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
