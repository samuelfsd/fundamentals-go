package main

import "fmt"

// trabalhar com ponteiros é conseguir trabalhar globalmente com as váriaveis com seu endereço de mémoria
func sum(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	var1 := 10
	var2 := 20

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Printf("Resultado da soma é: %d \n", sum(&var1, &var2))
	fmt.Println(var1)
	fmt.Println(var2)

  ///retorno 10 | 20 antes do print e 50 | 50 depois do print
}
