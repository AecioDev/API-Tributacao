package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type CodigosProdutoService struct {
	codigosprodutoRepo *repository.CodigosProdutoRepository
}

func NewCodigosProdutoService(repo *repository.CodigosProdutoRepository) *CodigosProdutoService {
	return &CodigosProdutoService{
		codigosprodutoRepo: repo,
	}
}

// Método para buscar todos os codigosprodutos
func (ps *CodigosProdutoService) GetAllCodigosProdutos() ([]estoque.CodigosProduto, error) {
	return ps.codigosprodutoRepo.FindAll()
}

// Método para buscar por ID
func (ps *CodigosProdutoService) GetCodigosProdutoByID(id uint) (*estoque.CodigosProduto, error) {
	return ps.codigosprodutoRepo.FindByID(id)
}

// Método para buscar codigosproduto por nome
func (ps *CodigosProdutoService) GetCodigosProdutosByFornecedorId(id uint) ([]estoque.CodigosProduto, error) {
	return ps.codigosprodutoRepo.BuscaCodigosProdutosByFornecedorId(id)
}

func (ps *CodigosProdutoService) CreateCodigosProduto(codigosproduto *estoque.CodigosProduto) error {

	err := ps.codigosprodutoRepo.Create(codigosproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *CodigosProdutoService) UpdateCodigosProduto(codigosproduto *estoque.CodigosProduto) error {

	err := ps.codigosprodutoRepo.Update(codigosproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *CodigosProdutoService) DeleteCodigosProduto(codigosprodutoId uint) error {

	err := ps.codigosprodutoRepo.Delete(codigosprodutoId)
	if err != nil {
		return err
	}

	return nil
}
