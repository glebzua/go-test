package config

import "os"

type configuration struct {
	DatabaseName     string
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
}

func GetConfiguration() *configuration {
	return &configuration{
		DatabaseName:     os.Getenv("DB_NAME"),
		DatabaseHost:     os.Getenv("DB_HOST"),
		DatabaseUser:     os.Getenv("DB_USER"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
	}
}
