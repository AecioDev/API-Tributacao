package estoque

import "api-tributacao/models"

type CodigosProduto struct {
	models.InitEntity

	ProdutoID    uint   `gorm:"not null"` // Chave estrangeira para Produto
	FornecedorID uint   `gorm:"not null"` // Código do Fonecedor e chave estrangeira de Fornecedores
	Codigo       string `gorm:"not null"` // Código utilizado pelo fornecedor na nota de compra

	Fornecedor models.Fornecedores `gorm:"foreignKey:FornecedorID"` // Relacionamento M:1
}
