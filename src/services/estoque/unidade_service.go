package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type UnidadeService struct {
	unidadeRepo *repository.UnidadeRepository
}

func NewUnidadeService(repo *repository.UnidadeRepository) *UnidadeService {
	return &UnidadeService{
		unidadeRepo: repo,
	}
}

// Método para buscar todos os unidades
func (ps *UnidadeService) GetAllUnidades() ([]estoque.Unidades, error) {
	return ps.unidadeRepo.FindAll()
}

// Método para buscar por ID
func (ps *UnidadeService) GetUnidadeByID(id uint) (*estoque.Unidades, error) {
	return ps.unidadeRepo.FindByID(id)
}

// Método para buscar unidade por nome
func (ps *UnidadeService) GetUnidadesByName(name string) ([]estoque.Unidades, error) {
	return ps.unidadeRepo.FindByName(name)
}

func (ps *UnidadeService) CreateUnidade(unidade *estoque.Unidades) error {

	err := ps.unidadeRepo.Create(unidade)
	if err != nil {
		return err
	}

	return nil
}

func (ps *UnidadeService) UpdateUnidade(unidade *estoque.Unidades) error {

	err := ps.unidadeRepo.Update(unidade)
	if err != nil {
		return err
	}

	return nil
}

func (ps *UnidadeService) DeleteUnidade(unidadeId uint) error {

	err := ps.unidadeRepo.Delete(unidadeId)
	if err != nil {
		return err
	}

	return nil
}
