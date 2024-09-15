package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var name, information string
var db *sql.DB

func main() {
	var choose int8
	fmt.Println("Hello!")
	for {
		fmt.Println("You can choose:\n1)New\n2)Get All\n3)Delete\n4)Get one\n5)Exit")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			fmt.Println("Enter name for youre ToDO:")
			fmt.Scan(&name)
			fmt.Println("Enter information:")
			reader := bufio.NewReader(os.Stdin)
			information, _ = reader.ReadString('\n')
			information = information[:len(information)-1]
			err := saveData(name, information)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Your data has been saved!")
				time.Sleep(2 * time.Second)
			}
		case 2:
			fmt.Println("Youre All Info:")
			GetAll()
		}
	}
}

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	createTable()
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS infos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name TEXT,
    	information TEXT
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func saveData(name, information string) error {
	query := `INSERT INTO infos (name, information) VALUES (?, ?)`
	_, err := db.Exec(query, name, information)
	return err
}

func GetAll() {
	query := `SELECT id, name, information FROM infos`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var information string
		err = rows.Scan(&id, &name, &information)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Id: ", id, "Name:", name, "Information:", information)
	}
}
