package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type Server struct{}

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("arquivo .env não encontrado")
	}
}

func (s *Server) Run() {

	fmt.Println("Server rodando em localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(err)
	}
}
