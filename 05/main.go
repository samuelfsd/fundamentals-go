package main

// como trabalhar com arrays em go?
// entendendo slice, map

import "fmt"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Samuel"
	e float64 = 15.8
	f ID      = 1
)

func main() {
  // o array terá um valor fixo de 3 posições, não poderá háver mais posições
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 5
	myArray[2] = 10

	// mostrando qtd de posições no array
	fmt.Println("tamanho do array: ", len(myArray))

	// pegando a última posição do meu array
	fmt.Println("última posição do array: ", myArray[len(myArray)-1])

	// como percorrer o array ?
	for i, v := range myArray {
		fmt.Printf("o índice -> %d | com valor ->  %d \n", i, v)
	}
}
