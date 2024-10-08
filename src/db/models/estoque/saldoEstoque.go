package estoque

import (
	model "api-tributacao/src/db/models"
)

type SaldoEstoqueProduto struct {
	model.InitEntity

	ProdutoID      uint    `gorm:"not null"` // Chave estrangeira para Produto
	LocalEstoqueID uint    `gorm:"not null"` // Chave estrangeira para Local de Estoque
	SaldoEstoque   float64 `gorm:"not null"` // Saldo do Estoque

	Local LocalEstoque `gorm:"foreignKey:LocalEstoqueID"` // Relacionamento M:1
}
