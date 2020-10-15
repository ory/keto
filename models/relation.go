package models

import "google.golang.org/protobuf/runtime/protoimpl"

type Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	ObjectID string `json:"object_id"`
}

var _ = Relation(GRPCRelation{})
var _ = GRPCRelation(Relation{})

func (r *Relation) ImportFromGRPC(gr *GRPCRelation) *Relation {
	//goland:noinspection GoVetCopyLock - state is reset afterwards
	*r = Relation(*gr)
	r.state = protoimpl.MessageState{}
	return r
}

func (x *GRPCRelation) ImportFromNormal(r *Relation) *GRPCRelation {
	//goland:noinspection GoVetCopyLock - state is reset afterwards
	*x = GRPCRelation(*r)
	x.state = protoimpl.MessageState{}
	return x
}
