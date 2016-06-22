package main

import (
	"github.com/kobeld/goutils"
	"github.com/kobeld/mgowrap"
	"gopkg.in/mgo.v2/bson"
)

type Shop struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Location Location
}

func (this *Shop) CollectionName() string {
	return "shops"
}

func (this *Shop) MakeId() interface{} {
	if !this.Id.Valid() {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

type Location struct {
	Type        string
	Coordinates []float64
}

func main() {
	mgowrap.SetupDatbase("localhost", "geo")

	var (
		err   error
		shops []*Shop

		long  = 139.701642
		lat   = 35.690647
		scope = 3000

		query = bson.M{
			"location": bson.M{"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
				"$maxDistance": scope,
			}},
		}
	)

	// err = mgowrap.FindAll(query, &shops)

	err = mgowrap.FindWithLimit(query, &shops, 1)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	for _, shop := range shops {
		println(shop.Name)
	}

}
