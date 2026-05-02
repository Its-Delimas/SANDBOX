package main

import "fmt"

func main() {
	//arrays
	var ages [4]int = [4]int{20, 25, 30, 40} //method 1
	var numbers = [4]int{20, 25, 30, 40}     //method 2

	//fmt.Println(ages) ~printing arrays
	fmt.Println(ages, len(ages)) //return with length of arrays,
	fmt.Println(numbers)

	//slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores = append(scores, 67)
	scores = append(scores, 97)
	fmt.Println(scores, len(scores))

	// slice ranges - inclusive of first number but no te second
	rangeOne := scores[1:3]  //index 1 and 2 but not 3
	rangeTwo := scores[2:]   //from index two going on forward to the very last
	rangeThree := scores[:3] //from start upto but not including position three

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

}
