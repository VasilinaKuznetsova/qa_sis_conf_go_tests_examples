package integration_tests

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	DSN    = "user=postgres password=postgres host=localhost port=6432 dbname=users sslmode=disable"
	DRIVER = "postgres"
)

func TestInsertUser(t *testing.T) {
	db, err := sql.Open(DRIVER, DSN)
	require.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT NOT NULL
        );
    `)
	require.NoError(t, err)

	user := User{Name: "John Doe", Email: "johndoe@example.com"}
	id, err := InsertUser(db, user)
	require.NoError(t, err)

	row := db.QueryRow("SELECT name, email FROM users WHERE id = ?", id)
	var name, email string
	err = row.Scan(&name, &email)
	require.NoError(t, err)

	require.Equal(t, "John Doe", name)
	require.Equal(t, "johndoe@example.com", email)
}
