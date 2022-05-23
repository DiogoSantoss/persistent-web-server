package middleware

import (
	"gitlab.com/psem/recruitment-software/diogosantoss/persistent-web-server/models"

	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"github.com/joho/godotenv"
)

// Global database connection
var db *gorm.DB

// Create connection to database and migrate schema
func CreateConnection() *gorm.DB {

	/* 
	This is no longer necessary because all env variables
	are now set in the docker-compose file.

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	*/
	
	// Load env variables to local variables
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	
	var err error

	retries := 5
	for i := 0; i < retries; i++ {
		// Open connection to database
		db, err = gorm.Open(dialect, "host="+host+" port="+port+" user="+user+" dbname="+dbName+" password="+password+" sslmode=disable")
		if err != nil {
			if i == retries-1 {
				log.Fatal("Error connecting to database with error: ", err)
			}
			log.Printf("Error connecting to database, retrying... (%v/%v)\n",i+1,retries)
			// Wait for 1 second before retrying
			time.Sleep(time.Second)
		} else {
			log.Println("Connected to database at " + host + ":" + port)
			break
		}
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