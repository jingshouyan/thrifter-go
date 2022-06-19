package test

import (
	"context"
	"testing"

	"github.com/jingshouyan/thrifter-go/test"
	"github.com/stretchr/testify/require"
)

func Test_decode_bool(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteBool(ctx, true)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(true, iter.ReadBool())

		buf, proto = c.CreateProtocol()
		proto.WriteBool(ctx, false)
		iter = c.CreateIterator(buf.Bytes())
		should.Equal(false, iter.ReadBool())
	}
}

func Test_unmarshal_bool(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		var val1 bool
		proto.WriteBool(ctx, true)
		should.NoError(c.Unmarshal(buf.Bytes(), &val1))
		should.Equal(true, val1)

		buf, proto = c.CreateProtocol()
		var val2 bool = true
		proto.WriteBool(ctx, false)
		should.NoError(c.Unmarshal(buf.Bytes(), &val2))
		should.Equal(false, val2)
	}
}

func Test_encode_bool(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteBool(true)
		iter := c.CreateIterator(stream.Buffer())
		should.Equal(true, iter.ReadBool())

		stream = c.CreateStream()
		stream.WriteBool(false)
		iter = c.CreateIterator(stream.Buffer())
		should.Equal(false, iter.ReadBool())
	}
}

func Test_marshal_bool(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(true)
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(true, iter.ReadBool())

		output, err = c.Marshal(false)
		should.NoError(err)
		iter = c.CreateIterator(output)
		should.Equal(false, iter.ReadBool())
	}
}
