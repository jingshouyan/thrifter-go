package test

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jingshouyan/thrifter/general"
	"github.com/jingshouyan/thrifter/protocol"
	"github.com/jingshouyan/thrifter/test"
	"github.com/stretchr/testify/require"
)

func Test_skip_message(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMessageBegin(ctx, "hello", thrift.CALL, 17)
		proto.WriteStructBegin(ctx, "args")
		proto.WriteFieldBegin(ctx, "field1", thrift.I64, 1)
		proto.WriteI64(ctx, 1)
		proto.WriteFieldBegin(ctx, "field2", thrift.I64, 2)
		proto.WriteI64(ctx, 2)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		proto.WriteMessageEnd(ctx)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipStruct(iter.SkipMessageHeader(nil)))
	}
}

func Test_unmarshal_message(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMessageBegin(ctx, "hello", thrift.CALL, 17)
		proto.WriteStructBegin(ctx, "args")
		proto.WriteFieldBegin(ctx, "field1", thrift.I64, 1)
		proto.WriteI64(ctx, 1)
		proto.WriteFieldBegin(ctx, "field2", thrift.I64, 2)
		proto.WriteI64(ctx, 2)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		proto.WriteMessageEnd(ctx)
		var msg general.Message
		should.NoError(c.Unmarshal(buf.Bytes(), &msg))
		should.Equal("hello", msg.MessageName)
		should.Equal(protocol.MessageTypeCall, msg.MessageType)
		should.Equal(protocol.SeqId(17), msg.SeqId)
		should.Equal(int64(1), msg.Arguments[protocol.FieldId(1)])
		should.Equal(int64(2), msg.Arguments[protocol.FieldId(2)])
	}
}

func Test_marshal_message(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		output, err := c.Marshal(general.Message{
			MessageHeader: protocol.MessageHeader{
				MessageType: protocol.MessageTypeCall,
				MessageName: "hello",
				SeqId:       protocol.SeqId(17),
			},
			Arguments: general.Struct{
				protocol.FieldId(1): int64(1),
				protocol.FieldId(2): int64(2),
			},
		})
		should.NoError(err)
		var msg general.Message
		should.NoError(c.Unmarshal(output, &msg))
		should.Equal("hello", msg.MessageName)
		should.Equal(protocol.MessageTypeCall, msg.MessageType)
		should.Equal(protocol.SeqId(17), msg.SeqId)
		should.Equal(int64(1), msg.Arguments[protocol.FieldId(1)])
		should.Equal(int64(2), msg.Arguments[protocol.FieldId(2)])
	}
}
