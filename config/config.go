package config

import (
	"os"
)

type Configuration struct {
	DatabaseName        string
	DatabaseHost        string
	DatabaseUser        string
	DatabasePassword    string
	AuthAccessKeySecret string
	MigrateToVersion    string
	MigrationLocation   string
}

func GetConfiguration() *Configuration {
	migrationLocation, set := os.LookupEnv("MIGRATION_LOCATION")
	if !set {
		migrationLocation = "internal/infra/database/migrations"
	}
	migrateToVersion, set := os.LookupEnv("MIGRATE")
	if !set {
		migrateToVersion = "latest"
	}
	return &Configuration{
		DatabaseName:        os.Getenv("DB_EVENTS"),
		DatabaseHost:        os.Getenv("DB_HOST"),
		DatabaseUser:        os.Getenv("DB_USER"),
		DatabasePassword:    os.Getenv("DB_PASSWORD"),
		MigrateToVersion:    migrateToVersion,
		MigrationLocation:   migrationLocation,
		AuthAccessKeySecret: os.Getenv("EVENTS_AUTH_ACCESS_SECRET"),
	}
}
