package belajar_go_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

	querycreateTable := `CREATE TABLE IF NOT EXISTS comments (id INT NOT NULL AUTO_INCREMENT, email VARCHAR(100) NOT NULL, comment TEXT NOT NULL,PRIMARY KEY (id))engine=InnoDB;`

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
	queryctx := `INSERT INTO comments(email, comment) VALUES('john.doe@example.com', 'This is a test comment')`

	_, err := db.ExecContext(ctx, queryctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert comment")
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

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "Jovansa@gmail.com"
	comment := "This is a test comment"
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert comment with ID:", id)
}

func TestSelectComments(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id, email, comment FROM comments")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var email string
		var comment string

		err := rows.Scan(&id, &email, &comment)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id, "\nEmail:", email, "\nComment:", comment)
		fmt.Println("=======================================")
	}
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO comments(email, comment) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "This is comment number " + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success insert comment with ID:", id)
	}

}
