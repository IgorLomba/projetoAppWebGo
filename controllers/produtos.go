package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/IgorLomba/projetoAppWebGo/models"
)

/* template must encapsula todos os templates e devolve 2 retornos (mensagem de erro) */
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	/* passei os dados para o Index (html) */
	produtos := models.SelectAllProducts()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "newProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	/* log.Println(r.Method) */
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
			/* 			panic(err.Error())
			 */
			http.Redirect(w, r, "/", http.StatusNotAcceptable)

		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
			/* 			panic(err.Error())
			 */
			http.Redirect(w, r, "/", http.StatusNotAcceptable)

		}
		if err == nil {
			models.CreateNewProduct(nome, descricao, precoConvertido, quantidadeConvertida)
			log.Println("CHEGOU NO CRIAR")
			/* volta pra pagina e diz o codigo correto */

			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
	}

	temp.ExecuteTemplate(w, "novoProduto", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	//idProdutoInt, err := strconv.Atoi(idProduto)
	_, err := strconv.Atoi(idProduto)

	if err == nil {
		models.DeleteProduct(idProduto)
	} else {
		log.Println("Erro ao deletar produto", err)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

/* só faz renderizar certinho o editarProduto */
func Edit(w http.ResponseWriter, r *http.Request) {
	/* para aprecer a tela de edit */
	idProduto := r.URL.Query().Get("id")

	produto := models.SelectOneProduct(idProduto)

	temp.ExecuteTemplate(w, "edit", produto)
}

func UpdateJson() {}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("CHEGOU NO UPDATE")
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)

			http.Redirect(w, r, "/", http.StatusNotAcceptable)

		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)

			http.Redirect(w, r, "/", http.StatusNotAcceptable)

		}

		/* log.Println(id)
		log.Println(nome)
		log.Println(descricao)
		log.Println(preco)
		log.Println(quantidade)
		log.Println(precoConvertido)
		log.Println(quantidadeConvertida) */

		if err == nil {
			models.UpdateProduct(id, nome, descricao, precoConvertido, quantidadeConvertida)
		}
		//log.Println(id, quantidadeConvertida, precoConvertido, nome, descricao)
	}
	log.Println("FEZ UPDATE")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
