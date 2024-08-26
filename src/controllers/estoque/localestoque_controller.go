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

type LocalEstoqueController struct {
	localestoqueService estserv.LocalEstoqueService
}

func NewLocalEstoqueController(service estserv.LocalEstoqueService) *LocalEstoqueController {
	return &LocalEstoqueController{
		localestoqueService: service,
	}
}

// GetLocalEstoques - Obtém a lista de localestoques
func (pc *LocalEstoqueController) GetLocalEstoques(ctx *gin.Context) {

	localestoques, err := pc.localestoqueService.GetAllLocalEstoques()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, localestoques)
}

// GetLocalEstoqueByID - Obtém um localestoque pelo ID
func (pc *LocalEstoqueController) GetLocalEstoqueByID(ctx *gin.Context) {
	localestoqueId, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	localestoque, err := pc.localestoqueService.GetLocalEstoqueByID(localestoqueId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, localestoque)
}

// GetLocalEstoqueByID - Obtém um localestoque pelo Nome
func (pc *LocalEstoqueController) GetLocalEstoquesByName(ctx *gin.Context) {
	localestoqueName := ctx.Param("name")
	if localestoqueName == "" {
		ctx.JSON(http.StatusBadRequest, `{error: "O Nome para LocalEstoque não foi informado!"}`)
	}
	localestoque, err := pc.localestoqueService.GetLocalEstoquesByName(localestoqueName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, localestoque)
}

// CreateLocalEstoque - Cria um novo localestoque
func (pc *LocalEstoqueController) CreateLocalEstoque(ctx *gin.Context) {
	var localestoque estoque.LocalEstoque
	if err := services.BindJsonOrSendErrorRes(ctx, &localestoque); err != nil {
		return
	}

	if err := pc.localestoqueService.CreateLocalEstoque(&localestoque); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, localestoque)
}

// UpdateLocalEstoque - Atualiza um localestoque existente
func (pc *LocalEstoqueController) UpdateLocalEstoque(ctx *gin.Context) {
	var localestoque estoque.LocalEstoque
	if err := services.BindJsonOrSendErrorRes(ctx, &localestoque); err != nil {
		return
	}

	if err := pc.localestoqueService.UpdateLocalEstoque(&localestoque); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, localestoque)
}

// DeleteLocalEstoque - Exclui um localestoque
func (pc *LocalEstoqueController) DeleteLocalEstoque(ctx *gin.Context) {
	id, err := utils.IdFromPathParamOrSendError(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingLocalEstoque, err := pc.localestoqueService.GetLocalEstoqueByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao obter o localestoque: %d", id))
		return
	}

	if existingLocalEstoque == nil {
		ctx.JSON(http.StatusNotFound, fmt.Errorf("localestoque com ID: %d não foi encontrado na base de dados", id))
		return
	} else {
		err := pc.localestoqueService.DeleteLocalEstoque(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": "LocalEstoque Deletado com Sucesso!!!"})
	}

}
