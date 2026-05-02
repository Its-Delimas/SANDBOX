package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("see you soon %v \n", n)
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	sayGreeting("Spencer")
	sayGreeting("Luigi")
	sayBye("Mario")

	cycleName([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleName([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(14.28)
	fmt.Printf("Circle is %0.2f", a1)

}
