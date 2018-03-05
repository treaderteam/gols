package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type I struct {
	S      *mgo.Session
	DBName string
}

func (i *I) Insert(c CollectingModificator) (err error) {
	c.Modificate()

	sess := i.S.Copy()
	defer sess.Close()

	err = sess.DB(i.DBName).C(c.GetCollectionName()).Insert(c)

	return
}

func (i *I) GetOne(id string, c CollectingModeller) (err error) {
	sess := i.S.Copy()
	defer sess.Close()

	err = sess.DB(i.DBName).C(c.GetCollectionName()).Find(bson.M{"_id": id}).One(c.GetModel())

	return
}

func (i *I) GetOneArbitrary(query map[string]string, c CollectingModeller) (err error) {
	sess := i.S.Copy()
	defer sess.Close()
	b := bson.M{}
	for k, v := range query {
		b[k] = v
	}

	err = sess.DB(i.DBName).C(c.GetCollectionName()).Find(b).One(c.GetModel())

	return
}

func (i *I) GetAll(c interface{}, modelname string) (err error) {

	sess := i.S.Copy()
	defer sess.Close()

	err = sess.DB(i.DBName).C(modelname).Find(bson.M{}).All(c)

	return
}

func (i *I) Update(id string, c CollectingModeller) (err error) {
	sess := i.S.Copy()
	defer sess.Close()

	err = sess.DB(i.DBName).C(c.GetCollectionName()).Update(bson.M{"_id": id}, c.GetModel())

	return
}

func (i *I) UpdateArbitrary(query map[string]string, c CollectingModeller) (err error) {
	sess := i.S.Copy()
	defer sess.Close()
	b := bson.M{}
	for k, v := range query {
		b[k] = v
	}

	err = sess.DB(i.DBName).C(c.GetCollectionName()).Update(b, c.GetModel())

	return
}

func (i *I) DeleteArbitrary(query map[string]string, c CollectingModeller) (err error) {
	sess := i.S.Copy()
	defer sess.Close()
	b := bson.M{}
	for k, v := range query {
		b[k] = v
	}

	err = sess.DB(i.DBName).C(c.GetCollectionName()).Remove(b)

	return
}
