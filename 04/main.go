package main

// importação como funciona ?
// trabalhando com formatação de valores
// package fmt | pacote de formatação, como importar ?
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
	fmt.Printf("O tipo de E é %T", e)
  println()
	fmt.Printf("O valor de E é %v", e)
}
