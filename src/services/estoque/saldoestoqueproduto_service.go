package services

import (
	"api-tributacao/src/db/models/estoque"
	repository "api-tributacao/src/db/repository/estoque"
)

type SaldoEstoqueProdutoService struct {
	saldoestoqueprodutoRepo *repository.SaldoEstoqueProdutoRepository
}

func NewSaldoEstoqueProdutoService(repo *repository.SaldoEstoqueProdutoRepository) *SaldoEstoqueProdutoService {
	return &SaldoEstoqueProdutoService{
		saldoestoqueprodutoRepo: repo,
	}
}

// Método para buscar todos os saldoestoqueprodutos
func (ps *SaldoEstoqueProdutoService) GetAllSaldoEstoqueProdutos() ([]estoque.SaldoEstoqueProduto, error) {
	return ps.saldoestoqueprodutoRepo.FindAll()
}

// Método para buscar por ID
func (ps *SaldoEstoqueProdutoService) GetSaldoEstoqueProdutoByID(id uint) (*estoque.SaldoEstoqueProduto, error) {
	return ps.saldoestoqueprodutoRepo.FindByID(id)
}

// Método para buscar saldoestoqueproduto por nome
func (ps *SaldoEstoqueProdutoService) GetSaldoEstoqueProdutosByName(name string) ([]estoque.SaldoEstoqueProduto, error) {
	return ps.saldoestoqueprodutoRepo.FindByName(name)
}

func (ps *SaldoEstoqueProdutoService) CreateSaldoEstoqueProduto(saldoestoqueproduto *estoque.SaldoEstoqueProduto) error {

	err := ps.saldoestoqueprodutoRepo.Create(saldoestoqueproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *SaldoEstoqueProdutoService) UpdateSaldoEstoqueProduto(saldoestoqueproduto *estoque.SaldoEstoqueProduto) error {

	err := ps.saldoestoqueprodutoRepo.Update(saldoestoqueproduto)
	if err != nil {
		return err
	}

	return nil
}

func (ps *SaldoEstoqueProdutoService) DeleteSaldoEstoqueProduto(saldoestoqueprodutoId uint) error {

	err := ps.saldoestoqueprodutoRepo.Delete(saldoestoqueprodutoId)
	if err != nil {
		return err
	}

	return nil
}
