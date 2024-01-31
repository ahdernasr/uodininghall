package queries

import (
	"fmt"

	"github.com/ahdernasr/dailydininghall/internal/db"
)

type User struct {
	ID    int
	Email string
}

func AddUser(user User) error {
	_, err := db.DB.Exec("INSERT INTO users (name) VALUES ($1)", user.Email)
	if err != nil {
		return fmt.Errorf("add user: %w", err)
	}
	return nil
}
