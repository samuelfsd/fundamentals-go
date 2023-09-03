package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(10, 100, 1000, 1000, 1000, 1000))
}

// trabalhando com parametros indefinidos
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
