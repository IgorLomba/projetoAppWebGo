package main

import (
	"fmt"
	"net/http"

	"github.com/IgorLomba/projetoAppWebGo/db"
	"github.com/IgorLomba/projetoAppWebGo/models"
	"github.com/IgorLomba/projetoAppWebGo/routes"
)

func main() {
	db := db.ConectaBd()
	//isso é o um insert into
	db.Create(&models.Produto{Nome: "Agua", Descricao: "Beba agua", Preco: 1, Quantidade: 10})
	db.Where("nome = ?", "Agua").Delete(&models.Produto{})

	/*
		//para criar uma tabela corretamente pela query
		var pessoa []models.Pessoa

		db.Raw(`create table if not exists pessoas (
			id SERIAL PRIMARY KEY,
			nome varchar(100)
			)`).Scan(&pessoa)

		db.Create(&models.Pessoa{Nome: "Teste"})
	*/

	routes.CarregaRotas()
	fmt.Println("SUCESSO!")
	/* ouça e sirvaa	 na porta 3 mil */
	http.ListenAndServe(":3000", nil)

}
