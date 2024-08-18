package estoque

import "api-tributacao/models"

type LocalEstoque struct {
	models.InitEntity

	Nome string `gorm:"not null"` // Nome da Marca
}