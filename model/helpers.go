package model
import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"reflect"
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
	transport := thrift.NewTMemoryBufferLen(1024)
	defer transport.Close()

	protocol := thrift.NewTBinaryProtocolTransport(transport)
	err := entity.Write(protocol)
	if err != nil {
		panic(fmt.Sprintf("Error during Thrift serialization of type %s: %s", reflect.TypeOf(entity),err))
	} else {
		return transport.Bytes()
	}
}


