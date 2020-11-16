package namespace

import "context"

type (
	Namespace struct {
		ID string `db:"name"`
	}
	Manager interface {
		NewNamespace(ctx context.Context, n *Namespace) error
	}
	ManagerProvider interface {
		NamespaceManagerProvider() Manager
	}
)

func NewNamespace(name string) *Namespace {
	return &Namespace{ID: name}
}

func (n *Namespace) TableName() string {
	return "keto_namespaces"
}
