package estoque

import "api-tributacao/models"

type TributacaoProdutos struct {
	models.InitEntity

	ProdutoID   uint    `gorm:"not null"` // Chave estrangeira para Produto
	TipoTributo string  `gorm:"not null"` // Exemplo: ICMS, IPI, PIS, COFINS
	CST         string  `gorm:"not null"` // Código de Situação Tributária
	Aliquota    float64 `gorm:"not null"` // Alíquota do tributo
	ReducaoBC   float64 `gorm:"not null"` // Redução de Base de Cálculo (%)
	ModBC       int     `gorm:"not null"` // Modalidade de determinação da base de cálculo
}
