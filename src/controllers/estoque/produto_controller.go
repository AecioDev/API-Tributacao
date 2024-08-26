package controller

import (
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/services"
	estserv "api-tributacao/src/services/estoque"
	"api-tributacao/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProdutoController struct {
	produtoService estserv.ProdutoService
}

func NewProdutoController(service estserv.ProdutoService) *ProdutoController {
	return &ProdutoController{
		produtoService: service,
	}
}

// GetProdutos - Obtém a lista de produtos
func (pc *ProdutoController) GetProdutos(ctx *gin.Context) {

	produtos, err := pc.produtoService.GetAllProdutos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, produtos)
}

// GetProdutoByID - Obtém um produto pelo ID
func (pc *ProdutoController) GetProdutoByID(ctx *gin.Context) {
	produtoId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	produto, err := pc.produtoService.GetProdutoByID(produtoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, produto)
}

// GetProdutoByID - Obtém um produto pelo Nome
func (pc *ProdutoController) GetProdutosByName(ctx *gin.Context) {
	produtoName := ctx.Param("name")
	if produtoName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome do Prouto não foi informado!"}`)
	}
	produto, err := pc.produtoService.GetProdutosByName(produtoName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, produto)
}

// CreateProduto - Cria um novo produto
func (pc *ProdutoController) CreateProduto(ctx *gin.Context) {
	var produto estoque.Produtos
	if err := services.BindJsonOrSendErrorRes(ctx, &produto); err != nil {
		return
	}

	if err := pc.produtoService.CreateProduto(&produto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, produto)
}

// UpdateProduto - Atualiza um produto existente
func (pc *ProdutoController) UpdateProduto(ctx *gin.Context) {
	var produto estoque.Produtos
	if err := services.BindJsonOrSendErrorRes(ctx, &produto); err != nil {
		return
	}

	if err := pc.produtoService.UpdateProduto(&produto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, produto)
}

// DeleteProduto - Exclui um produto
func (pc *ProdutoController) DeleteProduto(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingProduto, err := pc.produtoService.GetProdutoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o produto: %d", id))
		return
	}

	if existingProduto == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("produto com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.produtoService.DeleteProduto(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "Produto Deletado com Sucesso!!!"})
	}

}
