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
	produtoRepo := repository.ProdutoRepository.NewProdutoRepository(database, &cfg.App)

	// Inicialização dos Services
	produtoService := services.NewProdutoService(produtoRepo)

	// Inicialização dos Controllers
	produtoController := controller.NewProdutoController(produtoService)

	// Configuração das Rotas
	r := gin.Default()
	r.GET("/produtos", produtoController.GetAll)
	r.Run(":8080")

}

//------------------------------------------------
/*
func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de Repository
	ProdutoRepository := repository.BaseRepository // NewProductRepository(dbConnection)

	// Camada UseCase
	ProdutoService := services.NewProdutoService(ProdutoRepository)

	// Camada de Controllers
	ProdutoController := controller.NewProdutoController(ProdutoService)

	// Rotas

	// PRODUTOS
	server.GET("/products", ProdutoController.GetProducts)
	server.POST("/product", ProdutoController.CreateProduct)
	server.PUT("/product", ProdutoController.UpdateProduct)
	server.GET("/product/:id", ProdutoController.GetProductById)
	server.DELETE("/product/:id", ProdutoController.DeleteProductById)

	server.Run(":8000")

}
*/
