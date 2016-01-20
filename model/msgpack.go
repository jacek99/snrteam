package model
import (
	"gopkg.in/vmihailenco/msgpack.v2"
	"fmt"
	"errors"
)

// marshalls entity to  msgpack
func Marshall(entity interface{}) []byte {
	b, err := msgpack.Marshal(entity)
	if err != nil {
		panic(err)
	}
	return b
}

func Unmarshall(data []byte, entity interface{}) (interface{}) {
	err := msgpack.Unmarshal(data, &entity)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to unmarshall: %s (data bytes: %d)",err,len(data))))
	}
	return entity
}
