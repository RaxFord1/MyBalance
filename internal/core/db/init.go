package db

import (
	"MyBalance/internal/context"
	"MyBalance/internal/core/db/db_in_memory"
	"MyBalance/internal/core/db/interface_mono_db"
)

var db interface_mono_db.DBMono

func Init(ctx context.Context) error {
	if db == nil {
		db = db_in_memory.NewDatabaseMonoInMemory()
	}

	return nil
}
