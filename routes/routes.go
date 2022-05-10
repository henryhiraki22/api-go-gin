package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henryhiraki22/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.POST("/delete-aluno/:id", controllers.DeletaAluno)
	r.POST("/alunos/:id", controllers.EditaAluno)
	r.Run(":8080")
}
