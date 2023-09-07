package main

import "fmt"

// type assertion

func main() {

	var myvar interface{} = "mumu"

	println(myvar)

	res, ok := myvar.(int)

	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}
