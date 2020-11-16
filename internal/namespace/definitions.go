package namespace

import "context"

type (
	Namespace struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}
	Manager interface {
		NewNamespace(ctx context.Context, n *Namespace) error
	}
	ManagerProvider interface {
		NamespaceManagerProvider() Manager
	}
)

func (n *Namespace) TableName() string {
	return "keto_namespace"
}
