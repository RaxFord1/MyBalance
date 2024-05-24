package db

import "MyBalance/internal/context"

func GetCard(ctx context.Context, user string) (string, error) {
	return db.GetCard(ctx, user)
}

func SetCard(ctx context.Context, user string, card string) {
	db.SetCard(ctx, user, card)
}
