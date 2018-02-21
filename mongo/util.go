package mongo

import (
	"gopkg.in/mgo.v2"
)

func Cleanup(db *mgo.Database, coll string) (err error) {
	err = db.C(coll).DropCollection()
	return
}
