package mongo_test

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/alexnikita/gols/mongo"
)

func TestGetAll(t *testing.T) {
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

	for i := 0; i < 10; i++ {
		tm.Name = strconv.Itoa(i)
		tm.ID = strconv.Itoa(i)
		if err = dbi.Insert(&tm); err != nil {
			log.Panic(err)
		}
	}

	req := make([]testModel, 0)
	if err = dbi.GetAll(&req, tm.GetCollectionName()); err != nil {
		log.Panic(err)
	}

	assert.Equal(t, 10, len(req))
	for i := 0; i < 10; i++ {
		assert.Equal(t, strconv.Itoa(i), req[i].Name)
	}
}
