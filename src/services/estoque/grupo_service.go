package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type GrupoService struct {
	grupoRepo *repository.GrupoRepository
}

func NewGrupoService(repo *repository.GrupoRepository) *GrupoService {
	return &GrupoService{
		grupoRepo: repo,
	}
}

// Método para buscar todos os grupos
func (ps *GrupoService) GetAllGrupos() ([]estoque.Grupos, error) {
	return ps.grupoRepo.FindAll()
}

// Método para buscar por ID
func (ps *GrupoService) GetGrupoByID(id uint) (*estoque.Grupos, error) {
	return ps.grupoRepo.FindByID(id)
}

// Método para buscar grupo por nome
func (ps *GrupoService) GetGruposByName(name string) ([]estoque.Grupos, error) {
	return ps.grupoRepo.FindByName(name)
}

func (ps *GrupoService) CreateGrupo(grupo *estoque.Grupos) error {

	err := ps.grupoRepo.Create(grupo)
	if err != nil {
		return err
	}

	return nil
}

func (ps *GrupoService) UpdateGrupo(grupo *estoque.Grupos) error {

	err := ps.grupoRepo.Update(grupo)
	if err != nil {
		return err
	}

	return nil
}

func (ps *GrupoService) DeleteGrupo(grupoId uint) error {

	err := ps.grupoRepo.Delete(grupoId)
	if err != nil {
		return err
	}

	return nil
}
