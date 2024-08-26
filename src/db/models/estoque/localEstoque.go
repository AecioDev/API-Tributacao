package estoque

import (
	model "api-tributacao/src/db/models"
)

type LocalEstoque struct {
	model.InitEntity

	Nome string `gorm:"not null"` // Nome da Marca
}
