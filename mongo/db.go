// Package mongo responsible for connection and querying to database
package mongo

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

var (
	connectionURL string
	dbName        string
	session       *mgo.Session
)

// Connect perform database connection with specified URL and database name
func Connect(url, name, login, pwd string) {
	var err error
	log.SetFlags(log.Llongfile | log.Ldate)
	log.SetPrefix("DB: ")
	connectionURL = url
	dbName = name
	if session, err = mgo.Dial(connectionURL); err != nil {
		log.Fatal(err)
	}
	if err = session.Login(&mgo.Credential{Username: login, Password: pwd, Source: name}); err != nil {
		ConnectNoAuth(url)
	}
	if err = session.DB(dbName).C("users").EnsureIndexKey("email"); err != nil {
		log.Fatal(err)
	}
	if err = session.DB(dbName).C("sessions").EnsureIndexKey("user_email", "user_id"); err != nil {
		log.Fatal(err)
	}
	log.Printf("database connection to %s established\n", connectionURL)
}

// Connect without authorization
func ConnectNoAuth(url string) {
	var err error
	if session, err = mgo.Dial(connectionURL); err != nil {
		log.Fatal(err)
	}
	if err = session.DB(dbName).C("users").EnsureIndexKey("email"); err != nil {
		log.Fatal(err)
	}
	log.Printf("database connection to %s established\n", connectionURL)
}

// Connects to database and returns session
func PerformConnection(url, login, pwd, dbname string) (s *mgo.Session, e error) {
	if s, e = mgo.Dial(url); e != nil {
		return
	}

	if len(login) < 1 || len(pwd) < 1 {
		return
	}

	fmt.Printf("login with %s %s, %s\n", url, login, pwd)

	if e = s.Login(&mgo.Credential{Username: login, Password: pwd, Source: dbname}); e != nil {
		return
	}

	return
}

// GetName returns database name
func GetName() string {
	return dbName
}
