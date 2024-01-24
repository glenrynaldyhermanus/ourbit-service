package data

import (
	"database/sql"
	"fmt"
	"log"
	"ourbit-service/models"
	"ourbit-service/models/utils"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB(config utils.Config) (*sql.DB, error) {
	connectionString := getConnectionString(config)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func getConnectionString(config utils.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

}

// InsertBusiness inserts a new business into the database
func InsertBusiness(business models.Business) (int, error) {
	result, err := DB.Exec(`
		INSERT INTO businesses (name, category, username)
		VALUES (?, ?, ?)
	`, business.Name, business.Category, business.Username)
	if err != nil {
		log.Printf("Error inserting business: %v\n", err)
		return 0, err
	}

	// Retrieve the ID of the inserted business
	businessID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error retrieving last insert ID: %v\n", err)
		return 0, err
	}

	return int(businessID), nil
}

// InsertSeller inserts a new seller into the database
func InsertSeller(seller models.Seller) (int, error) {
	result, err := DB.Exec(`
		INSERT INTO sellers (uuid, created_at, updated_at, is_active, full_name, email, phone, password, is_phone_active, is_email_active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, seller.UUID, seller.CreatedAt, seller.UpdatedAt, seller.IsActive, seller.FullName, seller.Email, seller.Phone, seller.Password, seller.IsPhoneActive, seller.IsEmailActive)
	if err != nil {
		log.Printf("Error inserting seller: %v\n", err)
		return 0, err
	}

	// Retrieve the ID of the inserted seller
	sellerID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error retrieving last insert ID: %v\n", err)
		return 0, err
	}

	return int(sellerID), nil
}

// InsertStore inserts a new store into the database
func InsertStore(store models.Store) (int, error) {
	result, err := DB.Exec(`
		INSERT INTO stores (name, address, coordinate)
		VALUES (?, ?, ?)
	`, store.Name, store.Address, store.Coordinate)
	if err != nil {
		log.Printf("Error inserting store: %v\n", err)
		return 0, err
	}

	// Retrieve the ID of the inserted store
	storeID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error retrieving last insert ID: %v\n", err)
		return 0, err
	}

	return int(storeID), nil
}
