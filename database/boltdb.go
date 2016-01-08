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
