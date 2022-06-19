package codegen

import (
	"reflect"

	"github.com/jingshouyan/thrifter-go/spi"
)

type Extension struct {
	spi.Extension
	ExtTypes []reflect.Type
}

func (ext *Extension) MangledName() string {
	// TODO: hash extension to represent different config
	return "default"
}
