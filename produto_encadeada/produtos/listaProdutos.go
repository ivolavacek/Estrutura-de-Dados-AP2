package produtos

import (
	m "mcronalds/metricas"
	"strings"
)

const maxProdutos = 50

var Produtos [maxProdutos]Produto
var totalProdutos = 0

func tentarCriar(nome, descricao string, preco float64, id int) Produto {
	if id != -1 {
		_, idProcurado := BuscarId(id)
		if idProcurado != -1 { return Produto{} }
	}

	return criar(nome, descricao, preco, id)
}

/*
Adiciona um produto com nome, descrição e preço à lista de produtos.
Adiciona o produto primeiro espaço vazio da lista.
Caso já exista um produto com o mesmo id, não adiciona e retorna -3.
Caso já exista um produto com o mesmo nome, não adiciona e retorna erro -2.
Retorna -1 caso a lista esteja cheia, ou o número de produtos cadastrados em caso de sucesso.
*/
func AdicionarUnico(nome, descricao string, preco float64, id int) int {
	if totalProdutos == maxProdutos { return -1 } // Overflow

	for _, produto := range Produtos {
		if (produto == Produto{}) { break }
		if produto.Nome == nome {
			return -2
		}
	}

	produtoCriado := tentarCriar(nome, descricao, preco, id)
	if (produtoCriado == Produto{}) { return -3 }

	Produtos[totalProdutos] = produtoCriado
	totalProdutos++
	m.M.SomaProdutosCadastrados(1)
	return totalProdutos
}

/*
Localiza um produto a partir do seu id.
Retorna o produto encontrado e a sua posição na lista, em caso de sucesso.
Retorna um produto vazio e -1 em caso de erro.
*/
func BuscarId(id int) (Produto, int) {
	for no := ListaDeProdutos.ptlista; no != nil; no = no.prox {
		if (no.produto == Produto{}) { break }
		if no.produto.Id == id {
			return no.produto, no.produto.Id
		}
	}

	return Produto{}, -1
}

/*
Localiza produtos que iniciem com a string passada.
Retorna um slice com todos os produtos encontrados, e o tamanho do slice.
*/
func BuscarNome(comecaCom string) ([]Produto, int) {
	var produtosEncontrados []Produto

	for no := ListaDeProdutos.ptlista; no != nil; no = no.prox {
        if (no.produto == Produto{}) { break }

		if strings.HasPrefix(no.produto.Nome, comecaCom) {
			produtosEncontrados = append(produtosEncontrados, no.produto)
		}
    }

	return produtosEncontrados, len(produtosEncontrados)
}

/*
Exibe todos os produtos cadastrados.
*/
func Exibir() {
	for no := ListaDeProdutos.ptlista; no != nil; no = no.prox {
        if (no.produto == Produto{}) { break }
		no.produto.Exibir()
    }
}

/*
Remove um produto da lista a partir do seu id.
Retorna -2 caso não haja produtos na lista.
Retorna -1 caso não haja um produto com o id passado, ou 0 em caso de sucesso.
*/

func Excluir(id int) int {
	if totalProdutos == 0 { return -2 }

	_, idProcurado := BuscarId(id)
	if idProcurado == -1 { return -1 }

	ListaDeProdutos.RemovePorId(id)
	totalProdutos--

	m.M.SomaProdutosCadastrados(-1)
	return 0
}

func (l *Lista) RemovePorId(id int) Produto {
    ant, ptr := l.buscaPorId(id) // Esta função precisa ser definida para buscar o produto pelo ID

    if ptr == nil {
        return Produto{} // Produto não encontrado
    }

    if ant == nil {
        // O produto a ser removido é o primeiro da lista
        l.ptlista = ptr.prox
    } else {
        // O produto a ser removido está no meio ou no final da lista
        ant.prox = ptr.prox
    }

    return ptr.produto // Retorna o produto removido
}

func (l *Lista) buscaPorId(id int) (*No, *No) {
    var ant *No
    ptr := l.ptlista

    for ptr != nil && ptr.produto.Id != id {
        ant = ptr
        ptr = ptr.prox
    }

    return ant, ptr
}
