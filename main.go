package main

import (
	"skillz/modules"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:		modules.QueryType,
		Mutation:	modules.MutationType,
	})

	if err != nil {
		log.Fatal(err)
	}

	modules.InitDB("postgres", "postgres://dbuser:DBpassword@localhost/resourcedatabase")

	http.Handle("/skillz", handler(schema))
	log.Fatal(http.ListenAndServe(":3002", nil))
}
