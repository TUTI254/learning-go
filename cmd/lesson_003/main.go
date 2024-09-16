package main

import "fmt"

func main() {
	// arrays, slices, maps and loops
	var myArray [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)

	var mySlice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	mySlice = append(mySlice, 6)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 7)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 8)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 9)
	fmt.Println(mySlice)

	//maps
	var myMap map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Carol": 26, "Eddy": 27, "Lewis": 28, "Tuti": 23}
	fmt.Println(myMap2["Carol"])

	var age, ok = myMap2["Nelson"]
	if ok {
		fmt.Println(`Nelson is `, age, ` years old`)
	} else {
		fmt.Println(`User is not in the map`)
	}

	//loops
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for i, v := range mySlice {
		fmt.Println(i, v)
	}
}
