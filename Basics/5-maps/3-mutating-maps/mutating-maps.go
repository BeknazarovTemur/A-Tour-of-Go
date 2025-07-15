package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42 // add
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48 // add
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // delete
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // get
	fmt.Println("The Value:", v, "Present?", ok)
}
