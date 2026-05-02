package main

import "fmt"

func updateName(x *string) {
	*x = "Brave"

}
func main() {
	name := "Chrome"
	fmt.Println(name)

	//pointers &var_name
	fmt.Println("Memory address of the name is : ", &name)
	n := &name //pointer value can be stored in var...&
	fmt.Println("Memory address of n: ", &n)
	fmt.Println("Value at memory address:", *n) //to get the exact value of the pointer, we use *
	updateName(n)
	fmt.Println(name)
}

// non pointer and pointer wrappers
