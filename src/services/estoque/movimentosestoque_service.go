package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type MovimentosEstoqueService struct {
	movimentosestoqueRepo *repository.MovimentosEstoqueRepository
}

func NewMovimentosEstoqueService(repo *repository.MovimentosEstoqueRepository) *MovimentosEstoqueService {
	return &MovimentosEstoqueService{
		movimentosestoqueRepo: repo,
	}
}

// Método para buscar todos os movimentosestoques
func (ps *MovimentosEstoqueService) GetAllMovimentosEstoques() ([]estoque.MovimentosEstoque, error) {
	return ps.movimentosestoqueRepo.FindAll()
}

// Método para buscar por ID
func (ps *MovimentosEstoqueService) GetMovimentosEstoqueByID(id uint) (*estoque.MovimentosEstoque, error) {
	return ps.movimentosestoqueRepo.FindByID(id)
}

// Método para buscar movimentosestoque por nome
func (ps *MovimentosEstoqueService) GetMovimentosEstoquesByName(name string) ([]estoque.MovimentosEstoque, error) {
	return ps.movimentosestoqueRepo.FindByName(name)
}

func (ps *MovimentosEstoqueService) CreateMovimentosEstoque(movimentosestoque *estoque.MovimentosEstoque) error {

	err := ps.movimentosestoqueRepo.Create(movimentosestoque)
	if err != nil {
		return err
	}

	return nil
}

func (ps *MovimentosEstoqueService) UpdateMovimentosEstoque(movimentosestoque *estoque.MovimentosEstoque) error {

	err := ps.movimentosestoqueRepo.Update(movimentosestoque)
	if err != nil {
		return err
	}

	return nil
}

func (ps *MovimentosEstoqueService) DeleteMovimentosEstoque(movimentosestoqueId uint) error {

	err := ps.movimentosestoqueRepo.Delete(movimentosestoqueId)
	if err != nil {
		return err
	}

	return nil
}
