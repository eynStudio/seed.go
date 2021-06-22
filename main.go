package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/eynstudio/seed.go/base/db"
	"github.com/eynstudio/seed.go/schema"
	"github.com/graphql-go/graphql"
)

func init() {

}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	loadConfig()

	log.Println(viper.Get("hi.hi"))

	mongo := db.OpenMongo()
	defer mongo.Close()

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema.RootSchema)
		json.NewEncoder(w).Encode(result)
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(":"+viper.GetString("port"), nil)
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}
