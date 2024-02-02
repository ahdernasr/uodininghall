package queries

import (
	"github.com/ahdernasr/dailydininghall/internal/db"
	_ "github.com/lib/pq"
)

type Subscriber struct {
	Email string
}

func AddSubscriber(email string) error {
	statement := `INSERT INTO Subscribers (email) values ($1)`
	_, err := db.DB.Exec(statement, email)
	if err != nil {
		panic(err)
	}
	return nil
}

func RemoveSubscriber(email string) error {
	statement := `REMOVE FROM Subscribers WHERE email=$1`
	_, err := db.DB.Exec(statement, email)
	if err != nil {
		panic(err)
	}
	return nil
}
