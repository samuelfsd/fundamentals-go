package main

import "fmt"

// entendendo hashtables em go
// são representadas como -> Maps

func main() {

	salaries := map[string]int{"Samuel": 2000, "José": 1000, "Jobs": 10000000}

	fmt.Println(salaries)

	// salário especifico
	fmt.Println(salaries["Samuel"])

	//remover uma chave especifica da hashtable
	delete(salaries, "Jobs")

	fmt.Println(salaries)

	// criando com make
	// sal := make(map[string]int)

	// percorrendo a hashtalbe
	for name, salary := range salaries {
		fmt.Printf("O %s tem o salário de %d \n", name, salary)
	}

	// como ignorar a chave e percorrer apenas o valor ? blank identifier
	for _, salary := range salaries {
		fmt.Printf("O sálario é de %d \n", salary)
	}
}
