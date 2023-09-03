package main

// entendendo criação de tipos em go
// criando sua propria tipagem

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Samuel"
	e float64 = 15.8
	f ID      = 1
)

func main() {
	println(f)
}
