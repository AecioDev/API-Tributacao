package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type CodigosProdutoRepository struct {
	base.RepositoryBase[estoque.CodigosProduto]
}

func NewCodigosProdutoRepository(db *gorm.DB, cfgDb *config.AppConfig) *CodigosProdutoRepository {
	return &CodigosProdutoRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.CodigosProduto](db, cfgDb),
	}
}

func (r *CodigosProdutoRepository) BuscaCodigosProdutosByFornecedorId(id uint) ([]estoque.CodigosProduto, error) {
	var codigosproduto []estoque.CodigosProduto
	err := r.Db.Where("FornecedorID = ?", id).Find(&codigosproduto).Error
	if err != nil {
		return nil, err
	}
	return codigosproduto, nil
}
