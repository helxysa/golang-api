package main

import (
	controller "go-api-catalog/controller"
	"go-api-catalog/database"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	// Connect to the database
	db := database.ConnectDB()

	SolucosController := controller.NewSolucaoController(db)

	server.GET("/teste", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	server.GET("/solucoes", SolucosController.GetSolucoes)
	server.POST("/solucoes", SolucosController.CreateSolucao)
	server.PUT("/solucoes/:id", SolucosController.UpdateSolucao)
	server.DELETE("/solucoes/:id", SolucosController.DeleteSolucao)

	server.Run(":8000")
}
