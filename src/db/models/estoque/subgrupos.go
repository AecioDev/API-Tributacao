package estoque

import (
	model "api-tributacao/src/db/models"
)

type SubGrupos struct {
	model.InitEntity

	Nome    string `gorm:"not null"` // Nome do Grupo
	GrupoID uint   `gorm:"not null"` // Chave estrangeira para Grupo
}
