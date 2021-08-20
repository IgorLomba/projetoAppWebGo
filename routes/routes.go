package routes

import (
	"net/http"

	"github.com/IgorLomba/projetoAppWebGo/controllers"
)

/* sempre cria a rota primeiro para o que precisa */
func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)

	/* tela de inserir produtos */
	http.HandleFunc("/newProduct", controllers.NewProduct)

	/* chama o insert que está no models */
	/* ação de inserir */
	http.HandleFunc("/insert", controllers.Insert)

	/* ação de deletar */
	http.HandleFunc("/delete", controllers.Delete)

	/* tela de editar produtos */
	http.HandleFunc("/edit", controllers.Edit)

	/* ação de atualizar/editar */
	http.HandleFunc("/update", controllers.Update)

}
