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

type PrecosProdutoController struct {
	precosprodutoService estserv.PrecosProdutoService
}

func NewPrecosProdutoController(service estserv.PrecosProdutoService) *PrecosProdutoController {
	return &PrecosProdutoController{
		precosprodutoService: service,
	}
}

// GetPrecosProdutos - Obtém a lista de precosprodutos
func (pc *PrecosProdutoController) GetPrecosProdutos(ctx *gin.Context) {

	precosprodutos, err := pc.precosprodutoService.GetAllPrecosProdutos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, precosprodutos)
}

// GetPrecosProdutoByID - Obtém um precosproduto pelo ID
func (pc *PrecosProdutoController) GetPrecosProdutoByID(ctx *gin.Context) {
	precosprodutoId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	precosproduto, err := pc.precosprodutoService.GetPrecosProdutoByID(precosprodutoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, precosproduto)
}

// GetPrecosProdutoByID - Obtém um precosproduto pelo Nome
func (pc *PrecosProdutoController) GetPrecosProdutosByName(ctx *gin.Context) {
	precosprodutoName := ctx.Param("name")
	if precosprodutoName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para PrecosProduto não foi informado!"}`)
	}
	precosproduto, err := pc.precosprodutoService.GetPrecosProdutosByName(precosprodutoName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, precosproduto)
}

// CreatePrecosProduto - Cria um novo precosproduto
func (pc *PrecosProdutoController) CreatePrecosProduto(ctx *gin.Context) {
	var precosproduto estoque.PrecosProdutos
	if err := services.BindJsonOrSendErrorRes(ctx, &precosproduto); err != nil {
		return
	}

	if err := pc.precosprodutoService.CreatePrecosProduto(&precosproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, precosproduto)
}

// UpdatePrecosProduto - Atualiza um precosproduto existente
func (pc *PrecosProdutoController) UpdatePrecosProduto(ctx *gin.Context) {
	var precosproduto estoque.PrecosProdutos
	if err := services.BindJsonOrSendErrorRes(ctx, &precosproduto); err != nil {
		return
	}

	if err := pc.precosprodutoService.UpdatePrecosProduto(&precosproduto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, precosproduto)
}

// DeletePrecosProduto - Exclui um precosproduto
func (pc *PrecosProdutoController) DeletePrecosProduto(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingPrecosProduto, err := pc.precosprodutoService.GetPrecosProdutoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o precosproduto: %d", id))
		return
	}

	if existingPrecosProduto == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("precosproduto com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.precosprodutoService.DeletePrecosProduto(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "PrecosProduto Deletado com Sucesso!!!"})
	}

}
