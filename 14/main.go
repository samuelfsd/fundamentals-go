package main

type Conta struct {
	saldo int
}

// meio que um construtor que cria uma conta já apontada diretamente para o endereço
// na memoria dela
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

// saldo agora vai mudar no main também
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {

	conta := Conta{
		saldo: 100,
	}

	conta.simular(200)
	println(conta.saldo)
}
