package main

import (
	"github.com/theplant/qortexapi"
	"reflect"
	"unsafe"
)

type Dog struct {
	Name  string
	Legs  int
	Liked []string
}

func main() {

	dog := Dog{
		Name:  "Mike",
		Legs:  4,
		Liked: []string{"sleep", "bones", "roar"},
	}

	dogType := reflect.TypeOf(dog)
	println(dogType.Name())
	println(dogType.Kind().String())
	println(dogType.FieldAlign())

	// dogvalue := reflect.ValueOf(dog)

	embedUser := reflect.TypeOf(qortexapi.User{})
	println(embedUser.FieldAlign())

	numberOne := 1
	var hello float32 = 0.0000000000000000000000000000032132132132131

	println(unsafe.Sizeof(numberOne))
	println(unsafe.Sizeof(hello))
	println(hello)
	println(reflect.TypeOf(hello).String())
	println(reflect.TypeOf(hello).Size())
}
