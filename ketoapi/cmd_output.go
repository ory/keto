// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoapi

func (r *RelationTuple) Header() []string {
	return []string{
		"NAMESPACE",
		"OBJECT ID",
		"RELATION NAME",
		"SUBJECT",
	}
}

func (r *RelationTuple) Columns() []string {
	sub := "<!error no subject>"
	switch {
	case r.SubjectID != nil:
		sub = *r.SubjectID
	case r.SubjectSet != nil:
		sub = r.SubjectSet.String()
	}

	return []string{
		r.Namespace,
		r.Object,
		r.Relation,
		sub,
	}
}

func (r *RelationTuple) Interface() interface{} {
	return r
}
