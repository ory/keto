package expand

import (
	"context"
	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/x"

	"github.com/ory/keto/internal/relationtuple"
)

type (
	EngineDependencies interface {
		relationtuple.ManagerProvider
		namespace.ManagerProvider
	}
	Engine struct {
		d EngineDependencies
	}
	EngineProvider interface {
		ExpandEngine() *Engine
	}
)

func NewEngine(d EngineDependencies) *Engine {
	return &Engine{d: d}
}

func (e *Engine) BuildTree(ctx context.Context, subject relationtuple.Subject, restDepth int) (*Tree, error) {
	if restDepth <= 0 {
		return nil, nil
	}

	if us, isUserSet := subject.(*relationtuple.SubjectSet); isUserSet {
		nm, err := e.d.NamespaceManager()
		n, err := nm.GetNamespace(ctx, us.Namespace)

		subTree := &Tree{
			Type:    n.GetRelationRewriteOperator(ctx, us.Relation),
			Subject: subject,
		}

		var (
			rels     []*relationtuple.InternalRelationTuple
			nextPage string
		)
		for nextPage != x.PageTokenEnd {
			var err error
			rels, nextPage, err = e.d.RelationTupleManager().GetRelationTuples(
				ctx,
				&relationtuple.RelationQuery{
					Relation:  us.Relation,
					Object:    us.Object,
					Namespace: us.Namespace,
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
