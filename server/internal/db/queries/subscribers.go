package queries

import (
	"github.com/ahdernasr/dailydininghall/internal/db"
	_ "github.com/lib/pq"
)

type Subscriber struct {
	Email string
}

func GetAllSubscribers() ([]Subscriber, error) {
	statement := `SELECT email FROM subscribers`
	rows, err := db.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscribers []Subscriber
	for rows.Next() {
		var subscriber Subscriber
		if err := rows.Scan(&subscriber.Email); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, subscriber)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
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
