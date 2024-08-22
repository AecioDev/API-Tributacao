package repository

import (
	"api-tributacao/config"
	model "api-tributacao/src/db/models"

	"gorm.io/gorm"
)

type ProdutoRepository struct {
	BaseRepo RepositoryBase[model.Produtos]
}

func NewProdutoRepository(db *gorm.DB, cfgDb *config.AppConfig) *ProdutoRepository {
	return &ProdutoRepository{
		BaseRepo: *NewBaseRepository[model.Produtos](db, cfgDb),
	}
}

func (r *ProdutoRepository) FindByName(name string) (*model.Produtos, error) {
	var product model.Produtos
	err := r.BaseRepo.Db.Where("nome = ?", name).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
