package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type MarcaService struct {
	marcaRepo *repository.MarcaRepository
}

func NewMarcaService(repo *repository.MarcaRepository) *MarcaService {
	return &MarcaService{
		marcaRepo: repo,
	}
}

// Método para buscar todos os marcas
func (ps *MarcaService) GetAllMarcas() ([]estoque.Marcas, error) {
	return ps.marcaRepo.FindAll()
}

// Método para buscar por ID
func (ps *MarcaService) GetMarcaByID(id uint) (*estoque.Marcas, error) {
	return ps.marcaRepo.FindByID(id)
}

// Método para buscar marca por nome
func (ps *MarcaService) GetMarcasByName(name string) ([]estoque.Marcas, error) {
	return ps.marcaRepo.FindByName(name)
}

func (ps *MarcaService) CreateMarca(marca *estoque.Marcas) error {

	err := ps.marcaRepo.Create(marca)
	if err != nil {
		return err
	}

	return nil
}

func (ps *MarcaService) UpdateMarca(marca *estoque.Marcas) error {

	err := ps.marcaRepo.Update(marca)
	if err != nil {
		return err
	}

	return nil
}

func (ps *MarcaService) DeleteMarca(marcaId uint) error {

	err := ps.marcaRepo.Delete(marcaId)
	if err != nil {
		return err
	}

	return nil
}
