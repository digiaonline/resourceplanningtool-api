package main

import (
	"resourceplanningtool-api/modules"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

func handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := graphql.Do(graphql.Params{
			Schema:		schema,
			RequestString:	string(query),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

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

	var dbinfo string
	dbinfo = fmt.Sprintf("postgres://%s:%s@%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"))
	modules.InitDB("postgres", dbinfo)

	http.Handle("/skillz", handler(schema))
	log.Fatal(http.ListenAndServe(":3002", nil))
}
