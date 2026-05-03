package main

import "fmt"

func main() {
	myBill := newBill("Spencer's bill")

	fmt.Println(myBill.format())
}
