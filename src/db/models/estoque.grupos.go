package model

type Grupos struct {
	InitEntity

	Nome string `gorm:"not null"` // Nome do Grupo

	SubGrupos []SubGrupos `gorm:"foreignKey:GrupoID"` // Relacionamento 1:N
}
