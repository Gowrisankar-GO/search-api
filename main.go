package main

import (
	"log"
	"net/http"
	"search_api/controller"
	"search_api/dbconfig"
)

var (
	defaultPort = "8080"
)

func main() {

	srv := http.NewServeMux()

	db := dbconfig.DbConn()

	err := dbconfig.MigrateTable(db)

	if err != nil {

		log.Fatal("failed to migrate table")
	}

	err = dbconfig.CreateIndexAndExtensions(db)

	if err != nil {

		log.Fatal("failed to create index")
	}

	err = dbconfig.PopulateUser(db)

	if err != nil {

		log.Fatal("failed to populate user data")
	}

	dependency := controller.Dependency{DB: db}

	srv.HandleFunc("/search", dependency.SearchUser)

	log.Fatal(http.ListenAndServe(":"+defaultPort, srv))

}
