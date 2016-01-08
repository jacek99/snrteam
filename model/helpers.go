package model
import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

// Converts a byte array to a pre-existing instance of a Thrift struct
func Thrift2Go(data []byte, entity thrift.TStruct) error {
	transport := thrift.NewTMemoryBufferLen(1024)
	defer transport.Close()

	transport.Read(data)
	return entity.Read(thrift.NewTBinaryProtocolTransport(transport))
}
