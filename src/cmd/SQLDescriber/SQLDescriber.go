package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	_ "github.com/go-sql-driver/mysql"
)

var (
	usePassword bool
	username    string
	database    string
	host        string
	port        string
)

func init() {
	flag.BoolVar(&usePassword, "p", false, "Specify that you wish to provide a password (You will be prompted)")
	flag.StringVar(&username, "u", "root", "Username to connect with")
	flag.StringVar(&database, "db", "", "Database to use")
	flag.StringVar(&host, "host", "localhost", "Host to connect to")
	flag.StringVar(&port, "port", "3306", "Port to connect to")
}

func getPassword() string {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalln("Failed to read password (ABORTED)")
	}
	return strings.TrimSpace(string(bytePassword))
}

func main() {
	// Parse the args
	flag.Parse()
	if len(database) == 0 {
		log.Fatalln("A database is required")
	}

	password := ""
	if usePassword {
		password = getPassword()
	}
	// Open connection to database
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		log.Fatalln(err)
	}
	// Print the tables
	fmt.Println()
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatalln(err)
	}

	// Get the tables
	var tables []string
	for rows.Next() {
		var tableName string
		_ = rows.Scan(&tableName)
		tables = append(tables, tableName)
	}

	for _, val := range tables {
		rows, err = db.Query(fmt.Sprintf("SHOW CREATE TABLE %s", val))
		if err != nil {
			log.Fatalln(err)
		}

		for rows.Next() {

			var table string
			var createString string
			_ = rows.Scan(&table, &createString)
			fmt.Println(createString)
		}
		fmt.Println()
	}

}
