package main

import "fmt"

// entendendo uma struct
// semelhante a uma classe pórem diferente
type Address struct {
	city   string
	number int
	state  string
}

// structs podem ter composições
// ou encadeamento de structs
type Client struct {
	name     string
	age      int
	isActive bool
	Address
}

func main() {

	marlon := Client{
		name:     "Marlon",
		age:      30,
		isActive: true,
	}

	// é possivel mudar atributos
	marlon.isActive = false
	fmt.Println(marlon.name)
	fmt.Printf("Nome ->  %s \nIdade -> %d \nÉ ativo -> %t \n", marlon.name, marlon.age, marlon.isActive)

	marlon.city = "Paulista PB"

	fmt.Println(marlon.city)
}
