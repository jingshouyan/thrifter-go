package test

import (
	"github.com/jingshouyan/thrifter-go"
	"github.com/jingshouyan/thrifter-go/test/api/binding_test"
	"github.com/v2pro/wombat/generic"
)

var api = thrifter.Config{
	Protocol: thrifter.ProtocolBinary,
}.Froze()

//go:generate go install github.com/jingshouyan/thrifter-go/cmd/thrifter
//go:generate $GOPATH/bin/thrifter -pkg github.com/jingshouyan/thrifter-go/test/api
func init() {
	generic.Declare(func() {
		api.WillDecodeFromBuffer(
			(*binding_test.TestObject)(nil),
		)
	})
}
