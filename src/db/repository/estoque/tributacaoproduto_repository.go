package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type TributacaoProdutoRepository struct {
	base.RepositoryBase[estoque.TributacaoProdutos]
}

func NewTributacaoProdutoRepository(db *gorm.DB, cfgDb *config.AppConfig) *TributacaoProdutoRepository {
	return &TributacaoProdutoRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.TributacaoProdutos](db, cfgDb),
	}
}

func (r *TributacaoProdutoRepository) BuscaPeloNomeCompleto(name string) (*estoque.TributacaoProdutos, error) {
	var tributacaoproduto estoque.TributacaoProdutos
	err := r.Db.Where("NomeCompleto = ?", name).First(&tributacaoproduto).Error
	if err != nil {
		return nil, err
	}
	return &tributacaoproduto, nil
}
