package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"ourbit-service/api"
	"ourbit-service/data"
	"ourbit-service/models/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	db, err = data.InitDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	api.SetupRoutes(router, db)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func loadConfig(filename string) (utils.Config, error) {
	var config utils.Config

	// Use os.ReadFile instead of ioutil.ReadFile
	configFile, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configFile, &config)
	return config, err
}
