package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// Init sets up the database connection
func Init(driver string, host string, port int, user string, password string, database string) (gorm.DB, error) {

	protHost := ""

	if host == "" || host == "127.0.0.1" || host == "localhost" {
		protHost = "unix"
	} else {
		sPort := ""
		if port != 3306 {
			sPort = ":" + fmt.Sprintf("%d", port)
		}

		protHost = "tcp(" + host + sPort + ")"
	}

	connectionString := user + ":" + password + "@" + protHost + "/" + database +
		"?charset=utf8&parseTime=True" //&loc=Local"

	log.Printf("Connecting to %s", connectionString)

	db, err := gorm.Open(driver, connectionString)
	if err != nil {
		fmt.Printf("Connection failed: %s\n", connectionString)
		return db, err
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.DB().Ping()
	if err != nil {
		return db, err
	}

	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)

	return db, nil
}