package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//strings package
	greeting := "hello there ninja"

	fmt.Println(strings.Contains(greeting, "hello"))
	// method contains, 2 arg, 1arg is the searh body and the 2nd is the actual search, returns true or false
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	//third arg is the value to relace with, does not actaully alter the string, returns the new string
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	//returns the position containing the doule ll
	fmt.Println(strings.Split(greeting, " "))

	//sort package ~ changes the original slice
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"Spencer", "Delimas", "Bangoya", "Ethan", "zenira"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "zenira"))

}
