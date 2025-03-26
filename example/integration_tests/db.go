package integration_tests

import (
	"database/sql"
	"log"
)

type User struct {
	Name  string
	Email string
}

func InsertUser(db *sql.DB, user User) (int64, error) {
	query := `
        INSERT INTO users (name, email)
        VALUES ($1, $2)
        RETURNING id;
    `

	var id int64
	err := db.QueryRow(query, user.Name, user.Email).Scan(&id)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return 0, err
	}

	return id, nil
}
