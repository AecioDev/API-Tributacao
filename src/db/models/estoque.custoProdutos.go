package model

type CustoProdutos struct {
	InitEntity

	ProdutoID      uint     `gorm:"not null"`     // Chave estrangeira para Produto
	FatoGerador    string   `gorm:"not null"`     // O Fato Gerado pode ser Manual o Custo foi Digitado Manualmente, ou pode ter sido proveniente da Entrada de uma Nota de Compras.
	FatoGeradoId   *uint    `gorm:"default:null"` // Código da Nota de Entrada ou Outro Fato Gerador q altere o Custo.
	CustoBruto     *float64 `gorm:"default:null"` // Custo Utilizado na Composição do Preço
	CustoReposicao *float64 `gorm:"default:null"` // Custo Utilizado na Composição do Preço
	CustoMedio     *float64 `gorm:"default:null"` // Custo Utilizado na Composição do Preço

	//Impostos Vinculados ao Custo
	IcmsCreditado      *float64 `gorm:"default:null"`
	IcmsST             *float64 `gorm:"default:null"`
	IPIGerado          *float64 `gorm:"default:null"`
	IPICreditado       *float64 `gorm:"default:null"`
	PisCofinsCreditado *float64 `gorm:"default:null"`
	DifAliquota        *float64 `gorm:"default:null"`
	CustoImportacao    *float64 `gorm:"default:null"`
	Frete              *float64 `gorm:"default:null"`
	OutrosCustos       *float64 `gorm:"default:null"`

	//Dados Adicionais
	QuantidadeAnterior      *float64 `gorm:"default:null"` //Quantidade no Momento da Alteração do Custo
	QuantidadeEntrada       *float64 `gorm:"default:null"` //Quantidade de Entrada do Produto
	QuantidadeFinal         *float64 `gorm:"default:null"` //Quantidade Final após a Entrada do Produto
	ValorUnitarioEntrada    *float64 `gorm:"default:null"` //Valor unitário de Entrada do Produto
	HistoricoAlteracaoCusto *string  `gorm:"default:null"` //Observações da Alteração do Custo
}
