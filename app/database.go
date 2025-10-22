package app

import (
	"database/sql"
	"go-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	//db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/belajar_golang_restful_api")
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/belajar_golang_database_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	//migrate create -ext sql -dir db/migrations create_table_first
	//migrate create -ext sql -dir db/migrations create_table_second
	//migrate create -ext sql -dir db/migrations create_table_third

	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up
	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down

	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up (angka)
	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down (angka)

	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations version
	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations force 20251022140403
}
