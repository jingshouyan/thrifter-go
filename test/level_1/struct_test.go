package test

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jingshouyan/thrifter-go/general"
	"github.com/jingshouyan/thrifter-go/protocol"
	"github.com/jingshouyan/thrifter-go/raw"
	"github.com/jingshouyan/thrifter-go/test"
	"github.com/jingshouyan/thrifter-go/test/level_1/struct_test"
	"github.com/stretchr/testify/require"
)

func Test_decode_struct_by_iterator(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.I64, 1)
		proto.WriteI64(ctx, 1024)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		iter := c.CreateIterator(buf.Bytes())
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeI64, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(int64(1024), iter.ReadInt64())
		}
		should.NoError(iter.Error())
		should.True(called)
	}
}

func Test_decode_struct_with_bool_by_iterator(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.BOOL, 1)
		proto.WriteBool(ctx, true)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		iter := c.CreateIterator(buf.Bytes())
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeBool, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(true, iter.ReadBool())
		}
		should.True(called)
	}
}

func Test_encode_struct_by_stream(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteStructHeader()
		stream.WriteStructField(protocol.TypeI64, protocol.FieldId(1))
		stream.WriteInt64(1024)
		stream.WriteStructFieldStop()
		iter := c.CreateIterator(stream.Buffer())
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeI64, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(int64(1024), iter.ReadInt64())
		}
	}
}

func Test_encode_struct_with_bool_by_stream(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteStructHeader()
		stream.WriteStructField(protocol.TypeBool, protocol.FieldId(1))
		stream.WriteBool(true)
		stream.WriteStructFieldStop()
		iter := c.CreateIterator(stream.Buffer())
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeBool, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(true, iter.ReadBool())
		}
		should.True(called)
	}
}

func Test_skip_struct(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.I64, 1)
		proto.WriteI64(ctx, 1024)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipStruct(nil))
	}
}

func Test_unmarshal_general_struct(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.I64, 1)
		proto.WriteI64(ctx, 1024)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		var val general.Struct
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val)
	}
}

func Test_unmarshal_raw_struct(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.I64, 1)
		proto.WriteI64(ctx, 1024)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		var val raw.Struct
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(1, len(val))
		should.Equal(protocol.TypeI64, val[protocol.FieldId(1)].Type)
	}
}

func Test_unmarshal_struct(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.I64, 1)
		proto.WriteI64(ctx, 1024)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		var val struct_test.TestObject
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(struct_test.TestObject{1024}, val)
	}
}

func Test_marshal_general_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		output, err := c.Marshal(general.Struct{
			protocol.FieldId(1): int64(1024),
		})
		should.NoError(err)
		var val general.Struct
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, val)
	}
}

func Test_marshal_raw_struct(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.I64, 1)
		proto.WriteI64(ctx, 1024)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		var val raw.Struct
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		output, err := c.Marshal(val)
		should.NoError(err)
		var generalVal general.Struct
		should.NoError(c.Unmarshal(output, &generalVal))
		should.Equal(general.Struct{
			protocol.FieldId(1): int64(1024),
		}, generalVal)
	}
}

func Test_marshal_struct(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(struct_test.TestObject{1024})
		should.NoError(err)
		iter := c.CreateIterator(output)
		called := false
		iter.ReadStructHeader()
		for {
			fieldType, fieldId := iter.ReadStructField()
			if fieldType == protocol.TypeStop {
				break
			}
			should.False(called)
			called = true
			should.Equal(protocol.TypeI64, fieldType)
			should.Equal(protocol.FieldId(1), fieldId)
			should.Equal(int64(1024), iter.ReadInt64())
		}
		should.True(called)
	}
}
