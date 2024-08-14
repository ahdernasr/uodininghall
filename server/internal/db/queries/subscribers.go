package queries

import (
	"fmt"

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
	// Query to check if the email already exists
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM Subscribers WHERE email=$1)`
	err := db.DB.QueryRow(checkQuery, email).Scan(&exists)
	if err != nil {
		return err // Return error if query fails
	}

	// If email already exists, return an error
	if exists {
		return fmt.Errorf("email %s already exists in the mailing list", email)
	}

	// If email does not exist, proceed to insert
	insertQuery := `INSERT INTO Subscribers (email) values ($1)`
	_, err = db.DB.Exec(insertQuery, email)
	if err != nil {
		return err // Return error if insertion fails
	}

	return nil
}

func RemoveSubscriber(email string) error {
	statement := `DELETE FROM Subscribers WHERE email=$1`
	_, err := db.DB.Exec(statement, email)
	if err != nil {
		panic(err)
	}
	return nil
}
