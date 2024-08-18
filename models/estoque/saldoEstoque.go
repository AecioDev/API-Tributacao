package estoque

import "api-tributacao/models"

type SaldoEstoque struct {
	models.InitEntity

	ProdutoID uint `gorm:"not null"` // Chave estrangeira para Produto
	LocalEstoqueID uint `gorm:"not null"` // Chave estrangeira para Local de Estoque
	SaldoEstoque float64 `gorm:"not null"` // Saldo do Estoque	

	Local LocalEstoque `gorm:"foreignKey:LocalEstoqueID"` // Relacionamento M:1
}