package database

import (
	"github.com/jacek99/snrteam/model"
	"log"
	"github.com/boltdb/bolt"
	"container/list"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const user_bucket_name = "user"

func init() {
	tx, err := Database.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	log.Println("Creating user bucker")
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

	err := Database.View(func (tx *bolt.Tx) error {

		b := tx.Bucket([]byte(user_bucket_name))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			transport := thrift.NewTMemoryBuffer()
			defer transport.Close()

			user := model.NewUser()
			transport.Read(v)
			user.Read(thrift.NewTBinaryProtocolTransport(transport))

			l.PushBack(user)
		}

		return nil
    })
	if err !=  nil {
		return nil, err
	}

	users := make([]model.User,l.Len())
	index := 0
	for e := l.Front(); e != nil; e = e.Next() {
		users[index] = e.Value.(model.User)
		index++
	}

	return users, nil
}
