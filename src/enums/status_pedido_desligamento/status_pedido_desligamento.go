package status_pedido_desligamento

type StatusPedidoDesligamento string

const (
	// o pedido de desligamento foi aberto pela empresa
	// terceirizada e pende avaliaçâo do gestor
	Aberto StatusPedidoDesligamento = "aberto"

	// o pedido de desligamento foi aceito pelo gestor da unidade dona da equipe
	// a qual o terceirizado pertence, o pedido pende confirmação do coordenador
	AceitoPorGestor StatusPedidoDesligamento = "aceito_por_gestor"

	// o pedido de desligamento foi confirmado pelo coordenador e efetuado
	Confirmado StatusPedidoDesligamento = "confirmado_por_coordenador"

	// o pedido de desligamento foi recusado pelo gestor da unidade dona da equipe
	// a qual o terceirizado pertence
	RecusadoPorGestor StatusPedidoDesligamento = "recusado_por_gestor"

	// o pedido de desligamento foi aprovado pelo gestor, mas negado pelo coordenador
	RecusadoPorCoordenador StatusPedidoDesligamento = "recusado_por_coordenador"

	// o pedido de desligamento foi fechado por seu criador
	Fechado StatusPedidoDesligamento = "fechado"

	// o pedido de desligamento foi fechado porque outro pedido apontando para o mesmo colaborador
	// foi aceito ou o colaborador foi desligado em outro evento com o encerramento da empresa
	FechadoPorOutroPedido StatusPedidoDesligamento = "fechado_por_outro_pedido"
)
