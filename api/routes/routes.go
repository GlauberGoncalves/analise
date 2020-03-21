package routes

import (
	"net/http"

	"github.com/glaubergoncalves/analise/api/controllers"
)

// CarregaRotas Carrega todas as notas do sistema
func CarregaRotas() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/analise", controllers.Analise)
}
