package main

func main() {

	// Mémoria -> Endereço -> valor
	// váriavel -> ponteiro -> endereço -> valor
	a := 10

	println(a)
	println(&a)

	// apontando para o endereço
	var pointer *int = &a

	b := &a
	println(pointer)
	println(*b)
}
