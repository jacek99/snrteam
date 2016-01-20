package model
import "gopkg.in/vmihailenco/msgpack.v2"

// marshalls entity to  msgpack
func Marshall(entity interface{}) []byte {
	b, err := msgpack.Marshal(true)
	if err != nil {
		panic(err)
	}
	return b
}

func Unmarshall(entity interface{},data []byte) (interface{}) {
	err := msgpack.Unmarshal(data, entity)
	if err != nil {
		panic(err)
	}
	return entity
}
