package util

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
