package database

import (
	"InternBCC/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("init db failed,", err)
	}
	return DB
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{}, &entity.Gedung{}, &entity.Testimoni{}, &entity.Tag{})
}
