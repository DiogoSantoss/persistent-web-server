package middleware

import (
	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/models"

	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// Global database connection
var db *gorm.DB

// Create connection to database and migrate schema
func CreateConnection() *gorm.DB {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// Load env variables to local variables
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	// Open connection to database
	db, err = gorm.Open(dialect, "host="+host+" port="+port+" user="+user+" dbname="+dbName+" password="+password+" sslmode=disable")
	if err != nil {
		// If ssl error occurs, try disabling ssl
		// by adding sslmode=disable to the connection string
		log.Fatal(err)
	} else {
		log.Println("Connected to database at " + host + ":" + port)
	}
 
	// Migrate the schema
	db.AutoMigrate(&models.Data{})

	// Return db connection to close it later
	return db
}

// Add data to database
func addDataDatabase(data models.Data){
	db.Create(&data)
}

// Get data from database
func getDataDatabase() []models.Data {
	var data []models.Data
	db.Find(&data)
	return data
}