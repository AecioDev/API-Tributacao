package estoque

import "api-tributacao/models"

type SubGrupos struct {
	models.InitEntity

	Nome string `gorm:"not null"` // Nome do Grupo
	GrupoID uint `gorm:"not null"` // Chave estrangeira para Grupo
}