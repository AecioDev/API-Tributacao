package model

type LocalEstoque struct {
	InitEntity

	Nome string `gorm:"not null"` // Nome da Marca
}
