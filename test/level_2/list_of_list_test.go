package test

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jingshouyan/thrifter/general"
	"github.com/jingshouyan/thrifter/test"
	"github.com/stretchr/testify/require"
)

func Test_skip_list_of_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.LIST, 2)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 1)
		proto.WriteListEnd(ctx)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteListEnd(ctx)
		proto.WriteListEnd(ctx)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipList(nil))
	}
}

func Test_unmarshal_general_list_of_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.LIST, 2)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 1)
		proto.WriteListEnd(ctx)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteListEnd(ctx)
		proto.WriteListEnd(ctx)
		var val general.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}

func Test_unmarshal_list_of_general_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.LIST, 2)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 1)
		proto.WriteListEnd(ctx)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteListEnd(ctx)
		proto.WriteListEnd(ctx)
		var val []general.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}

func Test_unmarshal_list_of_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(ctx, thrift.LIST, 2)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 1)
		proto.WriteListEnd(ctx)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 2)
		proto.WriteListEnd(ctx)
		proto.WriteListEnd(ctx)
		var val [][]int64
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal([][]int64{
			{1}, {2},
		}, val)
	}
}

func Test_marshal_general_list_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		lst := general.List{
			general.List{
				int64(1),
			},
			general.List{
				int64(2),
			},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.List
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}

func Test_marshal_list_of_general_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		lst := []general.List{
			{
				int64(1),
			},
			{
				int64(2),
			},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.List
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}

func Test_marshal_list_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		lst := [][]int64{
			{1}, {2},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.List
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}
