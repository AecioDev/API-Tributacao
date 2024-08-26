package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type MovimentosEstoqueRepository struct {
	base.RepositoryBase[estoque.MovimentosEstoque]
}

func NewMovimentosEstoqueRepository(db *gorm.DB, cfgDb *config.AppConfig) *MovimentosEstoqueRepository {
	return &MovimentosEstoqueRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.MovimentosEstoque](db, cfgDb),
	}
}

func (r *MovimentosEstoqueRepository) BuscaPeloNomeCompleto(name string) (*estoque.MovimentosEstoque, error) {
	var movimentosestoque estoque.MovimentosEstoque
	err := r.Db.Where("NomeCompleto = ?", name).First(&movimentosestoque).Error
	if err != nil {
		return nil, err
	}
	return &movimentosestoque, nil
}
