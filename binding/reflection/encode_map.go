package reflection

import (
	"reflect"
	"unsafe"

	"github.com/jingshouyan/thrifter-go/protocol"
	"github.com/jingshouyan/thrifter-go/spi"
)

type mapEncoder struct {
	mapInterface emptyInterface
	keyEncoder   internalEncoder
	elemEncoder  internalEncoder
}

func (encoder *mapEncoder) encode(ptr unsafe.Pointer, stream spi.Stream) {
	mapInterface := encoder.mapInterface
	mapInterface.word = ptr
	realInterface := (*interface{})(unsafe.Pointer(&mapInterface))
	mapVal := reflect.ValueOf(*realInterface)
	keys := mapVal.MapKeys()
	stream.WriteMapHeader(encoder.keyEncoder.thriftType(), encoder.elemEncoder.thriftType(), len(keys))
	for _, key := range keys {
		keyObj := key.Interface()
		keyInf := (*emptyInterface)(unsafe.Pointer(&keyObj))
		encoder.keyEncoder.encode(keyInf.word, stream)
		elem := mapVal.MapIndex(key)
		elemObj := elem.Interface()
		elemInf := (*emptyInterface)(unsafe.Pointer(&elemObj))
		encoder.elemEncoder.encode(elemInf.word, stream)
	}
}

func (encoder *mapEncoder) thriftType() protocol.TType {
	return protocol.TypeMap
}
