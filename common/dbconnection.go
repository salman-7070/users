package common

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DBConnection *sql.DB

func GetDBConnection() (*sql.DB, bool) {

	err := DBConnection.Ping()

	if err != nil {

		log.Println("Error in Database PING : ", err)

		myDb, result := OpenConnection()

		if result {
			DBConnection = myDb
			return DBConnection, true
		} else {
			return DBConnection, false
		}
	}

	return DBConnection, true
}

func init() {

	log.Println("Connection Initializing... Init()")

	OpenConnection()
}

func OpenConnection() (*sql.DB, bool) {

	//Here you can add your database connection string or set it in a JSON file and marshal it.
	connectionString := "user=postgres dbname=inventory password=8080 host=localhost port=5432 sslmode=disable"

	myDb, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("Open Database Connection failed!")
		return nil, false
	}

	DBConnection = myDb

	return myDb, true
}
