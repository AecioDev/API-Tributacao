package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type SubGrupoRepository struct {
	base.RepositoryBase[estoque.SubGrupos]
}

func NewSubGrupoRepository(db *gorm.DB, cfgDb *config.AppConfig) *SubGrupoRepository {
	return &SubGrupoRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.SubGrupos](db, cfgDb),
	}
}

func (r *SubGrupoRepository) BuscaPeloNomeCompleto(name string) (*estoque.SubGrupos, error) {
	var subgrupo estoque.SubGrupos
	err := r.Db.Where("NomeCompleto = ?", name).First(&subgrupo).Error
	if err != nil {
		return nil, err
	}
	return &subgrupo, nil
}
