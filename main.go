package main

import (
	"github.com/glaubergoncalves/analise/api"
	"github.com/glaubergoncalves/analise/api/routes"
)

func main() {
	routes.CarregaRotas()

	server := api.Server{}
	server.Run()
}
