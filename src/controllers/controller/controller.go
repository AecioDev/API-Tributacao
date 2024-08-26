package controller

import (
	"api-tributacao/config"
	estoqueController "api-tributacao/src/controllers/estoque"
	estoqueRepository "api-tributacao/src/db/repository/estoque"
	estoqueServices "api-tributacao/src/services/estoque"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// todas as chaves de variáveis de contexto do GIN devem
// ser definidas aqui para evitar conflitos de nome

var (
	CTX_USER_KEY = "REQUEST_USER"
)

type ControllerServer struct {
	router   *gin.Engine
	database *gorm.DB
	cfgApp   *config.Config
}

func New(router *gin.Engine, database *gorm.DB, cfg *config.Config) *ControllerServer {

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsCfg.AllowAllOrigins = true

	router.Use(cors.New(corsCfg))

	return &ControllerServer{router, database, cfg} //, minIoService, zimbraService}

}

func (r *ControllerServer) StartServer(port int) {

	sys := r.router.Group("/workspace/v1")

	// ESTOQUE/PRODUTOS
	produtoRepo := estoqueRepository.NewProdutoRepository(r.database, &r.cfgApp.App)
	produtoService := estoqueServices.NewProdutoService(produtoRepo)
	produtoController := estoqueController.NewProdutoController(*produtoService)

	// Endpoints Para Produtos
	sys.GET("/estoque/produtos", produtoController.GetProdutos)
	sys.GET("/estoque/produto/:id", produtoController.GetProdutoByID)
	sys.GET("/estoque/produto-nome/:name", produtoController.GetProdutosByName)
	sys.POST("/estoque/produto", produtoController.CreateProduto)
	sys.PATCH("/estoque/produto", produtoController.UpdateProduto)
	sys.DELETE("/estoque/produto/:id", produtoController.DeleteProduto)

	// Inicia o Servidor GIN
	if err := r.router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
