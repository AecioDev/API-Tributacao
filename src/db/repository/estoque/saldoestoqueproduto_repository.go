package repository

import (
	"api-tributacao/config"
	"api-tributacao/src/db/models/estoque"
	"api-tributacao/src/db/repository/base"

	"gorm.io/gorm"
)

type SaldoEstoqueProdutoRepository struct {
	base.RepositoryBase[estoque.SaldoEstoqueProduto]
}

func NewSaldoEstoqueProdutoRepository(db *gorm.DB, cfgDb *config.AppConfig) *SaldoEstoqueProdutoRepository {
	return &SaldoEstoqueProdutoRepository{
		RepositoryBase: *base.NewBaseRepository[estoque.SaldoEstoqueProduto](db, cfgDb),
	}
}

func (r *SaldoEstoqueProdutoRepository) BuscaPeloNomeCompleto(name string) (*estoque.SaldoEstoqueProduto, error) {
	var saldoestoqueproduto estoque.SaldoEstoqueProduto
	err := r.Db.Where("NomeCompleto = ?", name).First(&saldoestoqueproduto).Error
	if err != nil {
		return nil, err
	}
	return &saldoestoqueproduto, nil
}
