package db_in_memory

import (
	"MyBalance/internal/core/db/interface_mono_db"
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
	"sync"
)

// database - string to string map. nothing hard

var _ interface_mono_db.DBMono = (*databaseMonoInMemory)(nil)

type databaseMonoInMemory struct {
	db sync.Map
}

func (r *databaseMonoInMemory) GetUsers(ctx context.Context) []string {
	keys := []string{}

	// Function to collect keys
	collectKeys := func(key, value interface{}) bool {
		keyStrVal, ok := key.(string)
		if ok {
			keys = append(keys, keyStrVal)
		}
		return true // continue iteration
	}

	r.db.Range(collectKeys)

	return keys
}

func (r *databaseMonoInMemory) GetCard(ctx context.Context, user string) (string, error) {
	card, found := r.db.Load(user)
	if found {
		val, ok := card.(string)
		if ok {
			return val, nil
		}
		return "", requesto.InternalError.New(ctx)
	}
	return "", requesto.UserNotFound.New(ctx)
}

func (r *databaseMonoInMemory) SetCard(ctx context.Context, user string, card string) error {
	r.db.Store(user, card)
	return nil
}

func NewDatabaseMonoInMemory() interface_mono_db.DBMono {
	return &databaseMonoInMemory{}
}
