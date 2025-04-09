package migrations

import (
	model "go-api-catalog/models"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	// Auto migrate the Solucao model
	err := db.AutoMigrate(&model.Solucao{})
	if err != nil {
		log.Fatal("Erro ao executar migração:", err)
	}

	log.Println("Migração executada com sucesso!")
}
