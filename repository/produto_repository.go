package repository

import (
	"api-tributacao/models/estoque"
	"fmt"

	"gorm.io/gorm"
)

type ProdutoRepository struct {
	connection *gorm.DB
}

func NewProductRepository(connection *gorm.DB) ProdutoRepository {
	return ProdutoRepository{
		connection: connection,
	}
}

func (pr *ProdutoRepository) GetProducts() ([]estoque.Produtos, error) {

	var productList []estoque.Produtos

	result := pr.connection.Find(&productList)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return []estoque.Produtos{}, result.Error
	}

	return productList, nil
}

func (pr *ProdutoRepository) GetProductById(produtoId uint) (*estoque.Produtos, error) {

	var productObj estoque.Produtos

	result := pr.connection.First(&productObj, produtoId)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return &estoque.Produtos{}, result.Error
	}

	return &productObj, nil
}

func (pr *ProdutoRepository) CreateProduct(product estoque.Produtos) (uint, error) {

	result := pr.connection.Create(&product)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return 0, result.Error
	}

	return product.ID, nil
}

func (pr *ProdutoRepository) UpdateProduct(product estoque.Produtos) (estoque.Produtos, error) {

	result := pr.connection.Save(&product)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return product, result.Error
	}

	return product, nil
}

func (pr *ProdutoRepository) DeleteProductById(productId uint) (string, error) {

	result := pr.connection.Delete(&estoque.Produtos{}, productId)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return "Erro ao Deletar o Produto:\n", result.Error
	}

	return "Produto Deletado com Sucesso!", nil
}
