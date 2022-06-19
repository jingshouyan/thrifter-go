package raw

import (
	"github.com/jingshouyan/thrifter-go/protocol"
	"github.com/jingshouyan/thrifter-go/spi"
)

type rawListEncoder struct {
}

func (encoder *rawListEncoder) Encode(val interface{}, stream spi.Stream) {
	obj := val.(List)
	stream.WriteListHeader(obj.ElementType, len(obj.Elements))
	for _, elem := range obj.Elements {
		stream.Write(elem)
	}
}

func (encoder *rawListEncoder) ThriftType() protocol.TType {
	return protocol.TypeList
}
