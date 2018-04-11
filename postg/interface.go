package postg

import (
	"database/sql"
	"fmt"

	// Inport postgresql driver
	_ "github.com/lib/pq"
)

// Connect makes connection to database with provided
// connection string
func Connect(connstr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// EnsureTable creates table in its not exist
func EnsureTable(db *sql.DB, tablename, table string) error {
	_, err := db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (%s);
		`, tablename, table))

	return err
}
