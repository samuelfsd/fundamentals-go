package main

import (
	"errors"
	"fmt"
)

func main() {

	valor, err := sum(49, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

}

// trabalhando com funções
// passando paramêtros e informando o tipo de retorno

// conseguir retornar masi de um tipo -> vai ser bem útil no tratamento de erros
// já que no go isso não é muito bem desenvolvido
func sum(a int, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	// o nil representa -> vazio ou
	return a + b, nil
}
