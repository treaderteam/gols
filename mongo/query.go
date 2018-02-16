package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

// GetOne querying model from database by specified ID
func GetOne(id string, c CollectingModeller) (err error) {
	sess := session.Copy()
	defer sess.Close()

	err = sess.DB(dbName).C(c.GetCollectionName()).FindId(id).One(c.GetModel())

	return
}

// GetOneArbitrary same as GetOne, but allow query with arbitrary data
func GetOneArbitrary(query map[string]string, c CollectingModeller) (err error) {
	sess := session.Copy()
	defer sess.Close()
	b := bson.M{}
	for k, v := range query {
		b[k] = v
	}

	err = sess.DB(dbName).C(c.GetCollectionName()).Find(b).One(c.GetModel())

	return
}
