package main

import (
	"fmt"
	"reflect"
)

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Samuel"
	e float64 = 15.8
	f ID      = 1
)

func main() {
	fmt.Println(f)
	fmt.Printf("Este é o tipo de e %T", e)          // -> indica o tipo
	fmt.Println("Este tipo é: ", reflect.TypeOf(f)) // outra forma de inficação de tipo -> main.ID
}
