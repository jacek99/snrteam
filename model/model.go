package model
import (
	"time"
	"errors"
	"fmt"
)

// custom formats for nice JSON API

type Date time.Time

func (t Date) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(stamp), nil
}
func (m *Date) UnmarshalJSON(data []byte) error {
	if m == nil {
		panic (errors.New("model.Date: UnmarshalJSON on nil pointer"))
	}
	if t, err := time.Parse("2015-01-02",string(data)); err != nil {
		return err
	} else {
		var dt Date = Date(t)
		m = &dt
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
	BirthDate    Date	`json:",string"`
	CreationDate Date	`json:",string"`
	Photo        []byte `json:"-"`
	PwdHash      []byte `json:"-"`
}
