package general

import (
	"github.com/jingshouyan/thrifter-go/protocol"
	"github.com/jingshouyan/thrifter-go/spi"
)

type generalListEncoder struct {
}

func (encoder *generalListEncoder) Encode(val interface{}, stream spi.Stream) {
	writeList(val, stream)
}

func (encoder *generalListEncoder) ThriftType() protocol.TType {
	return protocol.TypeList
}

func writeList(val interface{}, stream spi.Stream) {
	obj := val.(List)
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
