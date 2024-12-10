package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("len=%d - cap=%d - %v", len(s), cap(s), s) // tamanho - capacidade - slice

	fmt.Printf("len=%d - cap=%d - %v", len(s), cap(s), s[:2]) // uso do [:] para fatiar alguma parte do seu slice.
}
