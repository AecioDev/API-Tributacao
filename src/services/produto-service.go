package services

import (
	model "api-tributacao/src/db/models"
	"api-tributacao/src/db/repository"
)

type ProdutoService struct {
	ProdutoRepo *repository.ProdutoRepository
}

func NewProdutoService(repo *repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{ProdutoRepo: repo}
}

// Método para buscar todos os produtos
func (ps *ProdutoService) GetAll() ([]model.Produtos, error) {
	return ps.ProdutoRepo.BaseRepo.FindAll()
}

// Método para buscar por ID
func (ps *ProdutoService) GetById(id uint) (*model.Produtos, error) {
	return ps.ProdutoRepo.BaseRepo.FindByID(id)
}

// Método para buscar produto por nome
func (ps *ProdutoService) GetByName(name string) (*model.Produtos, error) {
	return ps.ProdutoRepo.FindByName(name)
}

func (ps *ProdutoService) Create(produto model.Produtos) (model.Produtos, error) {

	err := ps.ProdutoRepo.BaseRepo.Create(&produto)
	if err != nil {
		return model.Produtos{}, err
	}

	return produto, nil
}

func (ps *ProdutoService) Update(produto model.Produtos) (model.Produtos, error) {

	return produto, nil
}

// Outros métodos específicos para Produto...

/*
type ProdutoService struct {
	*BaseServiceImpl[estoque.Produtos]
}

func NewProdutoService(repo *repository.RepositoryBase[estoque.Produtos]) *ProdutoService {
	return &ProdutoService{
		BaseServiceImpl: NewBaseService(repo),
		// Inicialize outras dependências específicas do ProdutoService
	}
}

// Exemplo de método específico para Produto
func (s *ProdutoService) GetByName(name string) (*estoque.Produtos, error) {
	//return s.repo.GetByName(name)
	return nil, nil
}
*/
