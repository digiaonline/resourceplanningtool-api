package main

import (
	"resourceplanningtool-api/modules"
	"log"
	"fmt"
	"os"
	"net/http"

	"github.com/rs/cors"
	"github.com/joho/godotenv"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
		log.Fatal(err)
	}
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:		modules.QueryType,
		Mutation:	modules.MutationType,
	})

	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(&handler.Config{
		Schema:		&schema,
		Pretty:		true,
		GraphiQL:	true,
	})

	c := cors.New(cors.Options{
		AllowedOrigins:	[]string{"*"},
		AllowedMethods:	[]string{"GET", "POST"},
	})

	mh := c.Handler(h)

	var dbinfo string
	dbinfo = fmt.Sprintf("postgres://%s:%s@%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"))
	modules.InitDB("postgres", dbinfo)

	http.Handle("/skillz", mh)
	http.ListenAndServe(":3002", nil)
}
