package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	// connect without specifying a database
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", username, password)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// create the database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS colors")
	if err != nil {
		panic(err)
	}

	// connect to created database
	dsn = "root:rootroot@tcp(127.0.0.1:3306)/colors"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE colors (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL)")
	if err != nil {
		panic(err)
	}

	colors := []string{"Red", "Blue", "Green", "Yellow", "Purple"}
	for _, color := range colors {
		_, err = db.Exec("INSERT INTO colors (name) VALUES (?)", color)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Colors table created and populated!")
}
