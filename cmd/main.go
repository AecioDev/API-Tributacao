package main

import (
	"api-tributacao/config"
	"api-tributacao/src/controller"
	"api-tributacao/src/db"
	"api-tributacao/src/db/repository"
	"api-tributacao/src/globals"
	"api-tributacao/src/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalln("Failed to parse configuration: ", err)
	}

	database := db.New(cfg)
	defer db.CloseConnection(database)

	globals.SetDev(true)

	// Inicialização dos Repositórios
	produtoRepo := repository.NewProdutoRepository(database, &cfg.App)

	// Inicialização dos Services
	produtoService := services.NewProdutoService(produtoRepo)

	// Inicialização dos Controllers
	produtoController := controller.NewProdutoController(*produtoService)

	// Configuração das Rotas
	r := gin.Default()
	r.GET("/produtos", produtoController.GetProdutos)
	r.Run(":8080")

}
