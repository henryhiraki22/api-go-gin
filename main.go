package main

import (
	"github.com/henryhiraki22/api-go-gin/database"
	"github.com/henryhiraki22/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
