package status_solicitacao_colaborador

type StatusSolicitacao string

const (
	// a solicitação foi criada pelo gestor e esta "em aberto"
	Aberta StatusSolicitacao = "aberta"

	// a solicitação cancelada pelo gestor que a criou
	Cancelada StatusSolicitacao = "cancelada"

	// um efetivo ou terceirizado com permissão de atender a solicitação a recusou
	RecusadaPorAtendente StatusSolicitacao = "recusada_por_atendente"

	// a solicitação foi aprovada pelo atendente e encaminhada para a empresa terceirizada
	//
	// nota-se que solicitações criadas diretamente pelo atendente já vem com esse status
	// como o status inicial
	EncaminhadaParaEmpresa StatusSolicitacao = "encaminhada_para_empresa"

	// a solicitação foi encaminhada para a empresa e esta a recusou
	RecusadaPorEmpresa StatusSolicitacao = "recusada_por_empresa"

	// a solicitação foi criada, aprovada pelo atendente, encaminhada para a empresa
	// a qual cadastrou um novo colaborador para atender a solicitação.
	Atendida StatusSolicitacao = "atendida"

	// a vaga que a solicitação visava atender foi preenchida por outra solicitação
	AtendidaPorOutra StatusSolicitacao = "atendida_por_outra"

	// a vaga que a solicitação visava atender foi encerrada/removida
	FechadaPorEncerramentoDaVaga StatusSolicitacao = "fechada_por_encerramento_da_vaga"
)
