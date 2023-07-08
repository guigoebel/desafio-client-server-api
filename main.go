package main

import (
	"database/sql"

	"github.com/guigoebel/desafio-client-server-api/cotation"
	dbsqlite "github.com/guigoebel/desafio-client-server-api/cotation/sqlite"
	"github.com/guigoebel/desafio-client-server-api/cotation/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := dbsqlite.NewRepository(db)
	service := cotation.NewService(repo)
	handlers := web.Handlers(service)

	err = web.Start("8080", handlers)
	if err != nil {
		panic(err)
	}

}
