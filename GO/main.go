package main

import (
	"fmt"
	"sort"
)

func main() {
	x := 0
	//while loop but its for in GO
	for x < 5 {
		fmt.Println("value of x is :", x)
		x += 2
	}
	//normal for loop but does not use parenthesis
	for i := 0; i < 5; i += 1 {
		fmt.Println("Odd Numbers:", i)
	}

	names := []string{"Spencer", "Delimas", "Bangoya", "Luigi", "Armstrong"}
	sort.Strings(names)
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//for-in loop but in GO its RANGE keyword
	//cycle through a slice
	for index, value := range names {
		fmt.Printf("The position at index %v is value %v \n", index, value)
	}

	//if you dont want to use the index
	for _, value := range names {
		fmt.Printf("The value %v \n", value)
	}
}
