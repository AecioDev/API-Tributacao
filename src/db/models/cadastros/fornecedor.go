package cadastros

import (
	model "api-tributacao/src/db/models"
)

type Fornecedores struct {
	model.InitEntity
	model.Pessoa

	RazaoSocial string `gorm:"not null"` //Nome Simples do Produto sem Detalhes

}
