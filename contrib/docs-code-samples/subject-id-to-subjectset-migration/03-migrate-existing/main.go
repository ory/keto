// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
)

const batchSize = 100

// subjectIDToNamespacedSubject converts a prefixed subject ID to a namespace and object.
// The prefix encodes the type: "user_" → User namespace, "apikey_" → ApiKey namespace.
func subjectIDToNamespacedSubject(id string) (namespace, object string) {
	switch {
	case strings.HasPrefix(id, "user_"):
		return "User", strings.TrimPrefix(id, "user_")
	case strings.HasPrefix(id, "apikey_"):
		return "ApiKey", strings.TrimPrefix(id, "apikey_")
	default:
		panic("unknown subject ID prefix: " + id)
	}
}

func main() {
	rc, err := grpc.NewClient("127.0.0.1:4466", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	wc, err := grpc.NewClient("127.0.0.1:4467", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	readClient := rts.NewReadServiceClient(rc)
	writeClient := rts.NewWriteServiceClient(wc)
	ctx := context.Background()

	// Step 1: Paginate all File#viewer tuples and collect those with a SubjectID.
	// In production, persist these to a file, SQLite, or a database before writing back,
	// so the migration can be resumed if interrupted.
	var toMigrate []*rts.RelationTuple
	var pageToken string
	for {
		resp, err := readClient.ListRelationTuples(ctx, &rts.ListRelationTuplesRequest{
			RelationQuery: &rts.RelationQuery{
				Namespace: new("File"),
				Relation:  new("viewer"),
			},
			PageToken: pageToken,
			PageSize:  batchSize,
		})
		if err != nil {
			panic(err)
		}
		for _, t := range resp.RelationTuples {
			if t.Subject.GetId() != "" {
				toMigrate = append(toMigrate, t)
			}
		}
		if resp.NextPageToken == "" {
			break
		}
		pageToken = resp.NextPageToken
	}
	// Step 2: Write SubjectSet counterparts in batches.
	for i := 0; i < len(toMigrate); i += batchSize {
		end := min(i+batchSize, len(toMigrate))
		batch := toMigrate[i:end]

		deltas := make([]*rts.RelationTupleDelta, len(batch))
		for j, t := range batch {
			ns, obj := subjectIDToNamespacedSubject(t.Subject.GetId())
			deltas[j] = &rts.RelationTupleDelta{
				Action: rts.RelationTupleDelta_ACTION_INSERT,
				RelationTuple: &rts.RelationTuple{
					Namespace: t.Namespace,
					Object:    t.Object,
					Relation:  t.Relation,
					Subject:   rts.NewSubjectSet(ns, obj, ""),
				},
			}
		}

		if _, err := writeClient.TransactRelationTuples(ctx, &rts.TransactRelationTuplesRequest{
			RelationTupleDeltas: deltas,
		}); err != nil {
			panic(err)
		}
	}

	fmt.Println("Migration complete")
}
