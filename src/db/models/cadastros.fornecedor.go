package model

type Fornecedores struct {
	InitEntity
	Pessoa

	RazaoSocial string `gorm:"not null"` //Nome Simples do Produto sem Detalhes

}
