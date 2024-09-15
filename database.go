package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

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
		err = rows.Scan(&id, &name, &information)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Id: ", id, "Name:", name, "Information:", information)
	}
}

func deleteAll() {
	query := `DELETE FROM infos`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteByName(name string) {
	query := `DELETE FROM infos WHERE name = ?`
	_, err := db.Exec(query, name)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Deleted info:", name)
	}
}

func deleteById(id int) {
	query := `DELETE FROM infos WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Deleted info by id:", id)
	}
}

func getByName(name string) {
	query := `SELECT information FROM infos WHERE name = ?`
	rows, err := db.Query(query, name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&information)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(information)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func getById(id int) {
	query := `SELECT information FROM infos WHERE id = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&information)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(information)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
