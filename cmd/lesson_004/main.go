package main

import "fmt"

// more on strings, bytes, runes
func main() {
	// strings
	var myString string = "Hi"
	fmt.Println(myString)

	// bytes
	var myBytes []byte = []byte("Hi")
	fmt.Println(myBytes)

	// runes
	var myRunes []rune = []rune("Hi")
	fmt.Println(myRunes)

}
