package database

import (
	"github.com/boltdb/bolt"
	"github.com/jacek99/snrteam/model"
	"github.com/jacek99/snrteam/common"
	"github.com/nicksnyder/go-i18n/i18n"
	"golang.org/x/crypto/bcrypt"

	"log"
	"time"
)

const USER_NAME = "UserName"
const USER_ID = "UserId"

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

func GetAllUsers() ([]*model.User, error) {

	users := []*model.User{}

	err := Database.View(func(tx *bolt.Tx) error {

		b := getBucket(tx,user_bucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			users = append(users, model.Unmarshall(v,new(model.User)).(*model.User))
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
func GetUser(userId int64, T i18n.TranslateFunc) (*model.User, error) {

	var user *model.User

	err := Database.View(func(tx *bolt.Tx) error {

		b := getBucket(tx,user_bucket)
		if data := getInt64(b,userId); data != nil {
			user = model.Unmarshall(data,new(model.User)).(*model.User)
			return nil
		} else {
			return common.NotFoundError{T("user_id_not_found", userId),"User",USER_ID,userId}
		}
	})
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}

}

// may return null if not found
func GetUserByName(userName string, T i18n.TranslateFunc) (*model.User, error) {

	var user *model.User

	err := Database.View(func(tx *bolt.Tx) error {

		b := getBucket(tx,users_name2id_idx)
		if data := getString(b,userName); data != nil {
			var err error
			user, err = GetUser(btoi(data), T)
			return err
		} else {
			return common.NotFoundError{T("user_not_found", userName),"User",USER_NAME,userName}
		}
	})

	return user, err
}

// Saves a user, if it exists error occurs

func SaveUser(user *model.User, password string, T i18n.TranslateFunc)  error  {

	existing, _ := GetUser(user.UserId, T);

	if existing == nil {

		// convert password to hash
		if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost); err != nil {
			panic(err)
		} else {
			user.PwdHash = hashedPassword
		}

		user.CreationDate = model.Date(time.Now())

		return Database.Update(func(tx *bolt.Tx) error {
			// put both user by ID as well as the index by name
			b := getBucket(tx, user_bucket)
			idx := getBucket(tx, users_name2id_idx)

			// generate numeric user ID
			id, _ := b.NextSequence()
			user.UserId = int64(id)

			// validate before saving
			if err := common.Validate(user,"User"); err != nil {
				return err
			} else {

				// actual DB save
				if err := putInt64(b, user.UserId, model.Marshall(user));err != nil {
					log.Printf("Failed to save to users bucket for %s: %s",user.UserId,err)
					return err
				} else {
					if err = putString(idx,user.UserName, itob(user.UserId)); err != nil {
						log.Printf("Failed to save to user index bucket for user '%s': %s",user.UserName,err)
						return err
					}
				}

			}

			return nil
		})

	} else {
		// record already exists
		return common.ConflictError{T("user_exists", user.UserName),T("user"),USER_NAME,user.UserName}
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


