package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name    string
	Phone   string
	City_Id bson.ObjectId
}

type City struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string
}

func main() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("test")
	people := db.C("people")

	id := bson.NewObjectId()

	_ = people.Insert(&City{id, "New York"})
	_ = people.Insert(&Person{Name: "Ale", Phone: "+55 53 8116 9639", City_Id: id})

	person := Person{}
	err = people.Find(bson.M{"name": "Ale"}).One(&person)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("City Id:", person.City_Id)

	city := City{}
	err = people.Find(bson.M{"_id": id}).One(&city)
	if err != nil {
		println(err.Error())
	}
	fmt.Println("City Name:", city.Name)
}
