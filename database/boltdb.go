package database

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/jacek99/snrteam/common"
	"log"
	"os"
	"path"
	"time"
)

const (
	RECORD_DOES_NOT_EXIST = 0
	RECORD_ALREADY_EXISTS = 1
	QUERY_ERROR
)

// used for providing more error context on update/delete operations
type WriteError struct {
	ErrorType int
	Message string
}

var Database *bolt.DB = nil

func init() {
	db_folder := os.Getenv(common.ENV_DB_FOLDER)
	if db_folder == "" {
		db_folder = "."
	}

	fmt.Println(Database)

	db_path := path.Join(db_folder, "snrteam.db")

	var err error = nil
	Database, err = bolt.Open(db_path, 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
}
