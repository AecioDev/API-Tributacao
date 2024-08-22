package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Modelos Genéricos para serem utilizados em conjunto com os Modelos Específicos.

type Response struct {
	Message string
}

type Pessoa struct {
	Nome string `gorm:"not null"` // Nome da Pessoa
	CPF  uint   `gorm:"not null"` // CPF da Pessoa
	CNPJ uint   `gorm:"not null"` // CNPJ da Pessoa

	Enderecos []EnderecoPessoa `gorm:"foreignKey:PessoaID"` // Relacionamento 1:N
}

type TipoEndereco string

const (
	EnderecoEntrega     TipoEndereco = "Entrega"
	EnderecoCobranca    TipoEndereco = "Cobrança"
	EnderecoResidencial TipoEndereco = "Residencial"
	EnderecoComercial   TipoEndereco = "Comercial"
)

type EnderecoPessoa struct {
	InitEntity

	PessoaID          uint         `gorm:"not null"` // Chave estrangeira para Pessoa
	CEP               uint         `gorm:"not null"` // CEP do Endereço
	TipoLogradouro    string       `gorm:"not null"` // Tipo do Endereço Rua, Avenida, Travessa etc.
	NomeLogradouro    string       `gorm:"not null"` // Rua do Endereço
	Numero            string       `gorm:"not null"` // Numero do Endereço
	TipoBairro        string       `gorm:"not null"` // Tipo de Bairro do Endereço Vila, Bairro, Conjunto, Chácara, Fazenda etc.
	NomeBairro        string       `gorm:"not null"` // Bairro do Endereço
	Complemento       string       `gorm:"not null"` // Complemento do Endereço
	PontoDeReferencia string       `gorm:"not null"` // Ponto de Referência do Endereço
	TipoEndereco      TipoEndereco `gorm:"not null"` // Usando o tipo customizado para o tipo de endereço

	CidadeID uint    // Cidade do Endereço e Chave estrangeira para Cidades
	Cidade   Cidades `gorm:"foreignKey:CidadeID"` // Relacionamento M:1
}

type Cidades struct {
	ID         uint   `gorm:"primaryKey"`
	NomeCidade string `gorm:"not null"` // Nome da Cidade
	CodigoIBGE uint   `gorm:"not null"` // Código IBGE da Cidade

	EstadoID uint    // Chave estrangeira para o Estado
	Estado   Estados `gorm:"foreignKey:EstadoID"` // Relacionamento M:1
}

type Estados struct {
	ID                             uint   `gorm:"primaryKey"`
	EstadoNome                     string `gorm:"not null"` // Nome do Estado
	EstadoUF                       string `gorm:"not null"` // Sigla da UF do Estado
	EstadoCodigo                   uint   `gorm:"not null"` // Código IBGE do Estado
	AliquotaInternaICMS            decimal.Decimal
	AliquotaICMSVendaInterestadual decimal.Decimal
	AliquotaFCPNormal              decimal.Decimal
	AliquotaICMSST                 decimal.Decimal
	AliquotaFCPST                  decimal.Decimal

	PaisID uint   // Chave estrangeira para Paises
	Pais   Paises `gorm:"foreignKey:PaisID"` // Relacionamento M:1
}

type Paises struct {
	ID               uint   `gorm:"primaryKey"`
	CodigoBacen      uint   `gorm:"not null"` // Código Bacen do País
	NomePais         string `gorm:"not null"` // Nome do País
	CodigoTelefonico string `gorm:"not null"` // Código Telefônico do País ex. Brasil +55
}

type InitEntity struct {
	gorm.Model

	CreatedFor string // Nome do Usuário que Criou
	UpdatedFor string // Nome do Usuário que Alterou
	DeletedFor string // Nome do Usuário que Deletou
}

func (e *InitEntity) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreatedFor = "Sistema/Usuário Atual" // Insira a lógica para pegar o usuário atual
	return
}

func (e *InitEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdatedFor = "Sistema/Usuário Atual" // Insira a lógica para pegar o usuário atual
	return
}

func (e *InitEntity) BeforeDelete(tx *gorm.DB) (err error) {
	e.DeletedFor = "Sistema/Usuário Atual" // Insira a lógica para pegar o usuário atual
	return
}
