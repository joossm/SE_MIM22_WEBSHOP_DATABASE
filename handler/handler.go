package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := openDB()
		defer closeDB(db)
		fmt.Println("Initialization of the Database was executed")
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS `books` ( `Id` int(11) NOT NULL, `Titel` varchar(45) DEFAULT NULL, `EAN` varchar(45) DEFAULT NULL, `Content` varchar(45) DEFAULT NULL, `Price` float DEFAULT NULL, PRIMARY KEY (`Id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;")
		if err != nil {
			log.Printf("Error creating table: %s", err)
		}
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS `users` ( `Id` int(11) NOT NULL, `Username` varchar(45) DEFAULT NULL, `Password` varchar(45) DEFAULT NULL, `Firstname` varchar(45) DEFAULT NULL, `Lastname` varchar(45) DEFAULT NULL, `Housenumber` varchar(45) DEFAULT NULL, `Street` varchar(45) DEFAULT NULL, `Zipcode` varchar(45) DEFAULT NULL, `City` varchar(45) DEFAULT NULL, `Email` varchar(45) DEFAULT NULL, `Phone` varchar(45) DEFAULT NULL, PRIMARY KEY (`Id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;")
		if err != nil {
			log.Printf("Error creating table: %s", err)
		}
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS `orders` ( `id` int(11) NOT NULL AUTO_INCREMENT, `produktId` varchar(45) DEFAULT NULL, `userId` varchar(45) DEFAULT NULL, `amount` varchar(45) DEFAULT NULL, PRIMARY KEY (`id`) ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;")
		if err != nil {
			log.Printf("Error creating table: %s", err)
		}
		js, err := json.Marshal("Success")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	default:
		js, err := json.Marshal("THIS IS A GET REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

func closeDB(db *sql.DB) {
	err := db.Close()
	errorHandler(err)
}

func openDB() *sql.DB {
	fmt.Println("Opening DB")
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/books")
	fmt.Println(db.Ping())
	fmt.Println(db.Stats())
	db.SetMaxIdleConns(0)
	errorHandler(err)
	return db
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
