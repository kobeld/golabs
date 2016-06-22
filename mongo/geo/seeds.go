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

	names := []string{
		"Shop in Sydney",
		"Studio Alta",
		"ビックロ",
		"Keio Plaza Hotel",
		"明治神宮",
		"Hachiko",
		"Haneda Airport",
	}

	coords := [][]float64{
		{151.20699, -33.867487},
		{139.7011064, 35.692474},
		{139.70328368, 35.69146649},
		{139.69489306, 35.68999278},
		{139.69936833, 35.67612138},
		{139.7005894, 35.65905387},
		{139.78569388, 35.54958159},
	}

	for index, name := range names {
		shop := &Shop{
			Name: name,
			Location: Location{
				Type:        "Point",
				Coordinates: coords[index]},
		}

		err := mgowrap.Save(shop)
		if goutils.HasErrorAndPrintStack(err) {
			continue
		}
	}

}
