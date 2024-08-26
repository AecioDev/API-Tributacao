package estoque

import (
	model "api-tributacao/src/db/models"
)

type TributacaoProdutos struct {
	model.InitEntity

	ProdutoID   uint    `gorm:"not null"` // Chave estrangeira para Produto
	TipoTributo string  `gorm:"not null"` // Exemplo: ICMS, IPI, PIS, COFINS
	CST         string  `gorm:"not null"` // Código de Situação Tributária
	Aliquota    float64 `gorm:"not null"` // Alíquota do tributo
	ReducaoBC   float64 `gorm:"not null"` // Redução de Base de Cálculo (%)
	ModBC       int     `gorm:"not null"` // Modalidade de determinação da base de cálculo
}
