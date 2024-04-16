package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB インスタンスをパッケージレベルで保持
var DB *gorm.DB

// InitDatabase はデータベース接続を初期化します
func InitDatabase() {
    dsn := "root:koryo5040@tcp(localhost:3306)/giver?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    DB = db
    log.Println("Database connected")
}