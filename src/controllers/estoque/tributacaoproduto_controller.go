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

type TributacaoProdutoController struct {
	tributacaoprodutoService estserv.TributacaoProdutoService
}

func NewTributacaoProdutoController(service estserv.TributacaoProdutoService) *TributacaoProdutoController {
	return &TributacaoProdutoController{
		tributacaoprodutoService: service,
	}
}

// GetTributacaoProdutos - Obtém a lista de tributacaoprodutos
func (pc *TributacaoProdutoController) GetTributacaoProdutos(ctx *gin.Context) {

	tributacaoprodutos, err := pc.tributacaoprodutoService.GetAllTributacaoProdutos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tributacaoprodutos)
}

// GetTributacaoProdutoByID - Obtém um tributacaoproduto pelo ID
func (pc *TributacaoProdutoController) GetTributacaoProdutoByID(ctx *gin.Context) {
	tributacaoprodutoId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tributacaoproduto, err := pc.tributacaoprodutoService.GetTributacaoProdutoByID(tributacaoprodutoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tributacaoproduto)
}

// GetTributacaoProdutoByID - Obtém um tributacaoproduto pelo Nome
func (pc *TributacaoProdutoController) GetTributacaoProdutosByName(ctx *gin.Context) {
	tributacaoprodutoName := ctx.Param("name")
	if tributacaoprodutoName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para TributacaoProduto não foi informado!"}`)
	}
	tributacaoproduto, err := pc.tributacaoprodutoService.GetTributacaoProdutosByName(tributacaoprodutoName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tributacaoproduto)
}

// CreateTributacaoProduto - Cria um novo tributacaoproduto
func (pc *TributacaoProdutoController) CreateTributacaoProduto(ctx *gin.Context) {
	var tributacaoproduto estoque.TributacaoProdutos
	if err := services.BindJsonOrSendErrorRes(ctx, &tributacaoproduto); err != nil {
		return
	}

	if err := pc.tributacaoprodutoService.CreateTributacaoProduto(&tributacaoproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, tributacaoproduto)
}

// UpdateTributacaoProduto - Atualiza um tributacaoproduto existente
func (pc *TributacaoProdutoController) UpdateTributacaoProduto(ctx *gin.Context) {
	var tributacaoproduto estoque.TributacaoProdutos
	if err := services.BindJsonOrSendErrorRes(ctx, &tributacaoproduto); err != nil {
		return
	}

	if err := pc.tributacaoprodutoService.UpdateTributacaoProduto(&tributacaoproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tributacaoproduto)
}

// DeleteTributacaoProduto - Exclui um tributacaoproduto
func (pc *TributacaoProdutoController) DeleteTributacaoProduto(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingTributacaoProduto, err := pc.tributacaoprodutoService.GetTributacaoProdutoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o tributacaoproduto: %d", id))
		return
	}

	if existingTributacaoProduto == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("tributacaoproduto com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.tributacaoprodutoService.DeleteTributacaoProduto(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "TributacaoProduto Deletado com Sucesso!!!"})
	}

}
