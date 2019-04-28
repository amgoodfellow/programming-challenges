package main

import (
	"fmt"
)

func mapping(s string) float64 {
	var sum float64
	for _, j := range s {
		sum += float64(j)
	}
	return float64(sum) / float64(len(s))
}

func main() {
	fmt.Printf("%v\n", mapping("Hello"))
	fmt.Printf("%v\n", mapping("Test"))
	fmt.Printf("%v\n", mapping("QWERTY"))
	fmt.Printf("%v\n", mapping("!@#$%^&"))
	fmt.Printf("%v\n", mapping("This is fun!"))
}
