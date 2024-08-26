package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	MIN_PAGESIZE = 1
	MAX_PAGESIZE = 1000
)

type PaginationQuery struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// tenta extrair parametros de paginação da query da requisição.
//
// caso inválido(s) retorna uma mensagem safe para retorno
func GetPaginationFromQuery(c *gin.Context) (PaginationQuery, error) {
	pagination := PaginationQuery{}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 0 {
		return pagination, errors.New("page number must be an positive number")
	}

	if page == 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		return pagination, errors.New("page size must be an integer")
	}

	if pageSize < MIN_PAGESIZE || pageSize > MAX_PAGESIZE {
		return pagination, errors.New("pageSize must be between " + strconv.Itoa(MIN_PAGESIZE) + " and " + strconv.Itoa(MAX_PAGESIZE))
	}

	pagination.Offset = (page - 1) * pageSize
	pagination.Limit = pageSize
	pagination.Page = page

	return pagination, nil
}

func GetPaginationFromQueryOrSendErrorRes(c *gin.Context) (PaginationQuery, error) {
	pagination, err := GetPaginationFromQuery(c)
	if err != nil {
		Res(c).StatusBadRequest().SendError(err)
	}

	return pagination, err
}
