package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	DBUser     = "postgres"
// 	DBPassword = "password"
// 	DBName     = "todo_db"
// 	DBHost     = "127.0.0.1"
// 	DBPort     = "5432"
// 	DBType     = "postgres"
// )

const (
	DBUser     = "jxwfsdrmgkodop"
	DBPassword = "9f695dce1de225f9c1caf7d6ea1fffc9456f9f31a2108ccd3b99a310801703ba"
	DBName     = "df3k9vvjafhfb5"
	DBHost     = "ec2-44-196-170-156.compute-1.amazonaws.com"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	database := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require TimeZone=Asia/Jakarta",
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
