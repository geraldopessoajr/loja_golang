package models

import (
	"loja_camisetas/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaBancoDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func BuscarProduto(id string) Produto {
	db := db.ConectaBancoDados()
	selectProdutos, err := db.Query("select * from produtos where id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}
	defer db.Close()
	return produto
}

func CriarProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDados()
	insereDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insereDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletarProduto(id string) {
	db := db.ConectaBancoDados()
	deleteDados, err := db.Prepare("delete from produtos where id = ?")
	if err != nil {
		panic(err.Error())
	}
	deleteDados.Exec(id)
	defer db.Close()
}

func EditarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDados()
	updateDados, err := db.Prepare("update produtos set nome = ?, descricao = ?, preco =? , quantidade = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	updateDados.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
