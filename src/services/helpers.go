package services

import (
	utils "api-tributacao/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context) *utils.GinResponse {
	return utils.Res(c)
}

func handleBindError(c *gin.Context, err error) error {
	if err != nil {
		utils.Res(c).StatusBadRequest().SendError(err)
	}

	return err
}

func BindJsonOrSendErrorRes(c *gin.Context, obj any) error {
	return handleBindError(c, c.ShouldBindJSON(obj))
}

func BindQueryOrSendErrorRes(c *gin.Context, obj any) error {
	return handleBindError(c, c.ShouldBindQuery(obj))
}

func BindUriOrSendErrorRes(c *gin.Context, obj any) error {
	return handleBindError(c, c.ShouldBindUri(obj))
}

// utilidade para se buscar uma entidade por ID,
// enviando erro 404 caso a entidade nâo exista
func getEntityById[T any](c *gin.Context, getter func(uint) (T, error)) *T {
	id, err := utils.IdFromPathParamOrSendError(c)
	if err != nil {
		return nil
	}

	entidade, err := getter(id)
	if err != nil {
		utils.Res(c).StatusNotFound().SendErrorMsg(fmt.Sprintf("entidade de ID: %d não existe", id))
		return nil
	}

	return &entidade
}
