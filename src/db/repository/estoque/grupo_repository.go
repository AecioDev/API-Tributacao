package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type GrupoRepository struct {
	base.RepositoryBase[estoque.Grupos]
}

func NewGrupoRepository(db *gorm.DB, cfgDb *config.AppConfig) *GrupoRepository {
	return &GrupoRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.Grupos](db, cfgDb),
	}
}

func (r *GrupoRepository) BuscaPeloNomeCompleto(name string) (*estoque.Grupos, error) {
	var grupo estoque.Grupos
	err := r.Db.Where("NomeCompleto = ?", name).First(&grupo).Error
	if err != nil {
		return nil, err
	}
	return &grupo, nil
}
