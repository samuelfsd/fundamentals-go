package main

// É possível criar constrains -> são tipos especificos para serem subisituidos na generic ou algo do tipo

type Number interface {
	~int | float64
}

type MyNumber int

func Sum[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

// em vez de fazer duas funções iguais para retornos de tipos diferentes usamos generics
// func SumFloat(m map[string]float64) float64 {
// 	var soma float64

// 	for _, v := range m {
// 		soma += v
// 	}

//		return soma
//	}

// comparable -> apenas valores possíveis de serem comparados -> apenas igualdade
// package constrains
func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	x := map[string]int{"Samuel": 1000, "Jão": 500, "josé": 3000}
	y := map[string]float64{"Samuel": 50.0, "Jão": 10.0, "josé": 20.0}
	z := map[string]MyNumber{"Samuel": 50.0, "Jão": 10.0, "josé": 20.0}

	println(Sum(y))
	println(Sum(x))
	println(Sum(z))

	println(Compare(10, 10.0))
}
