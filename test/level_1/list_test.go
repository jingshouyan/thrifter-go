package test

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jingshouyan/thrifter/general"
	"github.com/jingshouyan/thrifter/protocol"
	"github.com/jingshouyan/thrifter/raw"
	"github.com/jingshouyan/thrifter/test"
	"github.com/stretchr/testify/require"
)

func Test_decode_list_by_iterator(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.I64, 3)
		proto.WriteI64(ctx, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteI64(ctx, 3)
		proto.WriteListEnd(ctx)
		iter := c.CreateIterator(buf.Bytes())
		elemType, length := iter.ReadListHeader()
		should.Equal(protocol.TypeI64, elemType)
		should.Equal(3, length)
		should.Equal(uint64(1), iter.ReadUint64())
		should.Equal(uint64(2), iter.ReadUint64())
		should.Equal(uint64(3), iter.ReadUint64())
	}
}

func Test_encode_list_by_stream(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteListHeader(protocol.TypeI64, 3)
		stream.WriteUint64(1)
		stream.WriteUint64(2)
		stream.WriteUint64(3)
		iter := c.CreateIterator(stream.Buffer())
		elemType, length := iter.ReadListHeader()
		should.Equal(protocol.TypeI64, elemType)
		should.Equal(3, length)
		should.Equal(uint64(1), iter.ReadUint64())
		should.Equal(uint64(2), iter.ReadUint64())
		should.Equal(uint64(3), iter.ReadUint64())
	}
}

func Test_skip_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.I64, 3)
		proto.WriteI64(ctx, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteI64(ctx, 3)
		proto.WriteListEnd(ctx)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipList(nil))
	}
}

func Test_unmarshal_general_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.I64, 3)
		proto.WriteI64(ctx, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteI64(ctx, 3)
		proto.WriteListEnd(ctx)
		var val general.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.List{int64(1), int64(2), int64(3)}, val)
	}
}

func Test_unmarshal_raw_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.I64, 3)
		proto.WriteI64(ctx, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteI64(ctx, 3)
		proto.WriteListEnd(ctx)
		var val raw.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(3, len(val.Elements))
		should.Equal(protocol.TypeI64, val.ElementType)
		iter := c.CreateIterator(val.Elements[0])
		should.Equal(int64(1), iter.ReadInt64())
	}
}

func Test_unmarshal_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.I64, 3)
		proto.WriteI64(ctx, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteI64(ctx, 3)
		proto.WriteListEnd(ctx)
		var val []int64
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal([]int64{int64(1), int64(2), int64(3)}, val)
	}
}

func Test_marshal_general_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		output, err := c.Marshal(general.List{
			int64(1), int64(2), int64(3),
		})
		should.NoError(err)
		iter := c.CreateIterator(output)
		elemType, length := iter.ReadListHeader()
		should.Equal(protocol.TypeI64, elemType)
		should.Equal(3, length)
		should.Equal(uint64(1), iter.ReadUint64())
		should.Equal(uint64(2), iter.ReadUint64())
		should.Equal(uint64(3), iter.ReadUint64())
	}
}

func Test_marshal_raw_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.I64, 3)
		proto.WriteI64(ctx, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteI64(ctx, 3)
		proto.WriteListEnd(ctx)
		var val raw.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		output, err := c.Marshal(val)
		should.NoError(err)
		var generalVal general.List
		should.NoError(c.Unmarshal(output, &generalVal))
		should.Equal(general.List{int64(1), int64(2), int64(3)}, generalVal)
	}
}

func Test_marshal_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal([]int64{1, 2, 3})
		should.NoError(err)
		iter := c.CreateIterator(output)
		elemType, length := iter.ReadListHeader()
		should.Equal(protocol.TypeI64, elemType)
		should.Equal(3, length)
		should.Equal(uint64(1), iter.ReadUint64())
		should.Equal(uint64(2), iter.ReadUint64())
		should.Equal(uint64(3), iter.ReadUint64())
	}
}

func Test_marshal_empty_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal([]int64{})
		should.NoError(err)
		iter := c.CreateIterator(output)
		elemType, length := iter.ReadListHeader()
		should.Equal(protocol.TypeI64, elemType)
		should.Equal(0, length)
	}
}
