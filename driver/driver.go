package driver

import "github.com/ory/keto/driver/configuration"

type Driver interface {
	Configuration() configuration.Provider
	Registry() Registry
}
