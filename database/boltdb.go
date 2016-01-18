package database

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/jacek99/snrteam/common"
	"log"
	"os"
	"path"
	"time"
	"encoding/binary"
	"bytes"
)

var Database *bolt.DB = nil

const DB_NAME = "snrteam.db"

const (
	user_bucket = "users"
	users_name2id_idx = "users_name2id_idx"
)


func init() {
	createDatabase()
	createBuckets()
}

func createDatabase() {
	db_folder := os.Getenv(common.ENV_DB_FOLDER)
	if db_folder == "" {
		db_folder = "."
	}

	fmt.Println(Database)

	db_path := path.Join(db_folder, DB_NAME)

	var err error = nil
	Database, err = bolt.Open(db_path, 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
}

func createBuckets() {
	tx, err := Database.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	createBucketsIfNotExists(tx, user_bucket, users_name2id_idx)

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

// creates a bucket if missing
func createBucketsIfNotExists(tx *bolt.Tx, buckets... string) {
	for _, bucket := range buckets {
		log.Printf("Creating %s bucket", bucket)
		if _, err := tx.CreateBucketIfNotExists([]byte(bucket)); err != nil {
			log.Fatalf("Unable to creat bucket %s:\n%s",bucket,err)
		}
	}
}

// itob returns an 8-byte big endian representation of v.
func itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(v []byte) int64 {
	var i int64
	buf := bytes.NewReader(v)
	binary.Read(buf, binary.BigEndian, &i)
	return i
}

// gets a bucket by string name
func getBucket(tx *bolt.Tx, bucket string) *bolt.Bucket {
	return tx.Bucket([]byte(bucket))
}

func getString(b *bolt.Bucket, key string) []byte {
	return b.Get([]byte(key))
}

func putString(b *bolt.Bucket, key string, value []byte) error {
	return b.Put([]byte(key), value)
}

func getInt64(b *bolt.Bucket, key int64) []byte {
	return b.Get(itob(key))
}

func putInt64(b *bolt.Bucket, key int64, value []byte) error {
	return b.Put(itob(key), value)
}

// TODO: protect
func Truncate() {
	Database.Close()

	os.Remove(DB_NAME)

	createDatabase()
	createBuckets()

}