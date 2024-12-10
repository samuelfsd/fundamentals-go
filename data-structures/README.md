### Estruturas de dados mais utilizadas em Go

Para falar sobre as estruturas de dados mais utilizadas, vou usar um padrão de pastas para cada estrutura e neste README vou introduzir conceitos teóricos sobre cada estrutura mas para visualização de códigos mais detalhados entre na pasta que determina a estrutura e si.

### Array

No go a estrutura de arrays é fixa e não podendo modificar o seu tamanho mas sim apenas os elementos que estão dentro dele, se o array for uma constante nem os elementos podem ser modificados.

- O tamanho do array faz parte do seu próprio tipo (portanto não pode mudar).
- Você pode atribuir um array a outro copiando todos os seus elementos.
- Se você passar um array como parâmetro de uma função, ele receberá a cópia do mesmo e não a referência.

### Map

...

### Slice

O Slice no go é bem semelhante ao array porém com o seu tamanho podendo ser relativo, quando você aumenta o tamanho do seu slice o go por baixo dos panos procura outro valor em memória que tenha maior capacidade para suprir a demanda que você precisa (a capacidade pode aumentar bastante caso você necessite de um tamanho muito maior, tornando isso um processo custoso para o go).

Dica: Tente sempre inicializar um slice com um valor próximo ao que você deseja armazenar para não ter grandes custos em memória (assim evitando o go procurar algum outro array maior em memória para satisfazer sua condição de espaço).