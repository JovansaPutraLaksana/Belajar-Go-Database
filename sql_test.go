package belajar_go_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// exec sql untuk insert data
	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name,email,balance,rating,birth_date,married) VALUES('c002', 'Putra',null,1500000,5,null,false)")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")

	// exec sql untuk update data
	// _, err = db.ExecContext(ctx, "UPDATE customer SET name = 'Eko Kurniawan' WHERE id = 'c001'")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Success Update Data")

	// exec sql untuk delete data
	// _, err = db.ExecContext(ctx, "DELETE FROM customer WHERE id = 'c001'")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Success Delete Data")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		var email sql.NullString
		var balance int
		var rating float64
		var birthDate sql.NullString
		var married bool
		var createdAt time.Time

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id, "\nName:", name, "\nEmail:", email.String, "\nBalance:", balance, "\nRating:", rating, "\nBirth Date:", birthDate.String, "\nMarried:", married, "\nCreated At:", createdAt)
		fmt.Println("=======================================")
	}
}

func TestCreateTable(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querycreateTable := `CREATE TABLE IF NOT EXISTS user (username VARCHAR(50), password VARCHAR(50) NOT NULL,PRIMARY KEY (username))engine=InnoDB;`

	_, err := db.ExecContext(ctx, querycreateTable)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success create table")
}

func TestInsertUser(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO user(username, password) VALUES('admin', 'admin')")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert user")
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"       // username yang dimasukkan oleh user wajib benar tanpa perlu password, karena akan diabaikan oleh query
	password := "admin or '1'='1" // password yang dimasukkan oleh user wajib benar tanpa perlu username, karena akan diabaikan oleh query
	query := fmt.Sprintf("SELECT username FROM user WHERE username = '%s' AND password = '%s' LIMIT 1", username, password)
	fmt.Println("Query:", query)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success, Welcome", username)
	} else {
		fmt.Println("Login Failed")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin '; #"
	password := "admin"
	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1" // menggunakan parameter untuk menghindari SQL Injection
	fmt.Println("Query:", query)
	rows, err := db.QueryContext(ctx, query, username, password) // menggunakan parameter untuk menghindari SQL Injection
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success, Welcome", username)
	} else {
		fmt.Println("Login Failed")
	}
}
