package util

import (
	"gitlab.com/alexnikita/gols/mongo"
	"gopkg.in/mgo.v2"
)

func Cleanup() error {
	connection, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}

	connection.DB(mongo.GetName()).DropDatabase()
	return nil
}
