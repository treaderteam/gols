package mongo_test

import (
	"log"
	"testing"

	"gitlab.com/alexnikita/gols/mongo"
)

var (
	dbURL  = "localhost"
	dbName = "treader_test"
)

func TestInsert(t *testing.T) {
	defer Cleanup()
	mongo.ConnectNoAuth(dbURL)

	var testModel testModel
	testModel.Name = "test"

	if err := mongo.Insert(&testModel); err != nil {
		log.Panic(err)
	}

}
