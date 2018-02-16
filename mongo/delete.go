package mongo

import "gopkg.in/mgo.v2/bson"

// Delete removes model from database
func Delete(query map[string]string, c CollectingModeller) (err error) {
	sess := session.Copy()
	defer sess.Close()
	b := bson.M{}
	for k, v := range query {
		b[k] = v
	}

	err = sess.DB(dbName).C(c.GetCollectionName()).Remove(b)

	return
}
