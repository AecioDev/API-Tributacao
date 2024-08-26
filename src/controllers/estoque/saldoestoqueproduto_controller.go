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

type SaldoEstoqueProdutoController struct {
	saldoestoqueprodutoService estserv.SaldoEstoqueProdutoService
}

func NewSaldoEstoqueProdutoController(service estserv.SaldoEstoqueProdutoService) *SaldoEstoqueProdutoController {
	return &SaldoEstoqueProdutoController{
		saldoestoqueprodutoService: service,
	}
}

// GetSaldoEstoqueProdutos - Obtém a lista de saldoestoqueprodutos
func (pc *SaldoEstoqueProdutoController) GetSaldoEstoqueProdutos(ctx *gin.Context) {

	saldoestoqueprodutos, err := pc.saldoestoqueprodutoService.GetAllSaldoEstoqueProdutos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, saldoestoqueprodutos)
}

// GetSaldoEstoqueProdutoByID - Obtém um saldoestoqueproduto pelo ID
func (pc *SaldoEstoqueProdutoController) GetSaldoEstoqueProdutoByID(ctx *gin.Context) {
	saldoestoqueprodutoId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	saldoestoqueproduto, err := pc.saldoestoqueprodutoService.GetSaldoEstoqueProdutoByID(saldoestoqueprodutoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, saldoestoqueproduto)
}

// GetSaldoEstoqueProdutoByID - Obtém um saldoestoqueproduto pelo Nome
func (pc *SaldoEstoqueProdutoController) GetSaldoEstoqueProdutosByName(ctx *gin.Context) {
	saldoestoqueprodutoName := ctx.Param("name")
	if saldoestoqueprodutoName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para SaldoEstoqueProduto não foi informado!"}`)
	}
	saldoestoqueproduto, err := pc.saldoestoqueprodutoService.GetSaldoEstoqueProdutosByName(saldoestoqueprodutoName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, saldoestoqueproduto)
}

// CreateSaldoEstoqueProduto - Cria um novo saldoestoqueproduto
func (pc *SaldoEstoqueProdutoController) CreateSaldoEstoqueProduto(ctx *gin.Context) {
	var saldoestoqueproduto estoque.SaldoEstoqueProduto
	if err := services.BindJsonOrSendErrorRes(ctx, &saldoestoqueproduto); err != nil {
		return
	}

	if err := pc.saldoestoqueprodutoService.CreateSaldoEstoqueProduto(&saldoestoqueproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, saldoestoqueproduto)
}

// UpdateSaldoEstoqueProduto - Atualiza um saldoestoqueproduto existente
func (pc *SaldoEstoqueProdutoController) UpdateSaldoEstoqueProduto(ctx *gin.Context) {
	var saldoestoqueproduto estoque.SaldoEstoqueProduto
	if err := services.BindJsonOrSendErrorRes(ctx, &saldoestoqueproduto); err != nil {
		return
	}

	if err := pc.saldoestoqueprodutoService.UpdateSaldoEstoqueProduto(&saldoestoqueproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, saldoestoqueproduto)
}

// DeleteSaldoEstoqueProduto - Exclui um saldoestoqueproduto
func (pc *SaldoEstoqueProdutoController) DeleteSaldoEstoqueProduto(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingSaldoEstoqueProduto, err := pc.saldoestoqueprodutoService.GetSaldoEstoqueProdutoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o saldoestoqueproduto: %d", id))
		return
	}

	if existingSaldoEstoqueProduto == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("saldoestoqueproduto com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.saldoestoqueprodutoService.DeleteSaldoEstoqueProduto(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "SaldoEstoqueProduto Deletado com Sucesso!!!"})
	}

}
