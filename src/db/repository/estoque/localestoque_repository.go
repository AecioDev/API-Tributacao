package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type LocalEstoqueRepository struct {
	base.RepositoryBase[estoque.LocalEstoque]
}

func NewLocalEstoqueRepository(db *gorm.DB, cfgDb *config.AppConfig) *LocalEstoqueRepository {
	return &LocalEstoqueRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.LocalEstoque](db, cfgDb),
	}
}

func (r *LocalEstoqueRepository) BuscaPeloNomeCompleto(name string) (*estoque.LocalEstoque, error) {
	var localestoque estoque.LocalEstoque
	err := r.Db.Where("NomeCompleto = ?", name).First(&localestoque).Error
	if err != nil {
		return nil, err
	}
	return &localestoque, nil
}
