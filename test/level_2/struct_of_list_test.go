package test

import (
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jingshouyan/thrifter/general"
	"github.com/jingshouyan/thrifter/protocol"
	"github.com/jingshouyan/thrifter/test"
	"github.com/jingshouyan/thrifter/test/level_2/struct_of_list_test"
	"github.com/stretchr/testify/require"
)

func Test_skip_struct_of_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.LIST, 1)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 1)
		proto.WriteListEnd(ctx)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipStruct(nil))
	}
}

func Test_unmarshal_general_struct_of_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.LIST, 1)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 1)
		proto.WriteListEnd(ctx)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		var val general.Struct
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.List{int64(1)}, val[protocol.FieldId(1)])
	}
}

func Test_unmarshal_struct_of_list(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteStructBegin(ctx, "hello")
		proto.WriteFieldBegin(ctx, "field1", thrift.LIST, 1)
		proto.WriteListBegin(ctx, thrift.I64, 1)
		proto.WriteI64(ctx, 1)
		proto.WriteListEnd(ctx)
		proto.WriteFieldEnd(ctx)
		proto.WriteFieldStop(ctx)
		proto.WriteStructEnd(ctx)
		var val struct_of_list_test.TestObject
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(struct_of_list_test.TestObject{
			[]int64{1},
		}, val)
	}
}

func Test_marshal_general_struct_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		obj := general.Struct{
			protocol.FieldId(1): general.List{
				int64(1),
			},
		}

		output, err := c.Marshal(obj)
		should.NoError(err)
		output1, err := c.Marshal(&obj)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Struct
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.List{int64(1)}, val[protocol.FieldId(1)])
	}
}

func Test_marshal_struct_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		obj := struct_of_list_test.TestObject{
			[]int64{1},
		}

		output, err := c.Marshal(obj)
		should.NoError(err)
		output1, err := c.Marshal(&obj)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.Struct
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.List{int64(1)}, val[protocol.FieldId(1)])
	}
}
