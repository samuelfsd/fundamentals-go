package main

import "fmt"

// entendendo uma struct
// semelhante a uma classe pórem diferente

type Cliente struct {
	name     string
	age      int
	isActive bool
}

func main() {

	marlon := Cliente{
		name:     "Marlon",
		age:      30,
		isActive: true,
	}

  // é possivel mudar atributos
  marlon.isActive = false
	fmt.Println(marlon.name)
	fmt.Printf("Nome ->  %s \nIdade -> %d \nÉ ativo -> %t \n", marlon.name, marlon.age, marlon.isActive)
}
