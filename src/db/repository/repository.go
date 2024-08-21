package repository

import (
	"api-tributacao/config"

	"gorm.io/gorm"
)

// BaseRepository define o repositório base com operações CRUD usando generics
type RepositoryBase[T any] struct {
	Db     *gorm.DB
	appCfg *config.AppConfig
}

// NewBaseRepository cria uma nova instância do repositório base
func NewBaseRepository[T any](db *gorm.DB, appCfg *config.AppConfig) *RepositoryBase[T] {
	return &RepositoryBase[T]{Db: db, appCfg: appCfg}
}

// Busca todos os registros paginados
func (r *RepositoryBase[T]) FindAllPaginated(page, pageSize int) ([]T, error) {
	var entities []T
	offset := (page - 1) * pageSize
	result := r.Db.Limit(pageSize).Offset(offset).Find(&entities)
	return entities, result.Error
}

// GetdAll busca uma entidade por ID
func (r *RepositoryBase[T]) FindAll() ([]T, error) {
	var entities []T
	result := r.Db.Find(&entities)

	return entities, result.Error
}

// FindByID busca uma entidade por ID
func (r *RepositoryBase[T]) FindByID(id uint) (*T, error) {
	var entity T
	err := r.Db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// Create insere uma nova entidade no banco de dados
func (r *RepositoryBase[T]) Create(entity *T) error {
	return r.Db.Create(entity).Error
}

// Update atualiza uma entidade existente
func (r *RepositoryBase[T]) Update(entity *T) error {
	return r.Db.Save(entity).Error
}

// Delete remove uma entidade por ID
func (r *RepositoryBase[T]) Delete(id uint) error {
	return r.Db.Delete(new(T), id).Error
}
