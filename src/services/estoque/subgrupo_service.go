package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type SubGrupoService struct {
	subgrupoRepo *repository.SubGrupoRepository
}

func NewSubGrupoService(repo *repository.SubGrupoRepository) *SubGrupoService {
	return &SubGrupoService{
		subgrupoRepo: repo,
	}
}

// Método para buscar todos os subgrupos
func (ps *SubGrupoService) GetAllSubGrupos() ([]estoque.SubGrupos, error) {
	return ps.subgrupoRepo.FindAll()
}

// Método para buscar por ID
func (ps *SubGrupoService) GetSubGrupoByID(id uint) (*estoque.SubGrupos, error) {
	return ps.subgrupoRepo.FindByID(id)
}

// Método para buscar subgrupo por nome
func (ps *SubGrupoService) GetSubGruposByName(name string) ([]estoque.SubGrupos, error) {
	return ps.subgrupoRepo.FindByName(name)
}

func (ps *SubGrupoService) CreateSubGrupo(subgrupo *estoque.SubGrupos) error {

	err := ps.subgrupoRepo.Create(subgrupo)
	if err != nil {
		return err
	}

	return nil
}

func (ps *SubGrupoService) UpdateSubGrupo(subgrupo *estoque.SubGrupos) error {

	err := ps.subgrupoRepo.Update(subgrupo)
	if err != nil {
		return err
	}

	return nil
}

func (ps *SubGrupoService) DeleteSubGrupo(subgrupoId uint) error {

	err := ps.subgrupoRepo.Delete(subgrupoId)
	if err != nil {
		return err
	}

	return nil
}
