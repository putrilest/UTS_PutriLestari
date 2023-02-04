package connection

import (
	"fmt"
	"log"
	"os"
	"projek-restoran/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatalln(" Error to Load Env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't Connect to Database")
	}

	db.AutoMigrate(&models.Kategori{})
	db.AutoMigrate(&models.Menu{})
	db.AutoMigrate(&models.Meja{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Reservasi{})
	db.AutoMigrate(&models.Kasir{})

	return db

}
