package main

import "fmt"

var points []int = []int{20, 30, 40, 50}

func sayHello(n string) {
	fmt.Printf("hello %v", n)
}

func showScore() {
	fmt.Println("You scored this many points:", score)
}
