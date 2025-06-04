package initializers

import (
	"log"
	"quickfit-server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:rootpassword@tcp(127.0.0.1:3306)/quickfit"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
}

func MigrateDatabase() {
	err := DB.AutoMigrate(
		models.Muscle{},
		models.ExerciseCategory{},
		models.Equipment{},
		models.Exercise{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}
}
