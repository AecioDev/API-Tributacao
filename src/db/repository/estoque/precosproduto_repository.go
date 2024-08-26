package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type PrecosProdutoRepository struct {
	base.RepositoryBase[estoque.PrecosProdutos]
}

func NewPrecosProdutoRepository(db *gorm.DB, cfgDb *config.AppConfig) *PrecosProdutoRepository {
	return &PrecosProdutoRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.PrecosProdutos](db, cfgDb),
	}
}

func (r *PrecosProdutoRepository) BuscaPeloNomeCompleto(name string) (*estoque.PrecosProdutos, error) {
	var precosproduto estoque.PrecosProdutos
	err := r.Db.Where("NomeCompleto = ?", name).First(&precosproduto).Error
	if err != nil {
		return nil, err
	}
	return &precosproduto, nil
}
