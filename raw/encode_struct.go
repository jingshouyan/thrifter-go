package raw

import (
	"github.com/jingshouyan/thrifter/protocol"
	"github.com/jingshouyan/thrifter/spi"
)

type rawStructEncoder struct {
}

func (encoder *rawStructEncoder) Encode(val interface{}, stream spi.Stream) {
	obj := val.(Struct)
	stream.WriteStructHeader()
	for fieldId, field := range obj {
		stream.WriteStructField(field.Type, fieldId)
		stream.Write(field.Buffer)
	}
	stream.WriteStructFieldStop()
}

func (encoder *rawStructEncoder) ThriftType() protocol.TType {
	return protocol.TypeStruct
}
