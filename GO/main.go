package main

import "fmt"

func main() {
	myBill := newBill("Spencer's bill")

	myBill.updateTip(10)
	myBill.addItem("onion", 2.50)
	myBill.addItem("coffee", 3.50)
	myBill.addItem("pie", 6.50)
	myBill.addItem("cake", 9.50)
	fmt.Println(myBill.format())
}
