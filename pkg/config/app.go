package config

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connect() {
	myDB_dsn := "host=localhost user=haydn password=tryguess dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	sqlDB, err := sql.Open("pgx", myDB_dsn)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}