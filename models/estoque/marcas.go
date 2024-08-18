package estoque

import "api-tributacao/models"

type Marcas struct {
	models.InitEntity

	Nome string `gorm:"not null"` // Nome da Marca

}