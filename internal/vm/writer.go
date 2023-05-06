package vm

import "io"

type Writer interface {
	Value
	io.Writer
}
