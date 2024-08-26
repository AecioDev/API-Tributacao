package estoque

import (
	model "api-tributacao/src/db/models"
	"time"
)

type MovimentoEstoque struct {
	model.InitEntity

	ProdutoID      uint      `gorm:"not null"` // Chave estrangeira para Produto
	LocalEstoqueID uint      `gorm:"not null"` // Chave estrangeira para Local de Estoque
	Data           time.Time `gorm:"not null"` // Data da Movimentação
	Tipo           string    `gorm:"not null"` // Tipo de Movimentação E - Entrada / S - Saída

	SaldoEstoque float64 `gorm:"not null"` // Saldo do Estoque

	Local LocalEstoque `gorm:"foreignKey:LocalEstoqueID"` // Relacionamento M:1

}
