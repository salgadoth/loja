package controllers

import (
	"log"
	"loja/models"
	"net/http"
	"strconv"
	"text/template"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	templ.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		estoque := r.FormValue("quantidade")

		precoConvFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		estoqueConvInt, err := strconv.Atoi(estoque)

		if err != nil {
			log.Println("Erro na conversão do estoque:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvFloat, estoqueConvInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idProduto)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	produto := models.BuscarProduto(idProduto)

	templ.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		estoque := r.FormValue("estoque")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int", err)
		}

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço para float", err)
		}

		estoqueInt, err := strconv.Atoi(estoque)
		if err != nil {
			log.Println("Erro na conversão do estoque para int", err)
		}

		models.AtualizaProduto(idInt, nome, descricao, precoFloat, estoqueInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
