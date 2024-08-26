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

type CustoProdutoController struct {
	custoprodutoService estserv.CustoProdutoService
}

func NewCustoProdutoController(service estserv.CustoProdutoService) *CustoProdutoController {
	return &CustoProdutoController{
		custoprodutoService: service,
	}
}

// GetCustoProdutos - Obtém a lista de custoprodutos
func (pc *CustoProdutoController) GetCustoProdutos(ctx *gin.Context) {

	custoprodutos, err := pc.custoprodutoService.GetAllCustoProdutos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, custoprodutos)
}

// GetCustoProdutoByID - Obtém um custoproduto pelo ID
func (pc *CustoProdutoController) GetCustoProdutoByID(ctx *gin.Context) {
	custoprodutoId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	custoproduto, err := pc.custoprodutoService.GetCustoProdutoByID(custoprodutoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, custoproduto)
}

// GetCustoProdutoByID - Obtém um custoproduto pelo Nome
func (pc *CustoProdutoController) GetCustoProdutosByName(ctx *gin.Context) {
	custoprodutoName := ctx.Param("name")
	if custoprodutoName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para CustoProduto não foi informado!"}`)
	}
	custoproduto, err := pc.custoprodutoService.GetCustoProdutosByName(custoprodutoName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, custoproduto)
}

// CreateCustoProduto - Cria um novo custoproduto
func (pc *CustoProdutoController) CreateCustoProduto(ctx *gin.Context) {
	var custoproduto estoque.CustoProdutos
	if err := services.BindJsonOrSendErrorRes(ctx, &custoproduto); err != nil {
		return
	}

	if err := pc.custoprodutoService.CreateCustoProduto(&custoproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, custoproduto)
}

// UpdateCustoProduto - Atualiza um custoproduto existente
func (pc *CustoProdutoController) UpdateCustoProduto(ctx *gin.Context) {
	var custoproduto estoque.CustoProdutos
	if err := services.BindJsonOrSendErrorRes(ctx, &custoproduto); err != nil {
		return
	}

	if err := pc.custoprodutoService.UpdateCustoProduto(&custoproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, custoproduto)
}

// DeleteCustoProduto - Exclui um custoproduto
func (pc *CustoProdutoController) DeleteCustoProduto(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingCustoProduto, err := pc.custoprodutoService.GetCustoProdutoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o custoproduto: %d", id))
		return
	}

	if existingCustoProduto == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("custoproduto com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.custoprodutoService.DeleteCustoProduto(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "CustoProduto Deletado com Sucesso!!!"})
	}

}
