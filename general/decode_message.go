package general

import (
	"github.com/jingshouyan/thrifter-go/protocol"
	"github.com/jingshouyan/thrifter-go/spi"
)

type messageDecoder struct {
}

func (decoder *messageDecoder) Decode(val interface{}, iter spi.Iterator) {
	*val.(*Message) = Message{
		MessageHeader: iter.ReadMessageHeader(),
		Arguments:     readStruct(iter).(Struct),
	}
}

type messageHeaderDecoder struct {
}

func (decoder *messageHeaderDecoder) Decode(val interface{}, iter spi.Iterator) {
	*val.(*protocol.MessageHeader) = iter.ReadMessageHeader()
}
