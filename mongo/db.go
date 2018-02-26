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
		ConnectNoAuth(url, dbName)
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
func ConnectNoAuth(url, dbName string) (sess *mgo.Session, err error) {
	if sess, err = mgo.Dial(url); err != nil {
		return nil, err
	}
	if err = sess.DB(dbName).C("users").EnsureIndexKey("email"); err != nil {
		return nil, err
	}
	return
}

// ConnectN new connection
func ConnectN(url, dbName string) (result I, err error) {
	var sess *mgo.Session
	if sess, err = mgo.Dial(url); err != nil {
		return
	}

	result = I{
		DBName: dbName,
		S:      sess,
	}

	return
}

// ConnectNAuth new secure connection
func ConnectNAuth(url, dbName, login, pwd string) (result I, err error) {
	if result, err = ConnectN(url, dbName); err != nil {
		return
	}

	fmt.Printf("login with %s %s, %s\n", url, login, pwd)

	if err = result.S.Login(&mgo.Credential{Username: login, Password: pwd, Source: dbName}); err != nil {
		return
	}

	return
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
