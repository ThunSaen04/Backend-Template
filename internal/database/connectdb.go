package database

import (
	"log"

	"backend-template/internal/config"

	"gorm.io/gorm"
)

// DB is the global database instance
var DB *gorm.DB

// ConnectDB establishes a connection to the database
func ConnectDB() {
	switch config.AppConfig.DBType {
	case "postgres":
		ConnectDBPostgreSQL()
	case "mysql":
		ConnectDBMySQL()
	case "sqlserver":
		ConnectDBSQLServer()
	case "sqlite":
		ConnectDBSQLite()
	default:
		log.Fatalf("Unsupported database type: %s", config.AppConfig.DBType)
	}
}
