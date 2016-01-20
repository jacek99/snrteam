package model
import (
	"time"
	"errors"
	"fmt"
	"log"
	"gopkg.in/vmihailenco/msgpack.v2"
)

// custom formats for nice JSON API
type Date time.Time

func (m Date) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(m).Format("2006-01-02"))
	return []byte(stamp), nil
}
func (m *Date) UnmarshalJSON(data []byte) error {
	if m == nil {
		log.Println("model.Date: UnmarshalJSON on nil pointer")
		panic (errors.New("model.Date: UnmarshalJSON on nil pointer"))
	}

	if t, err := time.Parse("\"2006-01-02\"",string(data)); err != nil {
		fmt.Printf("Failed to parse date '%s': %s",string(data),err)
		return err
	} else {
		*m = Date(t)
		return nil
	}
}
func (m *Date) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.EncodeTime(time.Time(*m).UTC())
}

func (m *Date) DecodeMsgpack(dec *msgpack.Decoder) error {
	if t, err := dec.DecodeTime(); err!= nil {
		return err
	} else {
		*m = Date(t.UTC())
		return nil
	}
}


type User struct {
	UserId       int64  `json:"-"`
	UserName     string
	EmailAddress string
	FirstName    string
	LastName     string
	Active       bool
	BirthDate    Date
	CreationDate Date
	Photo        []byte `json:"-"`
	PwdHash      []byte `json:"-"`
}
