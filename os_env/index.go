package main

import (
	"fmt"
	"os"
)

/*
	DB_HOST=10.10.100.1
	DB_USER=postgres
	DB_PORT=5432
	DB_NAME=db_enigma
	DB_PASSWORD=12345

	koneksi = postgress://postgres:12345@localhost:5432/db_enigma
*/

type Config struct {
	DataSourceName string
}

func NewConfig() *Config {
	config := new(Config)

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	config.DataSourceName = dsn
	return config
}

func main() {
	// config := NewConfig()
	// fmt.Println(config.DataSourceName)

	argsTest := os.Args
	fmt.Println(argsTest)
}
