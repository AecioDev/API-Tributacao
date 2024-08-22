package services

import (
	model "api-tributacao/src/db/models"
	"api-tributacao/src/db/repository"
)

type ProdutoService struct {
	ProdutoRepo *repository.ProdutoRepository
}

func NewProdutoService(repo *repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{ProdutoRepo: repo}
}

// Método para buscar todos os produtos
func (ps *ProdutoService) GetAll() ([]model.Produtos, error) {
	return ps.ProdutoRepo.BaseRepo.FindAll()
}

// Método para buscar por ID
func (ps *ProdutoService) GetById(id uint) (*model.Produtos, error) {
	return ps.ProdutoRepo.BaseRepo.FindByID(id)
}

// Método para buscar produto por nome
func (ps *ProdutoService) GetByName(name string) (*model.Produtos, error) {
	return ps.ProdutoRepo.FindByName(name)
}

func (ps *ProdutoService) Create(produto model.Produtos) (model.Produtos, error) {

	err := ps.ProdutoRepo.BaseRepo.Create(&produto)
	if err != nil {
		return model.Produtos{}, err
	}

	return produto, nil
}

func (ps *ProdutoService) Update(produto model.Produtos) (model.Produtos, error) {

	err := ps.ProdutoRepo.BaseRepo.Update(&produto)
	if err != nil {
		return model.Produtos{}, err
	}

	return produto, nil
}

func (ps *ProdutoService) Delete(produtoId uint) error {

	err := ps.ProdutoRepo.BaseRepo.Delete(produtoId)
	if err != nil {
		return err
	}

	return nil
}
