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

type SubGrupoController struct {
	subgrupoService estserv.SubGrupoService
}

func NewSubGrupoController(service estserv.SubGrupoService) *SubGrupoController {
	return &SubGrupoController{
		subgrupoService: service,
	}
}

// GetSubGrupos - Obtém a lista de subgrupos
func (pc *SubGrupoController) GetSubGrupos(ctx *gin.Context) {

	subgrupos, err := pc.subgrupoService.GetAllSubGrupos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subgrupos)
}

// GetSubGrupoByID - Obtém um subgrupo pelo ID
func (pc *SubGrupoController) GetSubGrupoByID(ctx *gin.Context) {
	subgrupoId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	subgrupo, err := pc.subgrupoService.GetSubGrupoByID(subgrupoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subgrupo)
}

// GetSubGrupoByID - Obtém um subgrupo pelo Nome
func (pc *SubGrupoController) GetSubGruposByName(ctx *gin.Context) {
	subgrupoName := ctx.Param("name")
	if subgrupoName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para SubGrupo não foi informado!"}`)
	}
	subgrupo, err := pc.subgrupoService.GetSubGruposByName(subgrupoName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subgrupo)
}

// CreateSubGrupo - Cria um novo subgrupo
func (pc *SubGrupoController) CreateSubGrupo(ctx *gin.Context) {
	var subgrupo estoque.SubGrupos
	if err := services.BindJsonOrSendErrorRes(ctx, &subgrupo); err != nil {
		return
	}

	if err := pc.subgrupoService.CreateSubGrupo(&subgrupo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, subgrupo)
}

// UpdateSubGrupo - Atualiza um subgrupo existente
func (pc *SubGrupoController) UpdateSubGrupo(ctx *gin.Context) {
	var subgrupo estoque.SubGrupos
	if err := services.BindJsonOrSendErrorRes(ctx, &subgrupo); err != nil {
		return
	}

	if err := pc.subgrupoService.UpdateSubGrupo(&subgrupo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, subgrupo)
}

// DeleteSubGrupo - Exclui um subgrupo
func (pc *SubGrupoController) DeleteSubGrupo(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingSubGrupo, err := pc.subgrupoService.GetSubGrupoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o subgrupo: %d", id))
		return
	}

	if existingSubGrupo == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("subgrupo com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.subgrupoService.DeleteSubGrupo(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "SubGrupo Deletado com Sucesso!!!"})
	}

}
