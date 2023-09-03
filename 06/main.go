package main

import "fmt"

// trabalhando com slices em go

func main() {
	// no slice quantidade de elementos fica em branco
	s := []int{2, 4, 6, 8, 10}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// quando você quiser adicionar alguma posição de um slice
	// ele não apenas aumenta quntas posições você quer aumentar da sua capacidade
	// e sim ele aponta para outro array que tem maior capacidade e dobra a capacidade do seu array atual
	s = append(s, 12)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

}
