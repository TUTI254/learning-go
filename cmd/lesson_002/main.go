package main

import "fmt"

//lets learn functions

func main() {
	profile()
	myAge(23)
	var result int = divide(10, 2)
	fmt.Println(`The result is: `, result)
}

func profile() {
	fmt.Println("hello world, my name is Tuti")
}

func myAge(value int) {
	fmt.Println(`I am `, value, ` years old`)
}

func divide(a, b int) int {
	var result int = a / b
	return result
}
