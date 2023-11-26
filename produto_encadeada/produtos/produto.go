package produtos

import "fmt"

var TotalProdutosJaCadastrados = 0
var listaDeProdutos Lista

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
}

type No struct {
	produto Produto
	prox *No
}

type Lista struct {
	ptlista *No
}

/*
Define um id de um produto, considerando todos os produtos já cadastrados.
*/
func (p *Produto) definirId() {
	//TotalProdutosJaCadastrados++
	p.Id = TotalProdutosJaCadastrados
}

/*
Exibe as informações de um produto no terminal.
*/
func (p *Produto) Exibir() {
	fmt.Println("\nProduto", p.Id)
	fmt.Println(p.Nome)
	fmt.Println(p.Descricao)
	fmt.Printf("Preço: R$ %.2f\n", p.Preco)
}

/*
Retorna um elemento do tipo Produto, com um id a ser definido ou com um id
pré-definido.
*/
func criar(nome, descricao string, preco float64, id int) Produto {
	p := Produto { Nome: nome, Descricao: descricao, Preco: preco }

	TotalProdutosJaCadastrados++ // conserta bug

	if id == -1 {
		p.definirId()
	} else {
		p.Id = id
	}

	listaDeProdutos.insere(p)

	return p
}

 // Requisito funcional 1
func AtualizarPrecoProduto(id int, novoPreco float64) bool {
    for i, produto := range Produtos {
        if produto.Id == id {
            Produtos[i].Preco = novoPreco
            return true
        }
    }
    return false
}

// Requisito não funcional 1
func (l *Lista) insere(novoProduto Produto) {
    novoNo := &No{produto: novoProduto}

    if l.ptlista == nil {
        l.ptlista = novoNo
    } else {
        pont := l.ptlista
        for pont.prox != nil {
            pont = pont.prox
        }

        pont.prox = novoNo
    }
}