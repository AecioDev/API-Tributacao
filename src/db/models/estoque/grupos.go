package estoque

import (
	model "api-tributacao/src/db/models"
)

type Grupos struct {
	model.InitEntity

	Nome string `gorm:"not null"` // Nome do Grupo

	SubGrupos []SubGrupos `gorm:"foreignKey:GrupoID"` // Relacionamento 1:N
}
