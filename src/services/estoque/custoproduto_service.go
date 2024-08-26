package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type CustoProdutoService struct {
	custoprodutoRepo *repository.CustoProdutoRepository
}

func NewCustoProdutoService(repo *repository.CustoProdutoRepository) *CustoProdutoService {
	return &CustoProdutoService{
		custoprodutoRepo: repo,
	}
}

// Método para buscar todos os custoprodutos
func (ps *CustoProdutoService) GetAllCustoProdutos() ([]estoque.CustoProdutos, error) {
	return ps.custoprodutoRepo.FindAll()
}

// Método para buscar por ID
func (ps *CustoProdutoService) GetCustoProdutoByID(id uint) (*estoque.CustoProdutos, error) {
	return ps.custoprodutoRepo.FindByID(id)
}

// Método para buscar custoproduto por nome
func (ps *CustoProdutoService) GetCustoProdutosByName(name string) ([]estoque.CustoProdutos, error) {
	return ps.custoprodutoRepo.FindByName(name)
}

func (ps *CustoProdutoService) CreateCustoProduto(custoproduto *estoque.CustoProdutos) error {

	err := ps.custoprodutoRepo.Create(custoproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *CustoProdutoService) UpdateCustoProduto(custoproduto *estoque.CustoProdutos) error {

	err := ps.custoprodutoRepo.Update(custoproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *CustoProdutoService) DeleteCustoProduto(custoprodutoId uint) error {

	err := ps.custoprodutoRepo.Delete(custoprodutoId)
	if err != nil {
		return err
	}

	return nil
}
