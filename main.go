package main

import (
	"github.com/henryhiraki22/api-go-gin/database"
	"github.com/henryhiraki22/api-go-gin/routes"
)

type Alunos struct {
	id   string
	nome string
}

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
