package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBUser     = "postgres"
	DBPassword = "password"
	DBName     = "todo_db"
	DBHost     = "127.0.0.1"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	database := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Jakarta",
		DBHost, DBPort, DBUser, DBName, DBPassword,
	)

	return database
}

var DB *gorm.DB

//NewDB to initiate Database Connection
func NewDB(params ...string) *gorm.DB {
	var err error
	conString := GetPostgresConnectionString()

	log.Print(conString)

	DB, err := gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return DB
}
