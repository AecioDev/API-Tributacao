package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type PrecosProdutoService struct {
	precosprodutoRepo *repository.PrecosProdutoRepository
}

func NewPrecosProdutoService(repo *repository.PrecosProdutoRepository) *PrecosProdutoService {
	return &PrecosProdutoService{
		precosprodutoRepo: repo,
	}
}

// Método para buscar todos os precosprodutos
func (ps *PrecosProdutoService) GetAllPrecosProdutos() ([]estoque.PrecosProdutos, error) {
	return ps.precosprodutoRepo.FindAll()
}

// Método para buscar por ID
func (ps *PrecosProdutoService) GetPrecosProdutoByID(id uint) (*estoque.PrecosProdutos, error) {
	return ps.precosprodutoRepo.FindByID(id)
}

// Método para buscar precosproduto por nome
func (ps *PrecosProdutoService) GetPrecosProdutosByName(name string) ([]estoque.PrecosProdutos, error) {
	return ps.precosprodutoRepo.FindByName(name)
}

func (ps *PrecosProdutoService) CreatePrecosProduto(precosproduto *estoque.PrecosProdutos) error {

	err := ps.precosprodutoRepo.Create(precosproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PrecosProdutoService) UpdatePrecosProduto(precosproduto *estoque.PrecosProdutos) error {

	err := ps.precosprodutoRepo.Update(precosproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PrecosProdutoService) DeletePrecosProduto(precosprodutoId uint) error {

	err := ps.precosprodutoRepo.Delete(precosprodutoId)
	if err != nil {
		return err
	}

	return nil
}
