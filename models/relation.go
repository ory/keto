package models

import (
	"google.golang.org/protobuf/runtime/protoimpl"

	"github.com/ory/x/cmdx"
)

type (
	Relation struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields

		UserID   string `json:"user_id"`
		Name     string `json:"name"`
		ObjectID string `json:"object_id"`
	}
	relationCollection struct {
		grpcRelations     []*GRPCRelation
		internalRelations []*Relation
	}
)

var _ = Relation(GRPCRelation{})
var _ = GRPCRelation(Relation{})
var _ cmdx.OutputEntry = &Relation{}

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

func (r *Relation) Header() []string {
	return []string{
		"RELATION NAME",
		"USER ID",
		"OBJECT ID",
	}
}

func (r *Relation) Fields() []string {
	return []string{
		r.Name,
		r.UserID,
		r.ObjectID,
	}
}

func (r *Relation) Interface() interface{} {
	return r
}

func NewGRPCRelationCollection(rels []*GRPCRelation) cmdx.OutputCollection {
	return &relationCollection{
		grpcRelations: rels,
	}
}

func NewRelationCollection(rels []*Relation) cmdx.OutputCollection {
	return &relationCollection{
		internalRelations: rels,
	}
}

func (r *relationCollection) Header() []string {
	return []string{
		"RELATION NAME",
		"USER ID",
		"OBJECT ID",
	}
}

func (r *relationCollection) Table() [][]string {
	if r.internalRelations == nil {
		for _, rel := range r.grpcRelations {
			r.internalRelations = append(r.internalRelations, (*Relation)(rel))
		}
	}

	data := make([][]string, len(r.internalRelations))
	for i, rel := range r.internalRelations {
		data[i] = []string{rel.Name, rel.UserID, rel.ObjectID}
	}

	return data
}

func (r *relationCollection) Interface() interface{} {
	if r.internalRelations == nil {
		return r.grpcRelations
	}
	return r.internalRelations
}

func (r relationCollection) Len() int {
	// one of them is zero so the sum is always correct
	return len(r.grpcRelations) + len(r.internalRelations)
}
