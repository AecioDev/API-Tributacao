package enums

// uma permissão que permite realizar uma ação ou acessar um recurso dessa API
//
// As permissões seguem um modelo simples, se o usuário tem a permissão "visualizar_noticias"
// então ele pode visualizar todas as notícias, caso haja alguma regra de negócio como por
// exemplo limitar a visualização de notícias para apenas as notícias de sua unidade, essa regra
// de negócio é expressa no código dos endpoints, não como uma permissão.
//
// Um nivel de acesso tem N permissões, um detalhe importante é não confundir relacionamentos
// de entidades expressos no banco e os níveis de acesso.
//
// Por exemplo um gestor é um colaborador_efetivo que tem um relacionamento de tipo "gestor" com a tabela local_trabalho
// é um gestor daquele local de trabalho.
type Permissao string

const (
	// permite criar, editar e deletar equipes de QUALQUER unidade (local_trabalho)
	//
	// importante: gestores de uma unidade podem cadastrar/editar equipes p/ sua unidade
	// mesmo sem ter essa permissão
	GerenciarEquipes Permissao = "gerenciar_equipes"

	// da ao usuário a possibilidade de gerenciar os níveis de acesso dos usuários
	// sendo possível criar, atualizar, deletar níveis de acesso e definir quais
	// usuários tem quais níveis de acesso.
	GerenciarNiveisAcesso Permissao = "gerenciar_niveis_acesso"

	// confirma um pedido de desligamento aprovado pelo gestor da unidade
	// dona da equipe a qual o terceirizado a ser desligado pertence.
	//
	// essa ação é a confirmação de coordenador da cotin necessária para executar o pedido
	// e de fato desligar o colaborador.
	ConfirmarPedidoDesligamento Permissao = "confirmar_pedido_desligamento"

	// atender (recusar ou aprovar e encaminhar p/ empresa) um pedido de solicitação
	// de colaborador.
	//
	// também permite que um pedido seja criado diretamente pelo atendente
	AtenderSolicitacaoColaborador Permissao = "atender_solicitacao_colaborador"

	// permite alterar a equipe de um colaborador,
	// registrando a alteração no histórico
	//
	// importante: gestores podem alterar equipe de seus colaboradores,
	// independente de terem essa permissão ou não
	AlterarEquipeDeColaborador Permissao = "alterar_equipe_de_colaborador"

	// permite alterar dados cadastrais de um colaborador terceirizado
	//
	// importante: usuarios da empresa terceirizada podem atualizar um terceirizado
	AtualizarTerceirizado Permissao = "atualizar_terceirizado"

	// permite alterar dados cadastrais de um colaborador efetivo
	AtualizarEfetivo Permissao = "atualizar_efetivo"

	// negar um pedido de desligamento já aprovado por um gestor.
	NegarPedidoDesligamento Permissao = "negar_pedido_desligamento"

	// avaliar (confirmar ou recusar) um pedido de reajuste salárial
	// que foi aprovado tanto pelo lider da equipe do terceirizado a
	// ter seu salário reajustado como quanto por seu gestor
	AvaliarPedidoReajusteSalarial Permissao = "avaliar_pedido_reajuste_salarial"

	// permite visualizar todos os pedidos de desligamento independentemente da equipe/unidade
	// ao qual o terceirizado do pedido pertence
	VisualizarPedidosDesligamento Permissao = "visualizar_pedidos_desligamento"

	// permite visualizar todos os pedidos de reajuste independentemente da equipe/unidade
	// ao qual o terceirizado do pedido pertence
	VisualizarPedidosReajusteSalarial Permissao = "visualizar_pedidos_reajuste_salarial"

	// permite visualizar todos os pedidos de ferias, independentemente
	// de se o terceirizado que pediu as férias pertence a mesma
	// equipe e/ou unidade do usuario requisitante
	VisualizarPedidosFerias Permissao = "visualizar_pedidos_ferias"

	// permite visualizar o historico de movimentação de equipes de um colaborador
	VisualizarHistoricoEquipeDeColaborador Permissao = "visualizar_historico_equipe_de_colaborador"
)

// Lista de todas as permissões da API
var TodasPermissoes []Permissao = []Permissao{
	ConfirmarPedidoDesligamento,
	NegarPedidoDesligamento,
	GerenciarNiveisAcesso,
	VisualizarPedidosDesligamento,
	VisualizarPedidosReajusteSalarial,
	AvaliarPedidoReajusteSalarial,
	VisualizarPedidosFerias,
	AlterarEquipeDeColaborador,
	AtualizarTerceirizado,
	AtualizarEfetivo,
	GerenciarEquipes,
}
