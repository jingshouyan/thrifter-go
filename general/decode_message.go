package general

import (
	"github.com/jingshouyan/thrifter/protocol"
	"github.com/jingshouyan/thrifter/spi"
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
