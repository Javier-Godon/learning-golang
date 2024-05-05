package main

import (
	"database/sql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"simpleBank/api"
	db "simpleBank/db/sqlc"
)

func main() {
	const (
		dbDriver = "postgres"
		//dbSource      = "postgres://postgres:postgres@localhost:31544/postgres?sslmode=disable&search_path=public"
		dbSource      = "postgres://postgres:postgres@localhost:5444/postgres?sslmode=disable&search_path=public"
		serverAddress = "0.0.0.0:8887"
	)
	//m, err := migrate.New(
	//	"file://db/migrations",
	//	//"postgres://postgres:postgres@localhost:31544/postgres?sslmode=disable&search_path=public")
	//	"postgres://postgres:postgres@localhost:5444/postgres?sslmode=disable&search_path=public")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err := m.Up(); err != nil {
	//	log.Fatal(err)
	//}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
