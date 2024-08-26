package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type ProdutoRepository struct {
	base.RepositoryBase[estoque.Produtos]
}

func NewProdutoRepository(db *gorm.DB, cfgDb *config.AppConfig) *ProdutoRepository {
	return &ProdutoRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.Produtos](db, cfgDb),
	}
}

func (r *ProdutoRepository) BuscaPeloNomeCompleto(name string) (*estoque.Produtos, error) {
	var product estoque.Produtos
	err := r.Db.Where("NomeCompleto = ?", name).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
