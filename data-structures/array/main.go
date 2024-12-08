package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	fmt.Println("Isso aqui é um Array: ", myArray) // -> [1,2,3]
	fmt.Println(len(myArray))                      // informa o tamanho do array -> 3

	// percorre o array mostrando o indíce e o valor (i e v)
	for i, v := range myArray {
		fmt.Printf("Indíce: %d - Valor: %d \n", i, v)
	}
}
