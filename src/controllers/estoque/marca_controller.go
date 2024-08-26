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

type MarcaController struct {
	marcaService estserv.MarcaService
}

func NewMarcaController(service estserv.MarcaService) *MarcaController {
	return &MarcaController{
		marcaService: service,
	}
}

// GetMarcas - Obtém a lista de marcas
func (pc *MarcaController) GetMarcas(ctx *gin.Context) {

	marcas, err := pc.marcaService.GetAllMarcas()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, marcas)
}

// GetMarcaByID - Obtém um marca pelo ID
func (pc *MarcaController) GetMarcaByID(ctx *gin.Context) {
	marcaId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	marca, err := pc.marcaService.GetMarcaByID(marcaId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, marca)
}

// GetMarcaByID - Obtém um marca pelo Nome
func (pc *MarcaController) GetMarcasByName(ctx *gin.Context) {
	marcaName := ctx.Param("name")
	if marcaName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para Marca não foi informado!"}`)
	}
	marca, err := pc.marcaService.GetMarcasByName(marcaName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, marca)
}

// CreateMarca - Cria um novo marca
func (pc *MarcaController) CreateMarca(ctx *gin.Context) {
	var marca estoque.Marcas
	if err := services.BindJsonOrSendErrorRes(ctx, &marca); err != nil {
		return
	}

	if err := pc.marcaService.CreateMarca(&marca); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, marca)
}

// UpdateMarca - Atualiza um marca existente
func (pc *MarcaController) UpdateMarca(ctx *gin.Context) {
	var marca estoque.Marcas
	if err := services.BindJsonOrSendErrorRes(ctx, &marca); err != nil {
		return
	}

	if err := pc.marcaService.UpdateMarca(&marca); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, marca)
}

// DeleteMarca - Exclui um marca
func (pc *MarcaController) DeleteMarca(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingMarca, err := pc.marcaService.GetMarcaByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o marca: %d", id))
		return
	}

	if existingMarca == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("marca com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.marcaService.DeleteMarca(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "Marca Deletado com Sucesso!!!"})
	}

}
