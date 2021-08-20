package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaBd() *gorm.DB {
	/* TimeZone=Asia/Shanghai */
	dsn := "host=127.0.0.1 user=postgres password=123456 dbname=lojaTest port=5432 sslmode=disable"
	// https://github.com/jackc/pgx
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
