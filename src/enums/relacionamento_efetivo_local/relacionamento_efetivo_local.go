package relacionamento_efetivo_local

type RelEfetivoLocal string

const (
	Gestor RelEfetivoLocal = "gestor"

	Secretario RelEfetivoLocal = "secretario"

	SecretarioSubstituto RelEfetivoLocal = "secretario_substituto"

	Assesor RelEfetivoLocal = "assesor"

	Superintendente RelEfetivoLocal = "superintendente"

	SuperintendenteSubstituto RelEfetivoLocal = "superintendente_substituto"

	Coordenador RelEfetivoLocal = "coordenador"

	CoordenadorSubstituto RelEfetivoLocal = "coordenador_substituto"
)
