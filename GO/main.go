package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	switch opt {
	case "a":
		name, _ := getInput("Item name", reader)
		price, _ := getInput("Item Price", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number:")
			PromptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Items added successfully")
		fmt.Println("item name: ", name, "item price: ", price)
		PromptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount (ksh)", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number:")
			PromptOptions(b)
		}
		b.updateTip(t)

		fmt.Println(tip)
		fmt.Println("Tip added successfully: ")
		PromptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved a file - ", b.name)
	default:
		fmt.Println("Invalid option...")
		PromptOptions(b)
	}

}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)

	}
	fmt.Println("Bill is saved succesfully")
}

func main() {
	myBill := createBill()
	PromptOptions(myBill)

}
