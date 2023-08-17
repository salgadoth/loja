package models

import (
	"loja/db"
)

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Estoque   int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	todosProdutos, err := db.Query("Select * from public.produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for todosProdutos.Next() {
		var id, estoque int
		var nome, descricao string
		var preco float64

		err = todosProdutos.Scan(&id, &nome, &descricao, &preco, &estoque)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Estoque = estoque

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, estoque int) {
	db := db.ConectaComBancoDeDados()

	insert, err := db.Prepare("insert into public.produtos(nome, descricao, preco, quantidade) values ($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(nome, descricao, preco, estoque)
	defer db.Close()
}

func DeletaProduto(idProduto string) {
	db := db.ConectaComBancoDeDados()

	delete, err := db.Prepare("delete from public.produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(idProduto)
	defer db.Close()
}

func BuscarProduto(idProduto string) Produto {
	db := db.ConectaComBancoDeDados()

	query, err := db.Query("select * from public.produtos where id = $1", idProduto)

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for query.Next() {
		var id, estoque int
		var nome, descricao string
		var preco float64

		err = query.Scan(&id, &nome, &descricao, &preco, &estoque)

		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Estoque = estoque
	}

	defer db.Close()
	return produto
}

func AtualizaProduto(id int, nome, descricao string, preco float64, estoque int) {
	db := db.ConectaComBancoDeDados()

	query, err := db.Prepare("update public.produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(nome, descricao, preco, estoque)

	defer db.Close()
}
