package controller

import (
	"api-tributacao/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type produtoController struct {
	produtoService services.ProdutoService
}

func NewProdutoController(service services.ProdutoService) produtoController {
	return produtoController{produtoService: service}
}

func (pc *produtoController) GetProdutos(ctx *gin.Context) {

	products, err := pc.produtoService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
