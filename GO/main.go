package main

import "fmt"

func main() {
	age := 45
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)
	// conditionals - if statement
	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not above than 40")
	}

	// nested ifs - if inside loops
	name := []string{"Zuko", "Sokka", "Katara", "Toph", "Aang"}
	for index, value := range name {
		if index == 1 {
			fmt.Println("continuing at pos \n", index)
			continue //means break out of this iteration and contnue the loop
		}
		if index > 2 {
			fmt.Println("Breaking at pos\n", index)
			break //breaks out of the loop completely
		}
		fmt.Printf("The value %v at index is %v \n", index, value)
	}
}
