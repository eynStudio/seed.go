package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/eynstudio/seed.go/base/db"
	"github.com/eynstudio/seed.go/base/utils"
	"github.com/eynstudio/seed.go/schema"
	"github.com/graphql-go/handler"
)

func init() {

}

func main() {
	loadConfig()

	log.Println(viper.Get("hi.hi"))

	mongo := db.OpenMongo()
	defer mongo.Close()

	h := handler.New(&handler.Config{
		Schema:   &schema.RootSchema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", utils.Cors(h))

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
