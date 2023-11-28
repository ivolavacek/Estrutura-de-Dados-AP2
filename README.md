# Estruturas de Dados - 2023.2

# AP2_Estrutura_de_Dados

## Participantes
- André Silveira
- Ivo Lavacek
- João Araújo
- Pedro Lustosa

## Como funciona
Ao executar o arquivo main.go, aparecerá um menu com todas as funcionalidades do programa, sendo estas:

1. Cadastrar produto;
2. Remover produto;
3. Buscar produto por id;
4. Buscar produto por nome;
5. Exibir todos os produtos;
6. Adicionar pedido;
7. Expedir pedido;
8. Exibir métricas do sistema;
9. *NEW* Atualizar preço de um produto;
10. *NEW* Exibir produtos por nome;
20. Exibir todos os pedidos em andamento;
21. Cadastrar produtos em lote;
100. Sair do programa;

Segue as melhorias propostas de algumas das funcionalidades do programa

### Considerando a lista sequencial de produtos

Seguem as melhorias do projeto sem considerar a transformação em lista encadeada

#### Atualização do preço de um produto

Foi feita a adição dessa função no menu de opções sendo pedido o Id do produto e o novo preço. Caso o Id não for encontradoo sistema retorna uma mensagem de erro, caso ao contrário o preço do produto é atualizado e é retornado uma mensagem de sucesso.  

#### Adição de métrica de Ticket Médio

Quando há uma atualização da expedição de um pedido, o novo atributo do struct "Metricas" (ticketMedio) é calculado como a divisão do "faturamentoTotal" pelo "pedidosEncerrados".

#### Exibição de produtos por nome

Para a ordenação dos produtos por nome foi feita uma nova função onde serão adicionados todos os produtos à uma variável slice de produtos. Foi adicionada também uma nova função de ordenação, Bubblesort, para ordenar os produtos da nova variável criada. Então a função "ListarProdutosPorNome" intera sobre a variável ordenada exibindo cada produto pela ordem alfabética.   

### Transformação da lista sequencial em lista encadeada

Para a transformação da lista de produtos sequencial para uma lista encadeada, começamos criando 2 structs: um struct "No" com 2 atributos ("produto" do tipo Produto e um "prox" com o endereço de um outro "No") e um outro struct do tipo "Lista", com um atributo "ptlista" (que é o endereço do primeiro ponteiro da lista).
Foi criada então a variável "ListaDeProdutos", do tipo Lista, para substituir a variável "Produtos".
A primeira função implementada para a substituição entre essas variáveis foi a função "insere" em produto.go e a chamada dessa função dentro da função "criar" no mesmo arquivo produto.go, assim já pudemos acessar a lista encadeada de qualquer lugar do código.
Em seguida fizemos a função "OrdenarPorNome", usando como base o material dado em aula sobre listas encadeadas, e a função "ExibirProdutosPorNome" que chama a função de ordenação, usa um "for" para interar sobre cada nó da lista e imprime o que está nesse endereço de nó usando, por exemplo, "no.produto.Id" para o Id de cada produto.
No código, principalmente nos arquivos produto.go e listaProdutos.go, há várias funções que interam sobre a variável Produtos usando algo parecido com "for produto := range Produtos", esse "for" deve ser corrigido para em todos os lugares do código ser usado a lista de produtos encadeada, e não a sequencial. Então fizemos a troca por algo parecido com "for no := ListaDeProdutos.ptlista; no != nil; no = no.prox", permitindo interar pela variável "ListaDeProdutos" que contém os nós que contém os produtos. Sendo assim, para interagir pelos produtos em si, as funções precisam chamar "no.produto", e para acessar os atributos de Produto, é preciso chamar, pro exemplo, no.produto.Nome para nome do produto ou no.produto.Id para o Id do produto.
Para todo o código rodar considerando a lista encadeada, foi necessário mais a criação de 2 funções, ambas para auxiliar a remoção de produtos, estas são "RemovePorId" e "buscaPorId".
A função "Excluir" em listaProdutos.go verifica se o Id passado como atributo é válido para exclusão, então chama a função "RemovePorId" que, por sua vez chama "buscaPorId". Esta última retorna os nós adjacentes ao nó que representa o produto com Id especificado, e então a função "RemovePorId" ajusta a variável "ListaDeProdutos", removendo o nó do produto específico, mas garantindo a integridade da lista encadeada.

### Correção do Bug 1

Esse erro se dá pelo índice de "inicioFila" e "fimFila" serem iguais depois que já tem mais de um pedido expedido, ou seja, a fila de pedidos está vazia. Foi necessário a inserção de um "if" para que nesse caso, a função "excluir" do "filaPedidos.go", retorne um pedido vazio ao invés de procurar um pedido com índice "-1". 

### Correção do Bug 2

A variável "TotalProdutosJaCadastrados" é responsável pela atribuição do id de um produto. Ela estava dentro de uma função própria, e essa função só é chamada quando não há um id definido para o produto. 
Os produtos que estão no "dados.csv" já vem com id atribuido, portanto a função que determina o id de um produto não era chamada. Então a variável que determina os ids criados manualmente no projeto não considerava os ids que já foram usados na importação "automática" permitindo assim a duplicação de ids.
Para corrigir esse erro bastou mover 
a variável "TotalProdutosJaCadastrados" de "definirId" para "criar" no arquivo "produtos.go", assim todos os produtos criados terão que atualizar essa variável, não apenas aqueles que não tem id pré-definido.          



------------------------------------------