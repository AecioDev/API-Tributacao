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

type UnidadeController struct {
	unidadeService estserv.UnidadeService
}

func NewUnidadeController(service estserv.UnidadeService) *UnidadeController {
	return &UnidadeController{
		unidadeService: service,
	}
}

// GetUnidades - Obtém a lista de unidades
func (pc *UnidadeController) GetUnidades(ctx *gin.Context) {

	unidades, err := pc.unidadeService.GetAllUnidades()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, unidades)
}

// GetUnidadeByID - Obtém um unidade pelo ID
func (pc *UnidadeController) GetUnidadeByID(ctx *gin.Context) {
	unidadeId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	unidade, err := pc.unidadeService.GetUnidadeByID(unidadeId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, unidade)
}

// GetUnidadeByID - Obtém um unidade pelo Nome
func (pc *UnidadeController) GetUnidadesByName(ctx *gin.Context) {
	unidadeName := ctx.Param("name")
	if unidadeName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para Unidade não foi informado!"}`)
	}
	unidade, err := pc.unidadeService.GetUnidadesByName(unidadeName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, unidade)
}

// CreateUnidade - Cria um novo unidade
func (pc *UnidadeController) CreateUnidade(ctx *gin.Context) {
	var unidade estoque.Unidades
	if err := services.BindJsonOrSendErrorRes(ctx, &unidade); err != nil {
		return
	}

	if err := pc.unidadeService.CreateUnidade(&unidade); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, unidade)
}

// UpdateUnidade - Atualiza um unidade existente
func (pc *UnidadeController) UpdateUnidade(ctx *gin.Context) {
	var unidade estoque.Unidades
	if err := services.BindJsonOrSendErrorRes(ctx, &unidade); err != nil {
		return
	}

	if err := pc.unidadeService.UpdateUnidade(&unidade); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, unidade)
}

// DeleteUnidade - Exclui um unidade
func (pc *UnidadeController) DeleteUnidade(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingUnidade, err := pc.unidadeService.GetUnidadeByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o unidade: %d", id))
		return
	}

	if existingUnidade == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("unidade com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.unidadeService.DeleteUnidade(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "Unidade Deletado com Sucesso!!!"})
	}

}
