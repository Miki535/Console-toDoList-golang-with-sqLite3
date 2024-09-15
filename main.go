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
var id int
var db *sql.DB

func main() {
	var choose1, choose2 int8
	fmt.Println("Hello!")
	for {
		fmt.Println("You can choose:\n1)New\n2)Get All\n3)Delete\n4)Get one\n5)Delete All\n6)Exit")
		fmt.Scan(&choose1)
		switch choose1 {
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
		case 3:
			fmt.Println("Choose:\n1)Delete by name\n2)Delete by id\n3)Exit")
			fmt.Scan(&choose2)
			switch choose2 {
			case 1:
				fmt.Println("Enter name for today:")
				fmt.Scan(&name)
				deleteByName(name)
			case 2:
				fmt.Println("Enter id for today:")
				fmt.Scan(&id)
				deleteById(id)
			case 3:
				fmt.Println("Youre Welcome!")
			}
		case 4:
			fmt.Println("Choose:\n1)Get by name\n2)Get by id\n3)Exit")
			fmt.Scan(&choose2)
			switch choose2 {
			case 1:
				fmt.Println("Enter name for today:")
				fmt.Scan(&name)
				getByName(name)
			case 2:
				fmt.Println("Enter id for today:")
				fmt.Scan(&id)
				getById(id)
			case 3:
			}

		case 5:
			deletAll()
		case 6:
			fmt.Println("Bye!")
			os.Exit(0)
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
		err = rows.Scan(&id, &name, &information)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Id: ", id, "Name:", name, "Information:", information)
	}
}

func deletAll() {
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
