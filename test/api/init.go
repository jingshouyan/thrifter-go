package test

import (
	"github.com/jingshouyan/thrifter"
	"github.com/jingshouyan/thrifter/test/api/binding_test"
	"github.com/v2pro/wombat/generic"
)

var api = thrifter.Config{
	Protocol: thrifter.ProtocolBinary,
}.Froze()

//go:generate go install github.com/jingshouyan/thrifter/cmd/thrifter
//go:generate $GOPATH/bin/thrifter -pkg github.com/jingshouyan/thrifter/test/api
func init() {
	generic.Declare(func() {
		api.WillDecodeFromBuffer(
			(*binding_test.TestObject)(nil),
		)
	})
}
