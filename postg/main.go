package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres sslmode=disable dbname=postgres port=5432"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS weather (
			city varchar(80),
			temp_lo int,
			temp_hi int,
			prcp real,
			date date
		)	
	`)
	if err != nil {
		log.Fatal(err)
	}

	qresult, err := db.Query(`
		SELECT * FROM weather WHERE temp_lo = $1;
	`, -2)
	if err != nil {
		log.Fatal(err)
	}

	defer qresult.Close()

	for qresult.Next() {
		var (
			city   string
			tempLo int
			tempHi int
			prcp   int
			date   time.Time
		)

		_ = qresult.Scan(&city, &tempLo, &tempHi, &prcp, &date)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		log.Println(city, tempLo, tempHi, date)
	}
}
