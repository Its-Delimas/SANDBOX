package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// receiver function to format bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...ksh %v \n", k+":", v) //%-25v - add empty space, uniform lineup: - for rigt space, + for left space
		total += v
	}
	//tip
	fs += fmt.Sprintf("%-25v ...ksh %v\n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...ksh %0.2f", "total:", total+b.tip) //%-25v - add empty space, uniform lineup: - for rigt space, + for left space
	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add items
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
