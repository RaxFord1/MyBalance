package interface_mono_db

import "MyBalance/internal/context"

// DBMono - interface for mono pkg
type DBMono interface {
	GetCard(context.Context, string) (string, error) // GET card code for user
	SetCard(context.Context, string, string)         // SET card code for user
}
