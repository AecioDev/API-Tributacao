package db

import (
	"database/sql"
	"fmt"

	"api-tributacao/models"
	"api-tributacao/models/estoque"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "bdapitrib"
)

func ConnectDB() (*gorm.DB, error) {

	fmt.Println("Verifica se o Banco de Dados Existe.")

	err := CreateBD(dbname)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println("String Conexão: ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Desativa a pluralização
		},
	})

	if err != nil {
		return nil, err
	}

	// Migrar automaticamente a estrutura das tabelas
	err = db.AutoMigrate(
		//Modelos Comuns
		&models.EnderecoPessoa{},
		&models.Cidades{},
		&models.Estados{},
		&models.Paises{},

		//Modelos Cadastros
		&models.Fornecedores{},

		//Modelos Estoque
		&estoque.Produtos{},
		&estoque.CodigosProduto{},
		&estoque.CustoProdutos{},
		&estoque.Grupos{},
		&estoque.LocalEstoque{},
		&estoque.Marcas{},
		&estoque.MovimentoEstoque{},
		&estoque.PrecoProdutos{},
		&estoque.SaldoEstoque{},
		&estoque.SubGrupos{},
		&estoque.TributacaoProdutos{},
		&estoque.Unidades{},
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao Banco de Dados: ", dbname)
	return db, nil
}

func CreateBD(targetDBName string) error {

	exists, err := checkDatabaseExists(targetDBName)
	if err != nil {
		panic(err)
	}

	if !exists {
		// Criar o banco de dados se ele não existir
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable", host, port, user, password)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return err
		}
		defer db.Close()

		// Query para criar o banco de dados
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", targetDBName)

		_, err = db.Exec(createDBQuery)
		if err != nil {
			return err
		}

		db.Close()

		fmt.Println("Banco de Dados Criado com Sucesso!")
	}

	return nil

}

func checkDatabaseExists(dbName string) (bool, error) {
	var exists bool

	// Criar o banco de dados se ele não existir
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable", host, port, user, password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return false, err
	}
	defer db.Close()

	query := `SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)`
	err = db.QueryRow(query, dbName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
