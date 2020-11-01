package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cynthiawilliamsa/meetmeup/config"
	"github.com/cynthiawilliamsa/meetmeup/database"
	"github.com/cynthiawilliamsa/meetmeup/graph"
	"github.com/cynthiawilliamsa/meetmeup/graph/generated"
	"github.com/go-pg/pg/v10"
)

const defaultPort = "8080"

func main() {
	db := database.New(&pg.Options{
		User:     config.PG_User,
		Password: config.PG_Pass,
		Database: config.PG_Database,
	})

	defer db.Close()

	db.AddQueryHook(database.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DBRepo: database.MeetMeUpRepo{DB: db}}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
