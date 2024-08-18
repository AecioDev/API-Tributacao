package services

import (
	"api-tributacao/models/estoque"
	"api-tributacao/repository"
)

type ProdutoService struct {
	repository repository.ProdutoRepository
}

func NewProdutoService(repo repository.ProdutoRepository) ProdutoService {
	return ProdutoService{
		repository: repo,
	}
}

func (pu *ProdutoService) GetProducts() ([]estoque.Produtos, error) {
	return pu.repository.GetProducts()
}

func (pu *ProdutoService) GetProductById(productId uint) (*estoque.Produtos, error) {
	return pu.repository.GetProductById(productId)
}

func (pu *ProdutoService) CreateProduct(product estoque.Produtos) (estoque.Produtos, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return estoque.Produtos{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProdutoService) UpdateProduct(product estoque.Produtos) (estoque.Produtos, error) {
	return pu.repository.UpdateProduct(product)
}

func (pu *ProdutoService) DeleteProductById(productId uint) (string, error) {
	return pu.repository.DeleteProductById(productId)
}
