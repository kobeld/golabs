package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

func main() {

	id := "4fd7814f558fbe76ff000038"

	oId := ToObjectId(id)

	println("------>")
	fmt.Printf("%+v\n", oId)

}

func ToObjectId(id string) bson.ObjectId {
	if id == "" {
		return ""
	}

	defer func() {
		if err := recover(); err != nil {
			println("hello")
		}
	}()

	return bson.ObjectIdHex(id)
}
