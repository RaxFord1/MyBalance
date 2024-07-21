package core

import (
	"MyBalance/internal/core/db"
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
	"MyBalance/internal/projkeys"
	"encoding/json"
	"fmt"
	"log"
)

// InitDefaultUsers inits default users from memory variables
//to database that WAS INITIALIZED BEFORE
func InitDefaultUsers(ctx context.Context) error {
	defaultUsersRaw, found := ctx.Get(projkeys.DefaultUsers)
	if !found {
		log.Println("Default users not found!")
		return nil
	}

	str, ok := defaultUsersRaw.(string)
	if !ok {
		log.Println("Default users not found!")
		return requesto.InternalError.NewWithMsg(ctx, "Error parsing default users to string")
	}

	// parse to a map[string]string
	var users = make(map[string]string)
	if err := json.Unmarshal([]byte(str), &users); err != nil {
		return requesto.InternalError.NewWithMsg(ctx,
			fmt.Sprintf(
				"Error parsing default users from json %v",
				err.Error(),
			),
		)
	}

	for user, creditCard := range users {
		if err := db.SetCard(ctx, user, creditCard); err != nil {
			return err
		}
	}

	return nil
}
