package main

import "fmt"

// structs and interfaces

type Person struct {
	Name string
	Age  int
}

type gender interface {
	getGender() string
}

type Male struct {
	gender string
}

type Female struct {
	gender string
}

func (m Male) getGender() string {
	return m.gender
}

func (f Female) getGender() string {
	return f.gender
}

func main() {
	var myPerson Person = Person{
		Name: "Tuti",
		Age:  23,
	}
	fmt.Println(myPerson)

	fmt.Println(myPerson.Name)
}
