package mongo_test

import (
	"log"

	"gopkg.in/mgo.v2"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func Cleanup() {
	log.Printf("cleaning up\n")
	connection, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}

	err = connection.DB(dbName).DropDatabase()
	if err != nil {
		log.Println(err)
	}

}
