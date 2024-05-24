package db_in_memory

import (
	"MyBalance/internal/context"
	"MyBalance/internal/core/db/interface_mono_db"
	"MyBalance/internal/http/requesto"
	"sync"
)

// database - string to string map. nothing hard

var _ interface_mono_db.DBMono = (*databaseMonoInMemory)(nil)

type databaseMonoInMemory struct {
	db sync.Map
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

func (r *databaseMonoInMemory) SetCard(ctx context.Context, user string, card string) {
	r.db.Store(user, card)
}

func NewDatabaseMonoInMemory() interface_mono_db.DBMono {
	return &databaseMonoInMemory{}
}
