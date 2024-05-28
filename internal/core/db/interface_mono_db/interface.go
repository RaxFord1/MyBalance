package interface_mono_db

import (
	"MyBalance/internal/http/context"
)

// DBMono - interface for mono pkg
type DBMono interface {
	GetCard(context.Context, string) (string, error) // GET card code for user
	SetCard(context.Context, string, string) error   // SET card code for user
	GetUsers(context.Context) []string               // Get all users
}
