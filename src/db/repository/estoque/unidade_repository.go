package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type UnidadeRepository struct {
	base.RepositoryBase[estoque.Unidades]
}

func NewUnidadeRepository(db *gorm.DB, cfgDb *config.AppConfig) *UnidadeRepository {
	return &UnidadeRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.Unidades](db, cfgDb),
	}
}
