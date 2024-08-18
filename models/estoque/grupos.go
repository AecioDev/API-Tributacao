package estoque

import "api-tributacao/models"

type Grupos struct { 
	models.InitEntity

	Nome string `gorm:"not null"` // Nome do Grupo

	SubGrupos []SubGrupos `gorm:"foreignKey:GrupoID"` // Relacionamento 1:N
}