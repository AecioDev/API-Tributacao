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

type GrupoController struct {
	grupoService estserv.GrupoService
}

func NewGrupoController(service estserv.GrupoService) *GrupoController {
	return &GrupoController{
		grupoService: service,
	}
}

// GetGrupos - Obtém a lista de grupos
func (pc *GrupoController) GetGrupos(ctx *gin.Context) {

	grupos, err := pc.grupoService.GetAllGrupos()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, grupos)
}

// GetGrupoByID - Obtém um grupo pelo ID
func (pc *GrupoController) GetGrupoByID(ctx *gin.Context) {
	grupoId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	grupo, err := pc.grupoService.GetGrupoByID(grupoId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, grupo)
}

// GetGrupoByID - Obtém um grupo pelo Nome
func (pc *GrupoController) GetGruposByName(ctx *gin.Context) {
	grupoName := ctx.Param("name")
	if grupoName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para Grupo não foi informado!"}`)
	}
	grupo, err := pc.grupoService.GetGruposByName(grupoName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, grupo)
}

// CreateGrupo - Cria um novo grupo
func (pc *GrupoController) CreateGrupo(ctx *gin.Context) {
	var grupo estoque.Grupos
	if err := services.BindJsonOrSendErrorRes(ctx, &grupo); err != nil {
		return
	}

	if err := pc.grupoService.CreateGrupo(&grupo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, grupo)
}

// UpdateGrupo - Atualiza um grupo existente
func (pc *GrupoController) UpdateGrupo(ctx *gin.Context) {
	var grupo estoque.Grupos
	if err := services.BindJsonOrSendErrorRes(ctx, &grupo); err != nil {
		return
	}

	if err := pc.grupoService.UpdateGrupo(&grupo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, grupo)
}

// DeleteGrupo - Exclui um grupo
func (pc *GrupoController) DeleteGrupo(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingGrupo, err := pc.grupoService.GetGrupoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o grupo: %d", id))
		return
	}

	if existingGrupo == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("grupo com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.grupoService.DeleteGrupo(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "Grupo Deletado com Sucesso!!!"})
	}

}
