package api

import (
	"database/sql"
	"ourbit-service/api/product"
	"ourbit-service/api/seller"

	"github.com/gorilla/mux"
)

// SetupRoutes configures the routes for your API
func SetupRoutes(router *mux.Router, db *sql.DB) {
	product.SetupProductRoutes(router, db)
	seller.SetupSellerRoutes(router, db)
}
