package controller

import (
	"api-tributacao/models"
	"api-tributacao/models/estoque"
	"api-tributacao/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type produtoController struct {
	//Services
	produtoservices services.ProdutoService
}

func NewProdutoController(services services.ProdutoService) produtoController {
	return produtoController{
		produtoservices: services,
	}
}

func (p *produtoController) GetProducts(ctx *gin.Context) {

	products, err := p.produtoservices.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *produtoController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		response := models.Response{
			Message: "Id do Produto não pode ser nulo!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := IdFromPathParamOrSendError(ctx)
	if err != nil {
		response := models.Response{
			Message: "Id do Produto precisa ser um número!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.produtoservices.GetProductById(productId)

	if err != nil {
		response := models.Response{
			Message: "Erro ao obter o produto: " + err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		response := models.Response{
			Message: "Produto não foi encontrado na base de dados!",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *produtoController) CreateProduct(ctx *gin.Context) {

	var product estoque.Produtos
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.produtoservices.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *produtoController) UpdateProduct(ctx *gin.Context) {

	var product estoque.Produtos
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updatedProduct, err := p.produtoservices.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *produtoController) DeleteProductById(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		response := models.Response{
			Message: "Id do Produto não pode ser nulo!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := IdFromPathParamOrSendError(ctx)
	if err != nil {
		response := models.Response{
			Message: "Id do Produto precisa ser um número!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.produtoservices.GetProductById(productId)

	if err != nil {
		response := models.Response{
			Message: "Erro ao obter o produto: " + err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		response := models.Response{
			Message: "Produto não foi encontrado na base de dados!",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	} else {

		result, err := p.produtoservices.DeleteProductById(productId)
		if err != nil {
			response := models.Response{
				Message: "Erro interno ao Deletar o Produto:\n" + err.Error(),
			}
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		response := models.Response{
			Message: result,
		}
		ctx.JSON(http.StatusOK, response)
	}
}

// busca e converte um uint de um parametro do path
func UintFromPathParam(ctx *gin.Context, paramName string) (uint, error) {
	paramUint, err := strconv.ParseUint(ctx.Param(paramName), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("path param: %s deve ser um integer positivo", paramName)
	}

	return uint(paramUint), nil
}

// busca um ID do path, respondendo com BadRequest no caso de erro
func IdFromPathParamOrSendError(ctx *gin.Context) (uint, error) {
	id, err := UintFromPathParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	return id, err
}
