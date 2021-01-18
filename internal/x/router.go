package x

import (
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
)

type (
	BasicRouter struct {
		*httprouter.Router
	}
	PrivilegedRouter struct {
		*httprouter.Router
	}
	BasicGRPCServer struct {
		*grpc.Server
	}
	PrivilegedGRPCServer struct {
		*grpc.Server
	}
)
