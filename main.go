package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial Start")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	//db, err := sql.Open("mssql", "test_user:123@tcp(127.0.0.1:51329)/NBA_TOPSHOT")
	db, err := sql.Open("sqlserver", "server=127.0.0.1;user id=test_user;password=123;port=51329;database=NBA_TOPSHOT;")
	fmt.Println("Connection Successful")
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	//defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO users VALUES('JOHN')")
	fmt.Println("Insert Successful")
	// if there is an error inserting, handle it
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	//defer insert.Close()
	insert.Close()
	db.Close()
}
