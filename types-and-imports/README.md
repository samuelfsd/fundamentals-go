### Tipagem no go

No go existem diversos tipos, também é possível criar tipos específicos para nossa determinada demanda.

```go
  var integer int = 10 // inteiro
  var floatNumber float64 = 10.10 // ponto flutuante
  var isLoading bool = false // valor boleano
  var name string = "John Doe" // string
```

Também no go é possível criar nossos próprios tipos ou nossa própria tipagem. Segue o exemplo:

```go
  type ID int

  var f ID = 1
```

Para utilizar uma forma de realizar formatações no nosso código, vamos precisar lidar com as importações do go.
O go com auxílio das IDE's famosas tem suporte a vários pacotes para uso, um deles muito comum é o `fmt` ou o `net/http`.

Para o uso do que foi dito acima, podemos ver no arquivo `main.go` deste package.