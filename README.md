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



### Correção do Bug 1

Esse erro se dá pelo índice de "inicioFila" e "fimFila" serem iguais depois que já tem mais de um pedido expedido, ou seja, a fila de pedidos está vazia. Foi necessário a inserção de um "if" para que nesse caso, a função "excluir" do "filaPedidos.go", retorne um pedido vazio ao invés de procurar um pedido com índice "-1". 

### Correção do Bug 2

A variável "TotalProdutosJaCadastrados" é responsável pela atribuição do id de um produto. Ela estava dentro de uma função própria, e essa função só é chamada quando não há um id definido para o produto. 
Os produtos que estão no "dados.csv" já vem com id atribuido, portanto a função que determina o id de um produto não era chamada. Então a variável que determina os ids criados manualmente no projeto não considerava os ids que já foram usados na importação "automática" permitindo assim a duplicação de ids.
Para corrigir esse erro bastou mover 
a variável "TotalProdutosJaCadastrados" de "definirId" para "criar" no arquivo "produtos.go", assim todos os produtos criados terão que atualizar essa variável, não apenas aqueles que não tem id pré-definido.          



------------------------------------------


**Comentários finais**
O trabalho foi feito no site 'replit.com', quando foi passado o código para o VSCode, para criar o .exe a partir do Command Prompt, algumas funcionalidades foram afetadas.
No VSCode, o arquivo main.go não rodou e o aviso de erro era que as funções definidas nos outros arquivos (metricas.go, pedido.go e produto.go) estavam 'undefined', mesmo com todos os arquivos incluidos no package 'main' e funções começadas em letras maiúsculas.
O arquivo .exe está funcionando bem, exceto no cadastro de produtos onde o programa pula o cadastro do nome dos produtos.
No replit.com, o programa funcionou perfeitamente. Segue o link [Replit](https://replit.com/@IvoLavacek/AP1-Goland-4#main.go)