package model

type Marcas struct {
	InitEntity

	Nome string `gorm:"not null"` // Nome da Marca

}
