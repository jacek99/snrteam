package model
import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"reflect"
	"log"
)

// Converts a byte array to a pre-existing instance of a Thrift struct
func Thrift2Go(data []byte, entity thrift.TStruct) thrift.TStruct {
	transport := thrift.NewTMemoryBufferLen(1024)
	defer transport.Close()

	transport.Read(data)
	err := entity.Read(thrift.NewTBinaryProtocolTransport(transport))

	// an error should never occur, means Thrift corruption occurred!
	if err != nil {
		panic(fmt.Sprintf("Error during Thrift deserialization of type %s: %s", reflect.TypeOf(entity),err))
	} else {
		return entity
	}
}

func Go2Thrift(entity thrift.TStruct) []byte {

	t := thrift.NewTSerializer()
	transport := thrift.NewTMemoryBufferLen(1024)
	defer transport.Close()
	t.Protocol = thrift.NewTBinaryProtocolTransport(transport)

	if data, err := t.Write(entity); err != nil {
		panic(fmt.Sprintf("Error during Thrift serialization of type %s: %s", reflect.TypeOf(entity),err))
	} else if (len(data) == 0) {
		panic(fmt.Sprintf("Error during Thrift serialization of type %s: %s", reflect.TypeOf(entity),"0 bytes serialized!"))
	} else {
		return data
	}
}


