package driver

import "github.com/ory/keto/internal/driver/configuration"

type Driver interface {
	Configuration() configuration.Provider
	Registry() Registry
}
