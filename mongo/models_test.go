package mongo_test

import (
	"gitlab.com/alexnikita/gols/mongo"
)

type testModel struct {
	mongo.SimpleModel
	ID   string `bson:"_id"`
	Name string `bson: "name"`
}

func (t *testModel) GetCollectionName() string {
	return "test-model"
}

func (t *testModel) GetModel() interface{} {
	return t
}
