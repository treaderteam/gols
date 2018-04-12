package postg

import (
	"database/sql"
	"fmt"
)

// Clean cleand specified table
func Clean(table string, db *sql.DB) error {
	deletetablequery := fmt.Sprintf(`
		DELETE FROM %s;	
	`, table)

	_, err := db.Exec(deletetablequery)

	return err
}
