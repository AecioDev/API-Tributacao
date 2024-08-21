package status_solicitacao_reajuste

type StatusReajusteSalarial string

const (
	// o pedido foi criado pelo terceirizado
	Aberto StatusReajusteSalarial = "aberto"

	// o pedido foi aprovado pelo lider da equipe do terceirizado
	// e pende avaliação do gestor
	AprovadoPorLider StatusReajusteSalarial = "aprovado_por_lider"

	// o pedido foi rejeitado pelo lider da equipe do terceirizado
	// e pende avaliação do gestor
	//
	// sim, um gestor pode aprovar um pedido recusado pelo lider
	RejeitadoPorLider StatusReajusteSalarial = "rejeitado_por_lider"

	// o pedido foi aprovado pelo gestor da unidade da equipe
	// do terceirizado e pende aprovação final do coordenador
	AprovadoPorGestor StatusReajusteSalarial = "aprovado_por_gestor"

	// o pedido foi rejeitado pelo gestor da unidade da equipe
	// do terceirizado e foi encerrado
	RejeitadoPorGestor StatusReajusteSalarial = "rejeitado_por_gestor"

	// o pedido foi aprovado pelo coordenador e pende atuação da empresa
	// responsável pelo colaborador terceirizado
	AprovadoPorCoordenador StatusReajusteSalarial = "aprovado_por_coordenador"

	// o pedido foi rejeitado pelo coordenador e encerrado
	RejeitadoPorCoodenador StatusReajusteSalarial = "rejeitado_por_coordenador"

	// o pedido foi aprovado por um usuario externo da empresa responsável
	// pelo terceirizado e foi finalmente atuado
	AprovadoPorEmpresa StatusReajusteSalarial = "aprovado_por_empresa"

	// o pedido foi rejeitado por um usuario externo da empresa responsável
	// pelo terceirizado
	RejeitadoPorEmpresa StatusReajusteSalarial = "rejeitado_por_empresa"
)
