package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// helper function
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("create a new bill name: ", reader)

	b := newBill(name)
	fmt.Printf("Created %v bill <3\n", b.name)

	return b
}
func PromptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a- add, s- save bill, t- add tip): ", reader)
	fmt.Println(opt)

}

func main() {
	myBill := createBill()
	PromptOptions(myBill)

}
