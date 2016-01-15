package database

import (
	"github.com/boltdb/bolt"
	"github.com/jacek99/snrteam/model"
	"log"
	"github.com/jacek99/snrteam/common"
)

const (
	user_bucket = "users"
	users_name2id_idx = "users_name2id_idx" //
)

func init() {
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

func thrift2User(v []byte) *model.User {
	user := model.NewUser()
	model.Thrift2Go(v,user)
	return user
}

func GetAllUsers() ([]*model.User, error) {

	users := []*model.User{}

	err := Database.View(func(tx *bolt.Tx) error {

		b := getBucket(tx,user_bucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			users = append(users, thrift2User(v))
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
func GetUser(userId int64) (*model.User, error) {

	var user *model.User

	err := Database.View(func(tx *bolt.Tx) error {

		b := getBucket(tx,user_bucket)
		if data := getInt64(b,userId); data != nil {
			user = thrift2User(data)
			return nil
		} else {
			return common.RECORD_NOT_FOUND_ERROR
		}
	})
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}

}

// may return null if not found
func GetUserByName(userName string) (*model.User, error) {

	var user *model.User

	err := Database.View(func(tx *bolt.Tx) error {

		b := getBucket(tx,users_name2id_idx)
		if data := getString(b,userName); data != nil {
			user, _= GetUser(btoi(data))
			return nil
		} else {
			return common.RECORD_NOT_FOUND_ERROR
		}
	})

	return user, err
}

// Saves a user, if it exists error occurs
func SaveUser(user *model.User)  error  {
	existing, err := GetUser(user.UserId)
	if err != nil {
		return err
	}

	if existing == nil {

		return Database.Update(func(tx *bolt.Tx) error {
			// put both user by ID as well as the index by name
			b := getBucket(tx, user_bucket)
			idx := getBucket(tx, users_name2id_idx)

			if err := putInt64(b, user.UserId, model.Go2Thrift(user));err != nil {
				return err
			} else {
				if err = putString(idx,user.UserName, itob(user.UserId)); err != nil {
					return err
				}
			}

			return nil
		})

	} else {
		// record already exists
		return common.RECORD_ALREADY_EXISTS_ERROR
	}}
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


