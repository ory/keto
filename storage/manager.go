package storage

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/open-policy-agent/opa/storage"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/pkg/errors"
)

type Manager interface {
	Get(ctx context.Context, collection string, key string, value interface{}) error
	List(ctx context.Context, collection string, value interface{}, limit, offset int) error
	Upsert(ctx context.Context, collection string, key string, value interface{}) error
	Delete(ctx context.Context, collection string, key string) error
	Storage(ctx context.Context, schema string, collections []string) (storage.Store, error)
}

func roundTrip(in, out interface{}) error {
	var b bytes.Buffer

	if err := json.NewEncoder(&b).Encode(in); err != nil {
		return errors.WithStack(err)
	}

	dec := json.NewDecoder(&b)
	dec.DisallowUnknownFields()
	if err := dec.Decode(out); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func toRegoStore(ctx context.Context, schema string, collections []string, query func(context.Context, string) ([]json.RawMessage, error)) (storage.Store, error) {
	var s map[string]interface{}
	dec := json.NewDecoder(bytes.NewBufferString(schema))
	dec.UseNumber()
	if err := dec.Decode(&s); err != nil {
		return nil, errors.WithStack(err)
	}

	db := inmem.NewFromObject(s)
	txn, err := db.NewTransaction(ctx, storage.WriteParams)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, c := range collections {
		path, ok := storage.ParsePath(c)
		if !ok {
			return nil, errors.Errorf("unable to parse storage path: %s", c)
		}

		var val []interface{}
		var b bytes.Buffer

		d, err := query(ctx, c)
		if err != nil {
			return nil, err
		}

		if err := json.NewEncoder(&b).Encode(d); err != nil {
			return nil, errors.WithStack(err)
		}

		dec := json.NewDecoder(&b)
		dec.UseNumber()
		if err := dec.Decode(&val); err != nil {
			return nil, errors.WithStack(err)
		}

		if err := db.Write(ctx, txn, storage.AddOp, path, val); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if err := db.Commit(ctx, txn); err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
