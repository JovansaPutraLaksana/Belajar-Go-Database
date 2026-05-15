package belajar_go_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// exec sql untuk insert data
	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES('c001', 'Jovansa')")
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
