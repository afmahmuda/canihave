package main

import (
	"fmt"

	"github.com/afmahmuda/canihave"
)

func main() {

	fmt.Println("==== string ======")

	var pointerString *string
	fmt.Println(canihave.ValueOfOrDefault(pointerString, "default value"))

	stringValue := "a value"
	pointerString = &stringValue
	fmt.Println(canihave.ValueOfOrDefault(pointerString, "default value"))
	fmt.Println(canihave.ValueOfOrDefault(canihave.PointerOf("also value"), "default value"))

	fmt.Println("==== int =========")

	var pointerInt *int
	fmt.Println(canihave.ValueOfOrDefault(pointerInt, 10))

	intValue := 11
	pointerInt = &intValue
	fmt.Println(canihave.ValueOfOrDefault(pointerInt, 10))
	fmt.Println(canihave.ValueOfOrDefault(canihave.PointerOf(12), 10))

	fmt.Println("==== struct =======")

	type person struct {
		name string
		age  int
	}

	var pointerPerson *person
	fmt.Println(canihave.ValueOfOrDefault(pointerPerson, person{name: "anna", age: 76}))

	personValue := person{name: "olaf", age: 19}
	pointerPerson = &personValue
	fmt.Println(canihave.ValueOfOrDefault(pointerPerson, person{name: "anna", age: 76}))
	fmt.Println(canihave.ValueOfOrDefault(canihave.PointerOf(person{name: "elsa", age: 80}), person{name: "anna", age: 76}))
}
