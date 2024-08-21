package model

type SubGrupos struct {
	InitEntity

	Nome    string `gorm:"not null"` // Nome do Grupo
	GrupoID uint   `gorm:"not null"` // Chave estrangeira para Grupo
}
