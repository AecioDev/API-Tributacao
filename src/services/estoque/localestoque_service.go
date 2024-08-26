package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type LocalEstoqueService struct {
	localestoqueRepo *repository.LocalEstoqueRepository
}

func NewLocalEstoqueService(repo *repository.LocalEstoqueRepository) *LocalEstoqueService {
	return &LocalEstoqueService{
		localestoqueRepo: repo,
	}
}

// Método para buscar todos os localestoques
func (ps *LocalEstoqueService) GetAllLocalEstoques() ([]estoque.LocalEstoque, error) {
	return ps.localestoqueRepo.FindAll()
}

// Método para buscar por ID
func (ps *LocalEstoqueService) GetLocalEstoqueByID(id uint) (*estoque.LocalEstoque, error) {
	return ps.localestoqueRepo.FindByID(id)
}

// Método para buscar localestoque por nome
func (ps *LocalEstoqueService) GetLocalEstoquesByName(name string) ([]estoque.LocalEstoque, error) {
	return ps.localestoqueRepo.FindByName(name)
}

func (ps *LocalEstoqueService) CreateLocalEstoque(localestoque *estoque.LocalEstoque) error {

	err := ps.localestoqueRepo.Create(localestoque)
	if err != nil {
		return err
	}

	return nil
}

func (ps *LocalEstoqueService) UpdateLocalEstoque(localestoque *estoque.LocalEstoque) error {

	err := ps.localestoqueRepo.Update(localestoque)
	if err != nil {
		return err
	}

	return nil
}

func (ps *LocalEstoqueService) DeleteLocalEstoque(localestoqueId uint) error {

	err := ps.localestoqueRepo.Delete(localestoqueId)
	if err != nil {
		return err
	}

	return nil
}
