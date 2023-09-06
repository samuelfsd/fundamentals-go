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

// vinculando um método a uma struct
func (c Client) logout() {
	c.isActive = false
	fmt.Printf("O Cliente %s foi desligado", c.name)
}

// entendendo o conceito de interface 
type People interface {
	Hello() string
}

func (c Client) Hello() string {
	hello := "hello"

	return hello
}

func Fala(c Client) {
	c.Hello()
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
	marlon.logout()

	fmt.Println(marlon.city)

	fmt.Println(marlon)
}
