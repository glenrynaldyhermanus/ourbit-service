package seller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func SetupSellerRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/seller/registration", RegisterSeller).Methods("POST")
}
