package ladon

import (
	"github.com/go-errors/errors"
	"github.com/pborman/uuid"

	kstorage "github.com/ory/keto/storage"
)

func validatePolicy(p kstorage.Policy) (kstorage.Policy, error) {
	if len(p.ID) == 0 {
		p.ID = uuid.New()
	}

	if p.Effect != "allow" && p.Effect != "deny" {
		return kstorage.Policy{}, errors.Errorf("invalid policy effect %s, only allow and deny are supported", p.Effect)
	}

	return p, nil
}
