package driver

import (
	"github.com/ory/keto/internal/driver/config"
)

type Driver interface {
	Configuration() config.Provider
	Registry() Registry
}
