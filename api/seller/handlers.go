package seller

import (
	"encoding/json"
	"log"
	"net/http"
	"ourbit-service/data"
	"ourbit-service/models"
	"ourbit-service/models/request"
)

// RegisterSeller handles seller registration
func RegisterSeller(w http.ResponseWriter, r *http.Request) {

	var request request.SellerRegistrationRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Insert seller into the database
	sellerID, err := data.InsertSeller(models.Seller{FullName: request.FullName, Email: request.Email, Phone: request.Phone, Password: request.Password})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error inserting seller")
		return
	}

	// Insert business into the database
	business := models.Business{Name: request.BusinessName, Username: request.BusinessCategory, Category: request.BusinessCategory}
	businessID, err := data.InsertBusiness(business)
	if err != nil {
		// Handle error, possibly rollback seller insertion
		RespondWithError(w, http.StatusInternalServerError, "Error inserting business")
		return
	}

	// Automatically create a new store (you may need to implement store creation logic)
	store := models.Store{Name: "New Store", BusinessID: businessID, SellerID: sellerID}
	storeID, err := data.InsertStore(store)
	_ = storeID

	if err != nil {
		// Handle error, possibly rollback seller and business insertion
		RespondWithError(w, http.StatusInternalServerError, "Error creating store")
		return
	}

	RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Seller registered successfully"})
}

// RespondWithError sends a JSON response with an error message and the specified HTTP status code
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, map[string]string{"error": message})
}

// RespondWithJSON sends a JSON response with the specified HTTP status code and payload
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Printf("Error encoding JSON response: %v\n", err)
		// You may choose to log this error or handle it differently based on your requirements
	}
}
