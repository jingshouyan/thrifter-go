package general

import (
	"github.com/thrift-iterator/go/spi"
	"github.com/thrift-iterator/go/protocol"
)

type generalListEncoder struct {
}

func (encoder *generalListEncoder) Encode(val interface{}, stream spi.Stream) {
	writeList(val, stream)
}

func writeList(val interface{}, stream spi.Stream) {
	obj := val.([]interface{})
	length := len(obj)
	if length == 0 {
		stream.WriteListHeader(protocol.TypeI64, 0)
		return
	}
	elemType, generalWriter := generalWriterOf(obj[0])
	stream.WriteListHeader(elemType, length)
	for _, elem := range obj {
		generalWriter(elem, stream)
	}
}