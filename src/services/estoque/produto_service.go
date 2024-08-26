package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type ProdutoService struct {
	produtoRepo *repository.ProdutoRepository
}

func NewProdutoService(repo *repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{
		produtoRepo: repo,
	}
}

// Método para buscar todos os produtos
func (ps *ProdutoService) GetAllProdutos() ([]estoque.Produtos, error) {
	return ps.produtoRepo.FindAll()
}

// Método para buscar por ID
func (ps *ProdutoService) GetProdutoByID(id uint) (*estoque.Produtos, error) {
	return ps.produtoRepo.FindByID(id)
}

// Método para buscar produto por nome
func (ps *ProdutoService) GetProdutosByName(name string) ([]estoque.Produtos, error) {
	return ps.produtoRepo.FindByName(name)
}

func (ps *ProdutoService) CreateProduto(produto *estoque.Produtos) error {

	err := ps.produtoRepo.Create(produto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProdutoService) UpdateProduto(produto *estoque.Produtos) error {

	err := ps.produtoRepo.Update(produto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProdutoService) DeleteProduto(produtoId uint) error {

	err := ps.produtoRepo.Delete(produtoId)
	if err != nil {
		return err
	}

	return nil
}
