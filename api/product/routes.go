package product

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// Product represents your product structure
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// SetupRoutes configures the routes for your API
func SetupProductRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/products", getProducts(db)).Methods("GET")
	// Add other routes as needed
}

// func getProducts(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		products, err := fetchProducts(db)
// 		if err != nil {
// 			http.Error(w, "500 internal server error", http.StatusInternalServerError)
// 			return
// 		}

// 		if len(products) == 0 {
// 			http.NotFound(w, r)
// 			return
// 		}

// 		respondJSON(w, products)
// 	}
// }

// func fetchProducts(db *sql.DB) ([]Product, error) {
// 	rows, err := db.Query("SELECT id, name FROM products")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var products []Product
// 	for rows.Next() {
// 		var product Product
// 		err := rows.Scan(&product.ID, &product.Name, &product.Price)
// 		if err != nil {
// 			return nil, err
// 		}
// 		products = append(products, product)
// 	}

// 	return products, nil
// }

// func respondJSON(w http.ResponseWriter, data interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }
