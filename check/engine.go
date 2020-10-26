package check

import (
	"context"
	"fmt"
	"os"

	"github.com/ory/keto/models"
	"github.com/ory/keto/relation"
)

type (
	Engine struct {
		d dependencies
	}
	dependencies interface {
		relation.ManagerProvider
	}
)

func NewEngine(d dependencies) *Engine {
	return &Engine{
		d: d,
	}
}

func equalRelation(a, b *models.Relation) bool {
	return a.SubjectID == b.SubjectID && a.Name == b.Name && a.ObjectID == b.ObjectID
}

func (e *Engine) subjectIsAllowed(ctx context.Context, requested *models.Relation, subjectRelations []*models.Relation) (bool, error) {
	// The question can be rephrased as "is requested.Name on requested.ObjectID reachable from requested.SubjectID"
	//
	// recursive breadth-first search
	// TODO replace by more performant algorithm

	fmt.Printf("%s\n    %s\n", subjectRelations, requested)

	var res bool
	for _, sr := range subjectRelations {

		if requested.Name == sr.Name && requested.ObjectID == sr.ObjectID {
			// found the requested relation
			return true, nil
		}

		prevRelationsLen := len(subjectRelations)

		// compute one indirection
		indirect, err := e.d.RelationManager().GetRelationsBySubject(ctx, sr.ToSubject())
		if err != nil {
			// TODO fix error handling
			_, _ = fmt.Fprintf(os.Stderr, "%+v", err)
		}

		for _, maybeRel := range indirect {
			var found bool
			for _, knownRel := range subjectRelations {
				if equalRelation(knownRel, maybeRel) {
					found = true
					break
				}
			}
			if !found {
				subjectRelations = append(subjectRelations, maybeRel)
			}
		}

		if prevRelationsLen < len(subjectRelations) {
			is, err := e.subjectIsAllowed(ctx, requested, subjectRelations)
			if err != nil {
				// TODO fix error handling
				_, _ = fmt.Fprintf(os.Stderr, "%+v", err)
			}
			res = res || is
		}
	}

	return res, nil
}

func (e *Engine) SubjectIsAllowed(ctx context.Context, r *models.Relation) (bool, error) {
	subjectRelations, err := e.d.RelationManager().GetRelationsBySubject(ctx, r.SubjectID)
	if err != nil {
		return false, err
	}

	return e.subjectIsAllowed(ctx, r, subjectRelations)
}
