package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type MarcaRepository struct {
	base.RepositoryBase[estoque.Marcas]
}

func NewMarcaRepository(db *gorm.DB, cfgDb *config.AppConfig) *MarcaRepository {
	return &MarcaRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.Marcas](db, cfgDb),
	}
}

func (r *MarcaRepository) BuscaPeloNomeCompleto(name string) (*estoque.Marcas, error) {
	var marca estoque.Marcas
	err := r.Db.Where("NomeCompleto = ?", name).First(&marca).Error
	if err != nil {
		return nil, err
	}
	return &marca, nil
}
