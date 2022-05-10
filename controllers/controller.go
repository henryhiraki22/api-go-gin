package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henryhiraki22/api-go-gin/database"
	"github.com/henryhiraki22/api-go-gin/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	err := c.ShouldBindJSON(&aluno)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(200, aluno)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})

}
