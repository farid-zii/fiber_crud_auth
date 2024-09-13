package databases

import (
	"fiber-crud-auth/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(){


	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database: ", err)
    }

    DB = db

    // Perform AutoMigrate or other database setup
    db.AutoMigrate(&models.User{}, &models.Book{})


	
	// cfg := config.GetConfig()
	// dsn := cfg.GetDBConnectionString()

	// var err  error
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil{
	// 	log.Fatal("Failed to connect to the database: ", err)
	// }

	// //
	// DB.AutoMigrate(&models.Book{}, &models.User{})

}