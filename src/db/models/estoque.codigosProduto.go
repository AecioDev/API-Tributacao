package model

type CodigosProduto struct {
	InitEntity

	ProdutoID    uint   `gorm:"not null"` // Chave estrangeira para Produto
	FornecedorID uint   `gorm:"not null"` // Código do Fonecedor e chave estrangeira de Fornecedores
	Codigo       string `gorm:"not null"` // Código utilizado pelo fornecedor na nota de compra

	Fornecedor Fornecedores `gorm:"foreignKey:FornecedorID"` // Relacionamento M:1
}
