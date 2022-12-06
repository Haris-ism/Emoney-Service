package initial

import (
	"gin/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbUrl := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("db connected")
	db.AutoMigrate(&model.Users{}, &model.Balances{},&model.History{})
	return db
}
