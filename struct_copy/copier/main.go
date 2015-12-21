package main

import (
	"fmt"

	"github.com/jinzhu/copier"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id   bson.ObjectId
	Name string
}

type Customer struct {
	Id   string
	Name string
}

func main() {

	var (
		user = &User{
			Id:   bson.NewObjectId(),
			Name: "Aaron"}
		customer = &Customer{}
	)

	copier.Copy(customer, user)

	fmt.Printf("%+v", customer)

}
