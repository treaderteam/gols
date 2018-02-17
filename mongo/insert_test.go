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
	sess, err := mongo.ConnectNoAuth(dbURL, dbName)
	if err != nil {
		t.Fatal(err)
	}
	dbi := mongo.I{
		DBName: dbName,
		S:      sess,
	}

	var testModel testModel
	testModel.Name = "test"

	if err = dbi.Insert(&testModel); err != nil {
		log.Panic(err)
	}

}
