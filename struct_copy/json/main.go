package main

import (
	"encoding/json"
	"fmt"

	"github.com/kobeld/goutils"

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
		user = User{bson.NewObjectId(), "Aaron"}
	)

	u, err := json.Marshal(user)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	var customer Customer
	err = json.Unmarshal(u, &customer)
	fmt.Printf("%+v", customer)
}
