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

// may return null if not found
func GetUser(userId string) (*model.User, error) {

	var user *model.User

	err := Database.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(user_bucket_name))
		data := b.Get([]byte(userId))
		if data != nil {
			user = thrift2User(data)
		}

		return nil
	})
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}

}

//// Saves a user, if it exists error occurs
//func SaveUser(user model.User)  WriteError  {
//	user, err := GetUser(user.UserId)
//	if err != nil {
//		return nil, WriteError{QUERY_ERROR, err}
//	}
//
//	if user == nil {
//		err = Database.Update(func(tx *bolt.Tx) error {
//			b := tx.Bucket([]byte(user_bucket_name))
//			return b.Put([]byte(user.UserId),model.Go2Thrift(user))
//		})
//		return user, err
//	} else {
//		// record already exists
//		return nil, WriteError{RECORD_ALREADY_EXISTS, nil}
//	}}
//
//// update an existing user
//func UpdateUser(user *model.User) (*model.User, error) {
//	user, err := GetUser(user.UserId)
//	if err != nil {
//		return nil, error
//	}
//
//	if user != nil {
//		err = Database.Update(func(tx *bolt.Tx) error {
//			b := tx.Bucket([]byte(user_bucket_name))
//			return b.Put([]byte(user.UserId),model.Go2Thrift(user))
//		})
//		return user, err
//	} else {
//		// TODO: need to return error type
//		return nil, "User does not exist"
//	}
//}


