package dto

// All the DTOS that are recieved from HTTP requests

type InLogin struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

type InCreateProduto struct {
	Codigo             string `json:"codigo" form:"codigo" binding:"max=20"`
	Nome               string `json:"nome" form:"nome" binding:"max=50"`
	NomeCompleto       string `json:"nomecompleto" form:"nomecompleto" binding:"max=1000"`
	NCM                string `json:"ncm" form:"ncm" binding:"max=10"`
	CEST               string `json:"cest" form:"cest" binding:"max=10"`
	UnidadeVenda       uint   `json:"unvenda" form:"unvenda"`
	UnidadeCompra      uint   `json:"uncompra" form:"uncompra"`
	QuantidadePorCaixa uint   `json:"qtdporcaixa" form:"qtdporcaixa"`
	TipoIndiceEntrada  uint   `json:"tipoentrada" form:"tipoentrada"`
	Origem             uint   `json:"origem" form:"origem"`
}
