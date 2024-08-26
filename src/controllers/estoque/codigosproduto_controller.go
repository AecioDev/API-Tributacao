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

type CodigosProdutoController struct {
	codigosprodutoService estserv.CodigosProdutoService
}

func NewCodigosProdutoController(service estserv.CodigosProdutoService) *CodigosProdutoController {
	return &CodigosProdutoController{
		codigosprodutoService: service,
	}
}

// GetCodigosProdutos - Obtém a lista de codigosprodutos
func (pc *CodigosProdutoController) GetCodigosProdutos(ctx *gin.Context) {

	codigosprodutos, err := pc.codigosprodutoService.GetAllCodigosProdutos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, codigosprodutos)
}

// GetCodigosProdutoByID - Obtém um codigosproduto pelo ID
func (pc *CodigosProdutoController) GetCodigosProdutoByID(ctx *gin.Context) {
	codigosprodutoId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	codigosproduto, err := pc.codigosprodutoService.GetCodigosProdutoByID(codigosprodutoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, codigosproduto)
}

// GetCodigosProdutoByID - Obtém um codigosproduto pelo Nome
func (pc *CodigosProdutoController) GetCodigosProdutosByFornecedorId(ctx *gin.Context) {
	fornecedorID, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	codigosproduto, err := pc.codigosprodutoService.GetCodigosProdutosByFornecedorId(fornecedorID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, codigosproduto)
}

// CreateCodigosProduto - Cria um novo codigosproduto
func (pc *CodigosProdutoController) CreateCodigosProduto(ctx *gin.Context) {
	var codigosproduto estoque.CodigosProduto
	if err := services.BindJsonOrSendErrorRes(ctx, &codigosproduto); err != nil {
		return
	}

	if err := pc.codigosprodutoService.CreateCodigosProduto(&codigosproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, codigosproduto)
}

// UpdateCodigosProduto - Atualiza um codigosproduto existente
func (pc *CodigosProdutoController) UpdateCodigosProduto(ctx *gin.Context) {
	var codigosproduto estoque.CodigosProduto
	if err := services.BindJsonOrSendErrorRes(ctx, &codigosproduto); err != nil {
		return
	}

	if err := pc.codigosprodutoService.UpdateCodigosProduto(&codigosproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, codigosproduto)
}

// DeleteCodigosProduto - Exclui um codigosproduto
func (pc *CodigosProdutoController) DeleteCodigosProduto(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingCodigosProduto, err := pc.codigosprodutoService.GetCodigosProdutoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o codigosproduto: %d", id))
		return
	}

	if existingCodigosProduto == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("codigosproduto com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.codigosprodutoService.DeleteCodigosProduto(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "CodigosProduto Deletado com Sucesso!!!"})
	}

}
