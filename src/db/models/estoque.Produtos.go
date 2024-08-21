package model

type Produtos struct {
	InitEntity

	Codigo             string `gorm:"unique;not null"` // Código do produto - Código de Barras ou para controle Interno
	NomeSimples        string `gorm:"not null"`        // Nome Simples do Produto sem Detalhes
	NomeCompleto       string `gorm:"not null"`        // Nome Completo do Produto Detalhado
	NCM                string `gorm:"not null"`        // NCM (Código da Nomenclatura Comum do Mercosul)
	CEST               string `gorm:"not null"`        // Código Especificador da Substituição Tributária
	UnidadeVenda       uint   `gorm:"not null"`        // Unidade de medida (ex: KG, UN, etc.)
	UnidadeCompra      uint   `gorm:"not null"`        // Unidade de medida (ex: KG, UN, etc.)
	QuantidadePorCaixa uint   `gorm:"default:1"`       // Default 1 e caso maior que 1 a Quantidade de Estoque Será Multiplicada ou dividida baseada no Tipo de Indice
	TipoIndiceEntrada  uint   `gorm:"default:0"`       // 0 - Mantém a Quantidade De Entrada da Nota | 1 - Multiplica a Quantidade de Entrada Pela Quantidade por Caixas | 2 - Divide a Quantidade de Entrada pela Quantidade por Caixas
	Origem             int    `gorm:"not null"`        // Código de origem do produto (Ex: Nacional, Importado)

	//Relacionamentos 1 pra 1
	UnidadeVnd Unidades `gorm:"foreignKey:UnidadeVenda"`  // Relacionamento M:1
	UnidadeCmp Unidades `gorm:"foreignKey:UnidadeCompra"` // Relacionamento M:1

	//Relacionamentos 1 pra Muitos
	Custos           []CustoProdutos      `gorm:"foreignKey:ProdutoID"` // Históri de Custos do produto
	Precos           []PrecoProdutos      `gorm:"foreignKey:ProdutoID"` // Preço do produto
	Codigos          []CodigosProduto     `gorm:"foreignKey:ProdutoID"` // Cadastro dos Códigos de Barras por Fornecedor
	Tributacoes      []TributacaoProdutos `gorm:"foreignKey:ProdutoID"` // Cadastro das Tributações do Produto
	SaldoEstoque     []SaldoEstoque       `gorm:"foreignKey:ProdutoID"` // O Produto pode ter saldo em mais de um Local de Estoque
	MovimentoEstoque []MovimentoEstoque   `gorm:"foreignKey:ProdutoID"` // Movimentações de Entrada e Saída do produto
}
