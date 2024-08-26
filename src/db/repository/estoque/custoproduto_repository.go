package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type CustoProdutoRepository struct {
	base.RepositoryBase[estoque.CustoProdutos]
}

func NewCustoProdutoRepository(db *gorm.DB, cfgDb *config.AppConfig) *CustoProdutoRepository {
	return &CustoProdutoRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.CustoProdutos](db, cfgDb),
	}
}

func (r *CustoProdutoRepository) BuscaPeloNomeCompleto(name string) (*estoque.CustoProdutos, error) {
	var custoproduto estoque.CustoProdutos
	err := r.Db.Where("NomeCompleto = ?", name).First(&custoproduto).Error
	if err != nil {
		return nil, err
	}
	return &custoproduto, nil
}
