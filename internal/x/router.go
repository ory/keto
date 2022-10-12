// Copyright © 2022 Ory Corp

package x

import (
	"github.com/julienschmidt/httprouter"
)

type (
	ReadRouter      struct{ *httprouter.Router }
	WriteRouter     struct{ *httprouter.Router }
	OPLSyntaxRouter struct{ *httprouter.Router }
)
