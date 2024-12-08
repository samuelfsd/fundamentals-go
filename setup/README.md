### Instalação do go

Para instalação do go você pode realizar pela plataforma mais indicada o próprio site do go:
[Go Dev](https://go.dev/dl/).

ou usar algum package manager do seu sistema operacional.


Quando é realizada a instalação do go ele vai ficar apontando para um path, para verificar use esse comando:

```go
  go env
```

Depois de verificar e entrar na pasta selecionada, podemos ver um padrão do go em seus arquivos com três pastas.

- bin (arquivos binários - pré compilações dentre outros motivos)
- pkg (pasta onde existem compilações utilizada geralmente em arquivos .a)
- src (onde é possível criar seus projetos executar arquivos .go e até mesmo rodar arquivos .c)
