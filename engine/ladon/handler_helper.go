package ladon

import (
	"github.com/go-errors/errors"
	"github.com/pborman/uuid"
)

func validatePolicy(p Policy) (Policy, error) {
	if len(p.ID) == 0 {
		p.ID = uuid.New()
	}

	if p.Effect != "allow" && p.Effect != "deny" {
		return Policy{}, errors.Errorf("invalid policy effect %s, only allow and deny are supported", p.Effect)
	}

	return p, nil
}
