package db

import (
	"api-tributacao/config"
	model "api-tributacao/src/db/models"
	"api-tributacao/src/db/seeder"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// tabelas a sincronizar com o GORM, a ser removido quando lançar para prod
var tablesToSync = []any{

	//Modelos Comuns
	&model.EnderecoPessoa{},
	&model.Cidades{},
	&model.Estados{},
	&model.Paises{},

	//Modelos Cadastros
	&model.Fornecedores{},

	//Modelos Estoque
	&model.Produtos{},
	&model.CodigosProduto{},
	&model.CustoProdutos{},
	&model.Grupos{},
	&model.LocalEstoque{},
	&model.Marcas{},
	&model.MovimentoEstoque{},
	&model.PrecoProdutos{},
	&model.SaldoEstoque{},
	&model.SubGrupos{},
	&model.TributacaoProdutos{},
	&model.Unidades{},
}

// Creates a new GORM instance and inits the connetion pool, panics if anything goes wrong
func New(cfg *config.Config) *gorm.DB {
	gormConfig := gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	// TODO: arrumar antes de prod
	if true {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(cfg.Db.Url), &gormConfig)
	if err != nil {
		log.Fatalf("failed to open db conn: %v", err)
	}

	// setup custom many to many relationships
	//err = db.SetupJoinTable(&model.LocalTrabalho{}, "Efetivos", &model.ColaboradorEfetivoLocalTrabalho{})
	//if err != nil {
	//	log.Fatalf("failed to setup many to many relationship: %v", err)
	//}

	dbSeeder := seeder.New(db)
	if cfg.Db.Clear {
		fmt.Println("[DB] clearing public schema")
		dbSeeder.Clear()
	}

	// sync with no fks to avoid errors due to circular dependencies
	if cfg.Db.Sync {
		db.DisableForeignKeyConstraintWhenMigrating = true

		fmt.Println("[DB] syncing tables with gorm model")
		if err := db.AutoMigrate(tablesToSync...); err != nil {
			log.Fatalf("failed to sync db schema db conn: %v", err)
		}
	}

	// por mais estranho que seja rodar o seeder antes das migrações, algumas
	// migrações atuão sobre os dados criados pelo seeder, então precisa ser posterior.
	if cfg.Db.Seed {
		fmt.Println("[DB] seeding database")
		dbSeeder.Seed()
	}

	if cfg.Db.Migrate {
		fmt.Println("[DB] running migrations")
		sqlDb, err := db.DB()
		if err != nil {
			log.Fatalf("failed to migrate schema: %v", err)
		}

		if err := migrateSchema(sqlDb, cfg.Db.Name); err != nil {
			log.Fatalf("failed to migrate schema: %v", err)
		}
	}

	// sync again to create the FKS, at this point the tables have been
	// created so circular dependencies are no problem
	if cfg.Db.Sync {
		db.DisableForeignKeyConstraintWhenMigrating = false

		fmt.Println("[DB] creating foreign keys")
		if err := db.AutoMigrate(tablesToSync...); err != nil {
			log.Fatalf("failed to add foreign keys: %v", err)
		}
	}

	return db
}

func CloseConnection(db *gorm.DB) {
	dbInstance, err := db.DB()
	if err == nil {
		_ = dbInstance.Close()
	}
}

// Migrates the database, closes the DB connection on failure
func migrateSchema(sqlDb *sql.DB, dbName string) error {
	driver, err := migratePostgres.WithInstance(sqlDb, &migratePostgres.Config{SchemaName: "public"})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://migrations", dbName, driver)
	if err != nil {
		return err
	}

	err = migrator.Up()

	// ErrNoChange means the migrator didnt find a new migration to execute
	// see: https://github.com/golang-migrate/migrate/issues/100
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}

/*
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
		&model.EnderecoPessoa{},
		&model.Cidades{},
		&model.Estados{},
		&model.Paises{},

		//Modelos Cadastros
		&model.Fornecedores{},

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
*/
