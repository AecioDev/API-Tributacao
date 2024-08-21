package status_solicitacao_ferias

type StatusSolicitacaoFerias string

const (
	// o pedido foi criado pelo terceirizado
	Aberto StatusSolicitacaoFerias = "aberto"

	// o pedido foi aprovado pelo lider da equipe do terceirizado
	// e pende avaliação do gestor
	AprovadoPorLider StatusSolicitacaoFerias = "aprovado_por_lider"

	// o pedido foi rejeitado pelo lider da equipe do terceirizado
	// e pende avaliação do gestor
	//
	// sim, um gestor pode aprovar um pedido recusado pelo lider
	RejeitadoPorLider StatusSolicitacaoFerias = "rejeitado_por_lider"

	// o pedido foi aprovado pelo gestor da unidade da equipe
	// do terceirizado e pende aprovação final do coordenador
	AprovadoPorGestor StatusSolicitacaoFerias = "aprovado_por_gestor"

	// o pedido foi rejeitado pelo gestor da unidade da equipe
	// do terceirizado e foi encerrado
	RejeitadoPorGestor StatusSolicitacaoFerias = "rejeitado_por_gestor"

	// o pedido foi aprovado pela empresa terceirizada responsável pelo
	// colaborador terceirizado
	AprovadoPorEmpresa StatusSolicitacaoFerias = "aprovado_por_empresa"

	// o pedido foi rejeitado pela empresa terceirizada responsável pelo
	// colaborador terceirizado
	RejeitadoPorEmpresa StatusSolicitacaoFerias = "rejeitado_por_empresa"
)
