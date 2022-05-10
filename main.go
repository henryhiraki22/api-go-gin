package main

import (
	"github.com/henryhiraki22/api-go-gin/models"
	"github.com/henryhiraki22/api-go-gin/routes"
)

type Alunos struct {
	id   string
	nome string
}

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Teste1", CPF: "123122311231", RG: "1231432423412"},
		{Nome: "Teste2", CPF: "123121115566", RG: "1231432423412"},
	}
	routes.HandleRequests()
}
