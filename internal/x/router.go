package x

import (
	"github.com/julienschmidt/httprouter"
)

type (
	ReadRouter struct {
		*httprouter.Router
	}
	WriteRouter struct {
		*httprouter.Router
	}
	MetricsRouter struct {
		*httprouter.Router
	}
)
