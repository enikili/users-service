package database

import (
	"fmt"
	"log"

	"github.com/enikili/users-service/internal/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Используем SQLite - файловая БД, не требует сервера
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	fmt.Println("✅ SQLite database connection established")
	
	// Auto migrate
	err = DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	
	fmt.Println("✅ Database migration completed")
}
