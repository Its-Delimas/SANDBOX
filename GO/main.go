package main

import "fmt"

func main() {
	age := 35
	name := "spencer"

	// print - same output on the same line
	fmt.Print("hello")
	fmt.Print("ninja")

	// Println - opens new line automatically
	fmt.Println("hello, Spencer")
	fmt.Println("Go lore")
	fmt.Println("my name is:", name, " and my name is: ", age)

	// formatted strings `Printf`, %_ = format specifiers
	fmt.Printf("my age is %v and my name is %v \n", age, name)  //%v = use default variable type
	fmt.Printf("my age is %v and my name is %q \n", age, name)  //%q= quotes around strings
	fmt.Printf("age is of type %T\n", age)                      //%T = returns type of var
	fmt.Printf("you scored %0.2f ponts in your gcpa\n", 225.59) //%f = floats, o.2 for 2d.p

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Printf("The saved string is : %v", str)

}
