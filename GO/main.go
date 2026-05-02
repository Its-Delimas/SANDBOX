package main

import "fmt"

func main() {

	//map format... varName := map[key type] valuetype{***}
	menu := map[string]float64{
		"Soup":   4.99,
		"Pie":    7.99,
		"salad":  6.99,
		"Coffee": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	//looping through a map
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
	//ints as key type
	Languages := map[int]string{
		001000: "Node",
		001001: "Express",
		001002: "NextJs",
		001003: "GO",
		001004: "PostgreSQL",
		001005: "MongoDB",
		001006: "Tailwind v4",
		001007: "React",
	}
	// changing value
	Languages[001003] = "GOLANG"
	for _, v := range Languages {
		fmt.Println(v)
	}

}
