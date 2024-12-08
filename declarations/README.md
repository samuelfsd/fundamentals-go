### Entendendo declarações e atribuições em go

Existem dois tipos de declarações em go mais utilizados.

- var (valores que podem ser modificados desde que não alterem os seus tipos)
- const (valores imutáveis)

Quando é realizada uma declaração no go se não existir uma atribuição o próprio go vai INFERIR o valor da sua variável.
Exemplo: 

```go
  var isLoading bool
  var age int
  println(isLoading) // -> false
  println(age) // -> 0
```

É possível atribuir mais de uma variável no mesmo escopo usando a sintaxe:

```go
  var (
	b bool    = true
	c int     = 10
	d string  = "Samuel"
	e float64 = 15.8
)
```

É possível atribuir e declarar usando um atalho (short hand) que seria:

```go
  const a := "A"
```

Caso você atribua uma variável mas não utiliza-lá em escopo local o go em tempo de compilação irá acusar um erro.