package belajar_go_database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	fmt.Println("Test Empty")
}

// koneksi sql ini bersifat rahasia (saya tidak menggunakan localhost karena laptop saya tidak terinstall mysql), namun ini adalah sql dummy yang bisa digunakan untuk belajar, jadi tidak masalah jika dibagikan
func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "u7jrrwm6d2do347l:jD4C66nOGnoFxyr3Vxws@tcp(blzvlgzphf7oinbjckkg-mysql.services.clever-cloud.com:3306)/belajar-go-database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
