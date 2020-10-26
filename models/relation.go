package models

import (
	"fmt"

	"github.com/ory/x/cmdx"
)

type (
	relationCollection struct {
		grpcRelations     []*RelationTuple
		internalRelations []*Relation
	}
	Object struct {
		ID        string
		Namespace string
	}
	User interface {
		String() string
	}
	UserID struct {
		User
		ID string
	}
	UserSet struct {
		User
		Object   Object
		Relation string
	}
	Relation struct {
		Object   Object
		Relation string
		User     User
	}
	RelationQuery struct {
		Object   Object
		Relation string
		User     User
	}
)

func (o Object) String() string {
	return fmt.Sprintf("%s:%s", o.Namespace, o.ID)
}

func (u UserID) String() string {
	return fmt.Sprintf("%s", u.ID)
}

func (u UserSet) String() string {
	return fmt.Sprintf("%s#%s", u.Object, u.Relation)
}

func (r Relation) String() string {
	return fmt.Sprintf("%s#%s@%s", r.Object, r.Relation, r.User)
}

func (r *Relation) ImportFromGRPC(gr *RelationTuple) *Relation {
	var user User
	switch gr.User.(type) {
	case *RelationTuple_UserId:
		user = UserID{
			ID: gr.GetUserId(),
		}
	case *RelationTuple_UserSet:
		user = UserSet{
			Object: Object{
				gr.GetUserSet().Object.Namespace,
				gr.GetUserSet().Object.ObjectId,
			},
			Relation: gr.GetUserSet().Relation,
		}
	}

	r.Object = Object{
		gr.Object.Namespace,
		gr.Object.ObjectId,
	}
	r.Relation = gr.Relation
	r.User = user

	return r
}

func (gr *RelationTuple) ImportFromNormal(r *Relation) *RelationTuple {
	var user isRelationTuple_User
	switch r.User.(type) {
	case UserID:
		user = &RelationTuple_UserId{
			r.User.(UserID).ID,
		}
	case UserSet:
		userSet := r.User.(UserSet)
		user = &RelationTuple_UserSet{
			UserSet: &RelationUserSet{
				Object: &RelationObject{
					ObjectId:  userSet.Object.ID,
					Namespace: userSet.Object.Namespace,
				},
			},
		}
	}

	gr.Object = &RelationObject{
		ObjectId:  r.Object.ID,
		Namespace: r.Object.Namespace,
	}
	gr.Relation = r.Relation
	gr.User = user

	return gr
}

func (rq *RelationQuery) ImportFromGRPC(rtq *ReadRelationTuplesRequest_Query) *RelationQuery {
	var user User
	switch rtq.User.(type) {
	case *ReadRelationTuplesRequest_Query_UserId:
		user = UserID{
			ID: rtq.GetUserId(),
		}
	case *ReadRelationTuplesRequest_Query_UserSet:
		user = UserSet{
			Object: Object{
				rtq.GetUserSet().Object.Namespace,
				rtq.GetUserSet().Object.ObjectId,
			},
			Relation: rtq.GetUserSet().Relation,
		}
	}

	rq.Object = Object{
		rtq.Object.Namespace,
		rtq.Object.ObjectId,
	}
	rq.Relation = rtq.Relation
	rq.User = user

	return rq
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
		r.Relation,
		r.User.String(),
		r.Object.String(),
	}
}

func (r *Relation) Interface() interface{} {
	return r
}

func NewGRPCRelationCollection(rels []*RelationTuple) cmdx.OutputCollection {
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
			r.internalRelations = append(r.internalRelations, (&Relation{}).ImportFromGRPC(rel))
		}
	}

	data := make([][]string, len(r.internalRelations))
	for i, rel := range r.internalRelations {
		data[i] = []string{rel.Relation, rel.User.String(), rel.Object.String()}
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
