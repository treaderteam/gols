package postg_test

import (
	"fmt"
	"log"
	"testing"

	"gitlab.com/alexnikita/gols/postg"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func TestConnect(t *testing.T) {
	connstr := "user=test password='test' dbname=test port=5433 sslmode=disable"

	db, err := postg.Connect(connstr)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
}

func TestCreateTable(t *testing.T) {
	connstr := "user=test password='test' dbname=test port=5433 sslmode=disable"

	db, err := postg.Connect(connstr)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}

	tablename := "users"
	table := `
	id SERIAL PRIMARY KEY,
	name varchar(80),
	city int
	`

	if err = postg.EnsureTable(db, tablename, table); err != nil {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	connstr := "user=test password='test' dbname=test port=5433 sslmode=disable"

	db, err := postg.Connect(connstr)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}

	tablename := "users"
	table := `
	id SERIAL PRIMARY KEY,
	name varchar(80),
	city int
	`

	if err = postg.EnsureTable(db, tablename, table); err != nil {
		t.Fatal(err)
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (name, city) VALUES ('%s', %d);
	`, tablename, "Nikita", 1)

	if _, err := db.Exec(query); err != nil {
		t.Fatal(err)
	}

	selectquery := fmt.Sprintf(`
		SELECT name, city FROM users;
	`)

	rows, err := db.Query(selectquery)
	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		name := ""
		city := 0
		if err = rows.Scan(&name, &city); err != nil {
			t.Fatal(err)
		}

		log.Println(name, city)
	}

}
