package test

import (
	"context"
	"testing"

	"github.com/jingshouyan/thrifter/test"
	"github.com/stretchr/testify/require"
)

func Test_decode_int64(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI64(ctx, -1)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(int64(-1), iter.ReadInt64())
	}
}

func Test_unmarshal_int64(t *testing.T) {
	should := require.New(t)
	ctx := context.Background()
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI64(ctx, -1)
		var val int64
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(int64(-1), val)
	}
}

func Test_encode_int64(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteInt64(-1)
		iter := c.CreateIterator(stream.Buffer())
		should.Equal(int64(-1), iter.ReadInt64())
	}
}

func Test_marshal_int64(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(int64(-1))
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(int64(-1), iter.ReadInt64())
	}
}
