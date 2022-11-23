package database

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPGPostgre() *gorm.DB {
	var (
		dbUser  = "postgres"
		dbPass  = "password"
		dbHost  = "postgres"
		dbName  = "postgres"
		dbPort  = "5432"
		sslMode = "disable"
	)

	// dsn
	dsn := fmt.Sprintf(`
		host=%s user=%s password=%s dbname=%s port=%s sslmode=%s`,
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
		sslMode,
	)

	log.Print("dsn:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		// set without default transaction
		// will call manually per-case query
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Warn("Connected to database local Failed:", err)
	}
	log.Warn("Connected to database local")

	return db
}
