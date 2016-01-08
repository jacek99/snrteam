package database

import (
	"container/list"
	"github.com/boltdb/bolt"
	"github.com/jacek99/snrteam/model"
	"log"
)

const user_bucket_name = "users"

func init() {
	tx, err := Database.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	log.Println("Creating users bucker")
	_, err = tx.CreateBucketIfNotExists([]byte(user_bucket_name))
	if err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func GetAllUsers() ([]model.User, error) {

	l := list.New()

	err := Database.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(user_bucket_name))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			l.PushBack(model.Thrift2Go(v,model.NewUser()))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// convert list to array
	users := make([]model.User, l.Len())
	index := 0
	for e := l.Front(); e != nil; e = e.Next() {
		users[index] = e.Value.(model.User)
		index++
	}
	return users, nil
}

// Saves a user, if it exists error occurs
func SaveUser(user model.User) error {
	return nil
}
