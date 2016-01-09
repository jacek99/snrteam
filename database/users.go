package database

import (
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

func thrift2User(v []byte) *model.User {
	user := model.NewUser()
	model.Thrift2Go(v,user)
	return user
}

func GetAllUsers() ([]model.User, error) {

	users := []model.User{}

	err := Database.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(user_bucket_name))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			users = append(users, *thrift2User(v))
		}

		return nil
	})
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

// Saves a user, if it exists error occurs
func SaveUser(user model.User) error {
	return nil
}
