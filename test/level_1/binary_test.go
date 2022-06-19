package test

import (
	"context"
	"testing"

	"github.com/jingshouyan/thrifter/test"
	"github.com/stretchr/testify/require"
)

func Test_decode_binary(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteBinary(ctx, []byte("hello"))
		iter := c.CreateIterator(buf.Bytes())
		should.Equal("hello", string(iter.ReadBinary()))
	}
}

func Test_unmarshal_binary(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteBinary(ctx, []byte("hello"))
		var val []byte
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal("hello", string(val))
	}
}

func Test_encode_binary(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteBinary([]byte(`hello world!`))
		iter := c.CreateIterator(stream.Buffer())
		should.Equal([]byte(`hello world!`), iter.ReadBinary())
	}
}

func Test_marshal_binary(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		val := []byte("hello")
		output, err := c.Marshal(val)
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal("hello", string(iter.ReadBinary()))
	}
}
