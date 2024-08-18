package estoque

import "api-tributacao/models"

type Unidades struct {
	models.InitEntity

	Sigla     string `gorm:"not null"` // Sigla da Unidade KG, MT, LT, CXA
	Descricao string `gorm:"not null"` // Descrição da Unidade Quilo, Metro, Litro, Caixa
}
