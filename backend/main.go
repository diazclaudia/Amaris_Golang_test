package main

import (
	"backend/api"
	"backend/config"
	"log"
	"net/http"

	read "backend/domain/read"
	update "backend/domain/write"
	read_repo "backend/repository/read"
	write_repo "backend/repository/write"
	"github.com/go-chi/chi"
)

const tableName = "points"

func main() {
	conf, _ := config.NewConfig("./config/config.yaml")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	log.Println("starting")

	log.Printf("creating DynaomDB client for DynamoDB Local running on port %v\n", conf.Database.Port)
	repo, dynamodbClient, err := write_repo.NewRepository(conf.Database.URL, conf.Database.DB, conf.Database.Port, conf.Database.Timeout)
	if err != nil {
		panic("error creating dynamodb client")
	}
	log.Printf("creating DynaomDB in aws \n")
	repoRead, dynamodbClientRead, err := read_repo.NewRepositoryRead(conf.Aws.Key, conf.Aws.Secret)
	if err != nil {
		panic("error creating dynamodb client")
	}

	log.Printf("creating table '%s' if it does not already exist\n", tableName)
	didCreateTableRead := read_repo.CreateTableIfNotExists(dynamodbClientRead)
	log.Printf("did create table '%s'? %v\n", tableName, didCreateTableRead)
	log.Printf("creating table '%s' if it does not already exist\n", tableName)
	didCreateTable := write_repo.CreateTableIfNotExists(dynamodbClient)
	log.Printf("did create table '%s'? %v\n", tableName, didCreateTable)
	log.Println("running seed items example")
	write_repo.SeedItems(dynamodbClient)
	log.Println("running seed read items example")
	read_repo.SeedItems(dynamodbClientRead)
	log.Println("completed")

	r := chi.NewRouter()
	serviceRead := read.NewPointServiceRead(repoRead)
	serviceUpdate := update.NewPointService(repo, conf)

	handler := api.NewHandler(serviceRead, serviceUpdate)

	r.Get("/id/{id}", handler.Get)
	r.Post("/points/{points}/id/{id}", handler.Post)

	log.Fatal(http.ListenAndServe(conf.Server.Port, r))
}
