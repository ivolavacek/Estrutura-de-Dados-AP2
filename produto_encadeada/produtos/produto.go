package produtos

import "fmt"
//import "sort"

var TotalProdutosJaCadastrados = 0
var ListaDeProdutos Lista

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

	ListaDeProdutos.insere(p)

	return p
}

 // Requisito funcional 1
func AtualizarPrecoProduto(id int, novoPreco float64) bool {
    for no := ListaDeProdutos.ptlista; no != nil; no = no.prox {
        fmt.Println(no.produto)
        if no.produto.Id == id {
            no.produto.Preco = novoPreco
            fmt.Println(no.produto)
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


func (l *Lista) OrdenarPorNome() {
    if l.ptlista == nil || l.ptlista.prox == nil {
        return // A lista está vazia ou tem apenas um elemento
    }

    trocado := true
    for trocado {
        trocado = false
        ptr := l.ptlista
        for ptr.prox != nil {
            if ptr.produto.Nome > ptr.prox.produto.Nome {
                // Troca os produtos
                ptr.produto, ptr.prox.produto = ptr.prox.produto, ptr.produto
                trocado = true
            }
            ptr = ptr.prox
        }
    }
}


func (l *Lista) ExibirProdutosPorNome() {
    l.OrdenarPorNome() // Primeiro ordena a lista pelo nome

    // Depois, exibe os produtos como antes
    for no := l.ptlista; no != nil; no = no.prox {
        fmt.Printf("Id: %d, Nome: %s, Descrição: %s, Preço: %.2f\n", no.produto.Id, no.produto.Nome, no.produto.Descricao, no.produto.Preco)
    }
}
