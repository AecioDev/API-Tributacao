package estoque

import (
	model "api-tributacao/src/db/models"
)

type Unidades struct {
	model.InitEntity

	Sigla string `gorm:"not null"` // Sigla da Unidade KG, MT, LT, CXA
	Nome  string `gorm:"not null"` // Descrição da Unidade Quilo, Metro, Litro, Caixa
}
