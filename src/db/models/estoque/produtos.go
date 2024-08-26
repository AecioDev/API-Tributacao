package estoque

import (
	model "api-tributacao/src/db/models"
)

type Produtos struct {
	model.InitEntity

	Codigo             string `json:"codigo" gorm:"unique;not null"` // Código do produto - Código de Barras ou para controle Interno
	Nome               string `json:"nome" gorm:"not null"`          // Nome Simples do Produto sem Detalhes
	NomeCompleto       string `json:"nomecompleto"`                  // Nome Completo do Produto Detalhado
	NCM                string `json:"ncm"`                           // NCM (Código da Nomenclatura Comum do Mercosul)
	CEST               string `json:"cest"`                          // Código Especificador da Substituição Tributária
	UnidadeVenda       uint   `json:"unvenda"`                       // Unidade de medida (ex: KG, UN, etc.)
	UnidadeCompra      uint   `json:"uncompra"`                      // Unidade de medida (ex: KG, UN, etc.)
	QuantidadePorCaixa uint   `json:"qtdporcaixa" gorm:"default:1"`  // Default 1 e caso maior que 1 a Quantidade de Estoque Será Multiplicada ou dividida baseada no Tipo de Indice
	TipoIndiceEntrada  uint   `json:"tipoentrada" gorm:"default:0"`  // 0 - Mantém a Quantidade De Entrada da Nota | 1 - Multiplica a Quantidade de Entrada Pela Quantidade por Caixas | 2 - Divide a Quantidade de Entrada pela Quantidade por Caixas
	Origem             uint   `json:"origem" gorm:"not null"`        // Código de origem do produto (Ex: Nacional, Importado)

	//Relacionamentos 1 pra 1
	UnidadeVnd *Unidades `gorm:"foreignKey:UnidadeVenda"`  // Relacionamento M:1
	UnidadeCmp *Unidades `gorm:"foreignKey:UnidadeCompra"` // Relacionamento M:1

	//Relacionamentos 1 pra Muitos
	Custos           []CustoProdutos      `gorm:"foreignKey:ProdutoID"` // Históri de Custos do produto
	Precos           []PrecoProdutos      `gorm:"foreignKey:ProdutoID"` // Preço do produto
	Codigos          []CodigosProduto     `gorm:"foreignKey:ProdutoID"` // Cadastro dos Códigos de Barras por Fornecedor
	Tributacoes      []TributacaoProdutos `gorm:"foreignKey:ProdutoID"` // Cadastro das Tributações do Produto
	SaldoEstoque     []SaldoEstoque       `gorm:"foreignKey:ProdutoID"` // O Produto pode ter saldo em mais de um Local de Estoque
	MovimentoEstoque []MovimentoEstoque   `gorm:"foreignKey:ProdutoID"` // Movimentações de Entrada e Saída do produto
}

func (Produtos) TableName() string {
	return "produtos"
}
