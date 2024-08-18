package estoque

import "api-tributacao/models"

type PrecoProdutos struct {
	models.InitEntity

	ProdutoID      uint     `gorm:"not null"`     // Chave estrangeira para Produto
	Custo          *float64 `gorm:"default:null"` // Custo Utilizado na Composição do Preço
	Indice         *float32 `gorm:"default:null"` // Índice Multiplicador do Custo para Obter o Preço
	PrecoVenda     float64  `gorm:"not null"`     // Preço de Venda
	AceitaDesconto bool     `gorm:"not null"`     // Permite Desconto?
}
