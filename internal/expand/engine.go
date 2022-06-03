package expand

import (
	"context"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/graph"

	"github.com/ory/keto/internal/relationtuple"
)

type (
	EngineDependencies interface {
		relationtuple.ManagerProvider
		config.Provider
		x.LoggerProvider
	}
	Engine struct {
		d EngineDependencies
	}
	EngineProvider interface {
		ExpandEngine() *Engine
	}
)

func NewEngine(d EngineDependencies) *Engine {
	return &Engine{
		d: d,
	}
}

func (e *Engine) BuildTree(ctx context.Context, subject relationtuple.Subject, restDepth int) (*Tree, error) {
	// global max-depth takes precedence when it is the lesser or if the request max-depth is less than or equal to 0
	if globalMaxDepth := e.d.Config(ctx).MaxReadDepth(); restDepth <= 0 || globalMaxDepth < restDepth {
		restDepth = globalMaxDepth
	}

	if us, isUserSet := subject.(*relationtuple.SubjectSet); isUserSet {
		ctx, wasAlreadyVisited := graph.CheckAndAddVisited(ctx, subject)
		if wasAlreadyVisited {
			return nil, nil
		}

		subTree := &Tree{
			Type:    Union,
			Subject: subject,
		}

		var (
			rels     []*relationtuple.InternalRelationTuple
			nextPage string
		)
		// do ... while nextPage != ""
		for ok := true; ok; ok = nextPage != "" {
			var err error
			rels, nextPage, err = e.d.RelationTupleManager().GetRelationTuples(
				ctx,
				&relationtuple.RelationQuery{
					Relation:  &us.Relation,
					Object:    &us.Object,
					Namespace: &us.Namespace,
				},
				x.WithToken(nextPage),
			)
			if err != nil {
				return nil, err
			} else if len(rels) == 0 {
				return nil, nil
			}

			if restDepth <= 1 {
				subTree.Type = Leaf
				return subTree, nil
			}

			children := make([]*Tree, len(rels))
			for ri, r := range rels {
				child, err := e.BuildTree(ctx, r.Subject, restDepth-1)
				if err != nil {
					return nil, err
				}
				if child == nil {
					child = &Tree{
						Type:    Leaf,
						Subject: r.Subject,
					}
				}
				children[ri] = child
			}
			subTree.Children = append(subTree.Children, children...)
		}

		return subTree, nil
	}

	// is SubjectID
	return &Tree{
		Type:    Leaf,
		Subject: subject,
	}, nil
}
