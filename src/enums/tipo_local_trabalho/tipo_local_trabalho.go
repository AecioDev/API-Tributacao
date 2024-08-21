package tipo_local_trabalho

type TipoLocalTrabalho string

const (
	// SEFAZ
	Sefaz TipoLocalTrabalho = "sefaz"

	// EX: Unidade de Gestão de Dados Analíticos Tributários (UGDAT)
	Unidade TipoLocalTrabalho = "unidade"

	// EX: Gabinete do Secretário de Estado (GAB/SEFAZ)
	Gabinete TipoLocalTrabalho = "gabinete"

	// EX: Coordenadoria de Tecnologia da Informação (COTIN)
	Coordenadoria TipoLocalTrabalho = "coordenadoria"

	// EX: Superintendência de Administração Tributária (SAT)
	Superintendencia TipoLocalTrabalho = "superintendencia"
)
