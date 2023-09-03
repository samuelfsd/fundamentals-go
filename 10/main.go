package main

import "fmt"

// trabalhando com funções anônimas -> Closures

func main() {

	total := func() int {
		return sum(10, 100, 1000, 1000, 1000, 1000) * 2
	}()

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
