package mongo_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/alexnikita/gols/mongo"
)

func TestGetOne(t *testing.T) {
	defer Cleanup()
	sess, err := mongo.ConnectNoAuth(dbURL, dbName)
	if err != nil {
		t.Fatal(err)
	}
	dbi := mongo.I{
		DBName: dbName,
		S:      sess,
	}

	var tm testModel
	var result testModel
	tm.Name = "test-get-one"
	tm.ID = "1"

	if err := dbi.Insert(&tm); err != nil {
		log.Panic(err)
	}

	if err := dbi.GetOne("1", &result); err != nil {
		log.Panic(err)
	}

	assert.Equal(t, tm.Name, result.Name)

}
