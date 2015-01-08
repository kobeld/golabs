package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string
	Email string
}

func (this *User) CollectionName() string {
	return "users"
}

func main() {
	println("start")

	users := []*User{}
	println(detect(&users))
}

func detect(result interface{}) (name string) {
	// Get the concrete value stored in the interface.
	// Here is to get the "&[]*User{}"
	userArrayPointer := reflect.ValueOf(result)
	// It should be pointer (reflect.Ptr)
	println(userArrayPointer.Kind())

	// Using Elem() to get the concrete value in the pointer.
	// Here is to get the "[]*User{}"
	userArray := userArrayPointer.Elem()

	// With the .Kind(), it now should be slice type (reflect.Slice)
	println(userArray.Kind())

	// With the .Type() to get the real data type "[]*main.User"
	fmt.Printf("%+v \n", userArray.Type())

	userElemt := userArray.Type().Elem().Elem()
	fmt.Printf("%+v \n", userElemt)

	userValue := reflect.New(userElemt)
	method := userValue.MethodByName("CollectionName")
	if method.IsValid() {
		name = method.Call([]reflect.Value{})[0].String()
	}

	return
}
