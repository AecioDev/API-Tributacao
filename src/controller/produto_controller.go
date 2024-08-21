package controller

import (
	"api-tributacao/src/services"
)

type produtoController struct {
	produtoService services.ProdutoService
}

func NewProdutoController(service services.ProdutoService) produtoController {
	return produtoController{produtoService: service}
}
