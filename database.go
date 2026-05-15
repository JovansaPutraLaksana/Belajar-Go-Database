package belajar_go_database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "u7jrrwm6d2do347l:jD4C66nOGnoFxyr3Vxws@tcp(blzvlgzphf7oinbjckkg-mysql.services.clever-cloud.com:3306)/blzvlgzphf7oinbjckkg")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(1)                   //minimal koneksi yang bisa dibuka
	db.SetMaxOpenConns(5)                   //maksimal koneksi yang bisa dibuka
	db.SetConnMaxIdleTime(5 * time.Minute)  //waktu maksimal koneksi idle sebelum ditutup
	db.SetConnMaxLifetime(60 * time.Minute) //waktu maksimal koneksi hidup sebelum ditutup

	return db
}
