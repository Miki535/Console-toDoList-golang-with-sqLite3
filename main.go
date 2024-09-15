package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"time"
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
