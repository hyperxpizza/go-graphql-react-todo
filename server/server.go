package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hyperxpizza/go-react-gql-todo/database"
	"github.com/hyperxpizza/go-react-gql-todo/graph"
	"github.com/hyperxpizza/go-react-gql-todo/graph/generated"
	"github.com/joho/godotenv"
)

func init() {
	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("[-] Can not load .env file")
	}
}

func main() {

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	port := os.Getenv("PORT")

	database := database.InitDB(dbUser, dbPassword, dbName)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Database: database}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
