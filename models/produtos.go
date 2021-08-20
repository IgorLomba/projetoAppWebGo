package models

import (
	"encoding/json"
	"log"

	"github.com/IgorLomba/projetoAppWebGo/db"
)

type Produto struct {
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Id         int     `json:"id"`
	Quantidade int     `json:"quantidade"`
}

type Pessoa struct {
	Nome string `json:"nome"`
	Id   int    `json:"id"`
}

func SelectAllProducts() []Produto {
	db := db.ConectaBd()
	var produtos []Produto
	db.Raw("select * from produtos order by id asc").Scan(&produtos)
	return produtos
}

func SelectAllProductsJson() string {
	produto := SelectAllProducts()
	json, _ := json.Marshal(produto)
	return string(json)
}

func CreateNewProduct(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBd()
	produtos := Produto{Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade}
	db.Select("nome", "descricao", "preco", "quantidade").Create(&produtos)
}

func DeleteProduct(idProduto string) {
	db := db.ConectaBd()
	log.Println("PRINTT", SelectOneProduct(idProduto))
	db.Delete(&Produto{}, idProduto)
}

func SelectOneProduct(idproduto string) Produto {
	db := db.ConectaBd()

	/* declaro um objeto tipo Produto */
	var produto Produto
	//faço um select pelo ID
	db.First(&produto, idproduto)

	return produto
}

/* atualiza o produto através do ID*/
func UpdateProduct(id, nome, descricao string, preco float64, quantidade int) {

	db := db.ConectaBd()

	produto := SelectOneProduct(id)
	produto.Nome = nome
	produto.Descricao = descricao
	produto.Preco = preco
	produto.Quantidade = quantidade

	//atualiza os dados do produto individual
	db.Save(&produto)
}
