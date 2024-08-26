package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type TributacaoProdutoService struct {
	tributacaoprodutoRepo *repository.TributacaoProdutoRepository
}

func NewTributacaoProdutoService(repo *repository.TributacaoProdutoRepository) *TributacaoProdutoService {
	return &TributacaoProdutoService{
		tributacaoprodutoRepo: repo,
	}
}

// Método para buscar todos os tributacaoprodutos
func (ps *TributacaoProdutoService) GetAllTributacaoProdutos() ([]estoque.TributacaoProdutos, error) {
	return ps.tributacaoprodutoRepo.FindAll()
}

// Método para buscar por ID
func (ps *TributacaoProdutoService) GetTributacaoProdutoByID(id uint) (*estoque.TributacaoProdutos, error) {
	return ps.tributacaoprodutoRepo.FindByID(id)
}

// Método para buscar tributacaoproduto por nome
func (ps *TributacaoProdutoService) GetTributacaoProdutosByName(name string) ([]estoque.TributacaoProdutos, error) {
	return ps.tributacaoprodutoRepo.FindByName(name)
}

func (ps *TributacaoProdutoService) CreateTributacaoProduto(tributacaoproduto *estoque.TributacaoProdutos) error {

	err := ps.tributacaoprodutoRepo.Create(tributacaoproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *TributacaoProdutoService) UpdateTributacaoProduto(tributacaoproduto *estoque.TributacaoProdutos) error {

	err := ps.tributacaoprodutoRepo.Update(tributacaoproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *TributacaoProdutoService) DeleteTributacaoProduto(tributacaoprodutoId uint) error {

	err := ps.tributacaoprodutoRepo.Delete(tributacaoprodutoId)
	if err != nil {
		return err
	}

	return nil
}
