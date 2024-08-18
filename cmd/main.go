package main

import (
	"api-tributacao/controller"
	"api-tributacao/db"
	"api-tributacao/repository"
	"api-tributacao/services"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de Repository
	ProdutoRepository := repository.NewProductRepository(dbConnection)

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
