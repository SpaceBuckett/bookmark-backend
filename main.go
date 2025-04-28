package main

import (
	"database/sql"
	"github.com/SpaceBuckett/bookmark-backend/api"
	db "github.com/SpaceBuckett/bookmark-backend/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/bookmarker?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
