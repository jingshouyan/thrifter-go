package reflection

import (
	"reflect"
	"unsafe"

	"github.com/jingshouyan/thrifter-go/protocol"
	"github.com/jingshouyan/thrifter-go/spi"
)

type pointerEncoder struct {
	valType    reflect.Type
	valEncoder internalEncoder
}

func (encoder *pointerEncoder) encode(ptr unsafe.Pointer, stream spi.Stream) {
	valPtr := *(*unsafe.Pointer)(ptr)
	if encoder.valType.Kind() == reflect.Map {
		valPtr = *(*unsafe.Pointer)(valPtr)
	}
	encoder.valEncoder.encode(valPtr, stream)
}

func (encoder *pointerEncoder) thriftType() protocol.TType {
	return encoder.valEncoder.thriftType()
}
