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

type MovimentosEstoqueController struct {
	movimentosestoqueService estserv.MovimentosEstoqueService
}

func NewMovimentosEstoqueController(service estserv.MovimentosEstoqueService) *MovimentosEstoqueController {
	return &MovimentosEstoqueController{
		movimentosestoqueService: service,
	}
}

// GetMovimentosEstoques - Obtém a lista de movimentosestoques
func (pc *MovimentosEstoqueController) GetMovimentosEstoques(ctx *gin.Context) {

	movimentosestoques, err := pc.movimentosestoqueService.GetAllMovimentosEstoques()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movimentosestoques)
}

// GetMovimentosEstoqueByID - Obtém um movimentosestoque pelo ID
func (pc *MovimentosEstoqueController) GetMovimentosEstoqueByID(ctx *gin.Context) {
	movimentosestoqueId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	movimentosestoque, err := pc.movimentosestoqueService.GetMovimentosEstoqueByID(movimentosestoqueId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movimentosestoque)
}

// GetMovimentosEstoqueByID - Obtém um movimentosestoque pelo Nome
func (pc *MovimentosEstoqueController) GetMovimentosEstoquesByName(ctx *gin.Context) {
	movimentosestoqueName := ctx.Param("name")
	if movimentosestoqueName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para MovimentosEstoque não foi informado!"}`)
	}
	movimentosestoque, err := pc.movimentosestoqueService.GetMovimentosEstoquesByName(movimentosestoqueName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movimentosestoque)
}

// CreateMovimentosEstoque - Cria um novo movimentosestoque
func (pc *MovimentosEstoqueController) CreateMovimentosEstoque(ctx *gin.Context) {
	var movimentosestoque estoque.MovimentosEstoque
	if err := services.BindJsonOrSendErrorRes(ctx, &movimentosestoque); err != nil {
		return
	}

	if err := pc.movimentosestoqueService.CreateMovimentosEstoque(&movimentosestoque); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, movimentosestoque)
}

// UpdateMovimentosEstoque - Atualiza um movimentosestoque existente
func (pc *MovimentosEstoqueController) UpdateMovimentosEstoque(ctx *gin.Context) {
	var movimentosestoque estoque.MovimentosEstoque
	if err := services.BindJsonOrSendErrorRes(ctx, &movimentosestoque); err != nil {
		return
	}

	if err := pc.movimentosestoqueService.UpdateMovimentosEstoque(&movimentosestoque); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movimentosestoque)
}

// DeleteMovimentosEstoque - Exclui um movimentosestoque
func (pc *MovimentosEstoqueController) DeleteMovimentosEstoque(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingMovimentosEstoque, err := pc.movimentosestoqueService.GetMovimentosEstoqueByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o movimentosestoque: %d", id))
		return
	}

	if existingMovimentosEstoque == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("movimentosestoque com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.movimentosestoqueService.DeleteMovimentosEstoque(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "MovimentosEstoque Deletado com Sucesso!!!"})
	}

}
