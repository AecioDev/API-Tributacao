package seeder

import (
	"log"

	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Seeder {
	return &Seeder{db}
}

func (s *Seeder) Clear() {
	if err := s.db.Exec("DROP SCHEMA public CASCADE").Error; err != nil {
		log.Fatalf("failed to drop public schema: %v", err)
	}

	if err := s.db.Exec("CREATE SCHEMA public").Error; err != nil {
		log.Fatalf("failed to create public schema: %v", err)
	}
}

func (s *Seeder) Seed() {

	/*
		// isso faz com que readExcel nunca execute e a IDE não entupa de warning com código morto
		//
		// ULTIMA CARGA COM PLANILHA ATUALIDA EM 17/07/24
		if (func() bool { return false })() {
			s.readExcelColaboradores()
		}

		// ULTIMA CARGA COM PLANILHA ATUALIDA EM 17/07/24
		if (func() bool { return false })() {
			s.readExcelRamais()
		}

		if (func() bool { return false })() {
			s.readExcelLocais()
		}
	*/

	s.seedAll()
}

func (s *Seeder) seedAll() {

}
