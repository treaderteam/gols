package mongo_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/alexnikita/gols/mongo"
)

func TestGetOne(t *testing.T) {
	defer Cleanup()
	mongo.ConnectNoAuth(dbURL)

	var tm testModel
	var result testModel
	tm.Name = "test-get-one"
	tm.ID = "1"

	if err := mongo.Insert(&tm); err != nil {
		log.Fatal(err)
	}

	if err := mongo.GetOne("1", &result); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, tm.Name, result.Name)

}
