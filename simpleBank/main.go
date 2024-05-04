package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	fmt.Print("Intentando conectar a la base de datos")

	m, err := migrate.New(
		"file://db/migrations",
		"postgres://postgres:postgres@localhost:31544/postgres?sslmode=disable&search_path=public")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

}
